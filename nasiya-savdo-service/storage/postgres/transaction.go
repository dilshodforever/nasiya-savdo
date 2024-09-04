package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/dilshodforever/nasiya-savdo/kafka"
	connect"github.com/dilshodforever/nasiya-savdo/kafkaconnect"
	"github.com/dilshodforever/nasiya-savdo/kafkasender"
	"github.com/dilshodforever/nasiya-savdo/model"
	"github.com/google/uuid"
)

type TransactionStorage struct {
	db    *sql.DB
	kafka kafka.KafkaProducer
}

func NewTransactionStorage(db *sql.DB, kafka kafka.KafkaProducer) *TransactionStorage {
	return &TransactionStorage{db: db, kafka: kafka}
}

func (p *TransactionStorage) CreateTransaction(req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	// Step 1: Retrieve amount and price from the exchange table for the given contract_id
	query := `SELECT amount, price FROM exchange WHERE contract_id = $1 and deleted_at = 0`
	rows, err := p.db.Query(query, req.ContractId)
	if err != nil {
		log.Printf("Failed to retrieve exchange data: %v", err)
		return nil, fmt.Errorf("failed to retrieve exchange data: %v", err)
	}
	defer rows.Close()

	var (
		amount int32
		price  float64
		total  float64
	)

	for rows.Next() {
		if err := rows.Scan(&amount, &price); err != nil {
			log.Printf("Failed to scan exchange data: %v", err)
			return nil, fmt.Errorf("failed to scan exchange data: %v", err)
		}
		total += price * float64(amount)
	}

	// Step 2: Calculate duration based on total price and the requested price
	duration := total / req.Price
	fmt.Println(duration, total, req.Price)

	// Step 3: Insert the new transaction into the transactions table
	query = `
		INSERT INTO transactions (id, contract_id, price, duration, created_at)
		VALUES ($1, $2, $3, $4, now())
	`
	_, err = p.db.Exec(query, uuid.New().String(), req.ContractId, req.Price, duration)
	if err != nil {
		log.Printf("Failed to create transaction: %v", err)
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	// Step 4: Calculate the total price and duration from all transactions for the contract
	query = `
		SELECT 
			SUM(t.price) AS total_price, 
			SUM(t.duration) AS total_duration, 
			c.duration AS contract_duration
		FROM 
			transactions AS t
		JOIN 
			contract AS c ON c.id = t.contract_id
		WHERE 
			t.contract_id = $1
		GROUP BY 
			c.duration
	`
	var (
		totalPrice       float64
		totalDuration    int32
		contractDuration int32
	)
	err = p.db.QueryRow(query, req.ContractId).Scan(&totalPrice, &totalDuration, &contractDuration)
	if err != nil {
		log.Printf("Failed to calculate totals: %v", err)
		return nil, fmt.Errorf("failed to calculate totals: %v", err)
	}

	// Step 5: Update the contract status if the total duration meets or exceeds the contract duration
	if totalDuration >= contractDuration {
		query = `UPDATE contract SET status = 'finished', deleted_at = now() WHERE id = $1`
		_, err = p.db.Exec(query, req.ContractId)
		if err != nil {
			log.Printf("Failed to update contract status: %v", err)
			return nil, fmt.Errorf("failed to update contract status: %v", err)
		}
		return &pb.TransactionResponse{
			Message: "Transaction created successfully. The contract is now finished.",
			Success: true,
		}, nil
	}

	remainingPending := totalPrice - req.Price
	return &pb.TransactionResponse{
		Message: fmt.Sprintf("Transaction created successfully. Remaining balance: %.2f", remainingPending),
		Success: true,
	}, nil
}

func (p *TransactionStorage) GetTransaction(req *pb.TransactionIdRequest) (*pb.GetTransactionResponse, error) {
	query := `
		SELECT id, contract_id, price, duration, created_at
		FROM transactions
		WHERE id = $1
	`
	var transaction pb.GetTransactionResponse
	err := p.db.QueryRow(query, req.Id).Scan(
		&transaction.Id, &transaction.ContractId, &transaction.Price, &transaction.Duration, &transaction.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
func (p *TransactionStorage) UpdateTransaction(req *pb.UpdateTransactionRequest) (*pb.TransactionResponse, error) {
	// Initialize a map to hold the fields that need to be updated
	update := map[string]interface{}{}
	if req.ContractId != "" {
		update["contract_id"] = req.ContractId
	}
	if req.Price != 0 {
		update["price"] = req.Price
	}
	if req.Duration != 0 {
		update["duration"] = req.Duration
	}

	// Check if there's anything to update
	if len(update) == 0 {
		return &pb.TransactionResponse{Message: "Nothing to update", Success: false}, nil
	}

	// Build the dynamic SQL query based on the fields to be updated
	setClauses := []string{}
	args := []interface{}{}
	argCount := 1

	for column, value := range update {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, argCount))
		args = append(args, value)
		argCount++
	}

	// Join the SET clauses
	setQuery := strings.Join(setClauses, ", ")

	// Final update query
	query := fmt.Sprintf(`
		UPDATE transactions
		SET %s, updated_at = now()
		WHERE id = $%d
	`, setQuery, argCount)

	// Append the id as the last argument
	args = append(args, req.Id)

	// Execute the query
	_, err := p.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.TransactionResponse{
		Message: "Transaction successfully updated",
		Success: true,
	}, nil
}

func (p *TransactionStorage) DeleteTransaction(req *pb.TransactionIdRequest) (*pb.TransactionResponse, error) {
	query := `
		DELETE FROM transactions
		WHERE id = $1
	`
	_, err := p.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.TransactionResponse{
		Message: "Transaction successfully deleted",
		Success: true,
	}, nil
}

func (p *TransactionStorage) ListTransactions(req *pb.GetAllTransactionRequest) (*pb.GetAllTransactionResponse, error) {
	transactions := pb.GetAllTransactionResponse{}
	query := `
		SELECT id, contract_id, price, duration, created_at
		FROM transactions
		WHERE true
	`
	var args []interface{}
	count := 1

	if req.ContractId != "" {
		query += fmt.Sprintf(" AND contract_id = $%d", count)
		args = append(args, req.ContractId)
		count++
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction pb.GetTransactionResponse
		err := rows.Scan(
			&transaction.Id, &transaction.ContractId, &transaction.Price, &transaction.Duration, &transaction.CreatedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		transactions.AllTransactions = append(transactions.AllTransactions, &transaction)
	}

	return &transactions, nil
}

// V2

func (p *TransactionStorage) CheckTransactions(req *pb.CheckRequest) (*pb.CheckResponse, error) {
	query := `
		SELECT c.id, c.created_at, COALESCE(SUM(t.duration), 0) as total_duration
		FROM contract AS c
		LEFT JOIN transactions AS t ON t.contract_id = c.id
		WHERE c.status = 'pending'
		GROUP BY c.id, c.created_at
	`
	rows, err := p.db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contractID, createdAt string
		var totalDuration sql.NullInt32

		if err := rows.Scan(&contractID, &createdAt, &totalDuration); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		creationTime, err := time.Parse(time.RFC3339, createdAt)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			return nil, err
		}

		// Calculate the due date
		durationMonths := int32(0)
		if totalDuration.Valid {
			durationMonths = totalDuration.Int32
		}
		dueDate := creationTime.AddDate(0, int(durationMonths), 0)

		// Check the due date and send notifications
		now := time.Now()
		switch {
		case dueDate.Year() == now.Year() && dueDate.Month() == now.Month() && dueDate.Day() == now.Day():
			kafka:=connect.ConnectToKafka()
			kafkasender.CreateNotification(kafka, model.Send{
				Userid:     req.UserId,
				Message:    "Payment due today for contract " + contractID,
				ContractId: contractID,
			})

		case dueDate.Before(now):
			kafka:=connect.ConnectToKafka()
			kafkasender.CreateNotification(kafka, model.Send{
				Userid:     req.UserId,
				Message:    "Payment overdue by 1 month for contract " + contractID,
				ContractId: contractID,
			})

		case dueDate.Month() == now.Month() && dueDate.Day() > now.Day():
			kafka:=connect.ConnectToKafka()
			kafkasender.CreateNotification(kafka, model.Send{
				Userid:     req.UserId,
				Message:    "Payment due this month for contract " + contractID,
				ContractId: contractID,
			})
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return nil, err
	}
	return &pb.CheckResponse{Message: "Payment checks completed successfully."}, nil
}



func (p *TransactionStorage) TestNotification(req *pb.Testresponse) (*pb.Testrequest, error) {
	kafka:=connect.ConnectToKafka()
	kafkasender.CreateNotification(kafka, model.Send{
		Userid:     "1111",
		Message:    "Payment due this month for contract " + "contractID",
		ContractId: "2222222",
	})
	return &pb.Testrequest{Message: "Success"}, nil
}

// V1

// func (p *TransactionStorage) CheckTransactions(req *pb.CheckRequest) (*pb.CheckResponse, error) {
// 	query := `
// 		SELECT c.id, c.created_at, SUM(t.duration)
// 		FROM contract AS c
// 		LEFT JOIN transactions AS t ON t.contract_id = c.id
// 		WHERE c.status = 'pending'
// 		GROUP BY c.id, c.created_at
// 	`
// 	rows, err := p.db.Query(query)
// 	if err != nil {
// 		log.Printf("Error executing query: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var dueIDs []string

// 	for rows.Next() {
// 		var contractID, createdAt string
// 		var totalDuration sql.NullInt32

// 		if err := rows.Scan(&contractID, &createdAt, &totalDuration); err != nil {
// 			log.Printf("Error scanning row: %v", err)
// 			return nil, err
// 		}
// 		creationTime, err := time.Parse(time.RFC3339, createdAt)
// 		if err != nil {
// 			log.Printf("Error parsing date: %v", err)
// 			return nil, err
// 		}

// 		duration := int32(0)

// 		if totalDuration.Valid {
// 			duration = totalDuration.Int32
// 		}

// 		dueDay := int(creationTime.Month()) + int(duration)
// 		switch {
// 		case dueDay == int(time.Now().Month()) && creationTime.Day() == time.Now().Day():
// 			kafkasender.CreateNotification(p.kafka, model.Send{Userid: req.UserId,Message: "This contact's paymant day come today", ContractId: contractID})
// 		case dueDay >= int(time.Now().Month()) && creationTime.Day() >= time.Now().Day():
// 			kafkasender.CreateNotification(p.kafka, model.Send{Userid: req.UserId,Message: "1 oydan oship ketgan tolanmagan contact  ", ContractId: contractID})

// 		case dueDay == int(time.Now().Month()) && creationTime.Day() >= time.Now().Day():
// 			kafkasender.CreateNotification(p.kafka, model.Send{Userid: req.UserId,Message: "shu oy tolanishi kere bolgan contract", ContractId: contractID})
// 		}

// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("Error during rows iteration: %v", err)
// 		return nil, err
// 	}

// 	if len(dueIDs) > 0 {
// 		return &pb.CheckResponse{Message: "" + dueIDs[0]}, nil
// 	}

// 	return &pb.CheckResponse{Message: "No payments are due today."}, nil
// }

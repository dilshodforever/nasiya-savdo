package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
)

type TransactionStorage struct {
	db *sql.DB
}

func NewTransactionStorage(db *sql.DB) *TransactionStorage {
	return &TransactionStorage{db: db}
}

func (p *TransactionStorage) CreateTransaction(req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	// Step 1: Retrieve amount and price from the exchange table for the given contract_id
	query := `SELECT amount, price FROM exchange WHERE contract_id = $1`
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

	// Step 3: Insert the new transaction into the transactions table
	query = `
		INSERT INTO transactions (contract_id, price, duration, created_at)
		VALUES ($1, $2, $3, now())
	`
	_, err = p.db.Exec(query, req.ContractId, req.Price, duration)
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

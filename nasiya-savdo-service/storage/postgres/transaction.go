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
	query := `
		INSERT INTO transactions (contract_id, price, duration, created_at)
		VALUES ($1, $2, $3, now())
		RETURNING id
	`
	var id string
	err := p.db.QueryRow(query, req.ContractId, req.Price, req.Duration).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.TransactionResponse{
		Message: "Transaction successfully created with ID: " + id,
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

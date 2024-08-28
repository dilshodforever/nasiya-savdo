package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/google/uuid"
)

type ExchangeStorage struct {
	db *sql.DB
}

func NewExchangeStorage(db *sql.DB) *ExchangeStorage {
	return &ExchangeStorage{db: db}
}

func (p *ExchangeStorage) CreateExchange(req *pb.CreateExchangeRequest) (*pb.ExchangeResponse, error) {
	id :=uuid.NewString()
	if req.Status == `sell` {
		query := `
			INSERT INTO exchange (product_id, amount, price, status, contract_id, created_at, deleted_at)
			VALUES ($1, $2, $3, $4, $5, $6, now(), 0)
		`
		_, err := p.db.Exec(query, id, req.ProductId, req.Amount, req.Price, req.Status, req.ContractId)
		if err != nil {
			return nil, err
		}
	} else{
		query := `
		INSERT INTO exchange (id, product_id, amount, price, status,  created_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, now(), 0)
	`
	_, err := p.db.Exec(query, id, req.ProductId, req.Amount, req.Price, req.Status)
	if err != nil {
		return nil, err
	}
	}
	return &pb.ExchangeResponse{
		Message: "Exchange created successfully ",
		Success: true,
	}, nil
}

func (p *ExchangeStorage) GetExchange(req *pb.ExchangeIdRequest) (*pb.GetExchangeResponse, error) {
	query := `
		SELECT id, product_id, amount, price, status, contract_id, created_at, updated_at, deleted_at
		FROM exchange
		WHERE id = $1 AND deleted_at = 0
	`
	var exchange pb.GetExchangeResponse
	err := p.db.QueryRow(query, req.Id).Scan(
		&exchange.Id, &exchange.ProductId, &exchange.Amount, &exchange.Price, &exchange.Status,
		&exchange.ContractId, &exchange.CreatedAt, &exchange.UpdatedAt, &exchange.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &exchange, nil
}

func (p *ExchangeStorage) UpdateExchange(req *pb.UpdateExchangeRequest) (*pb.ExchangeResponse, error) {
	// Initialize a map to hold the fields that need to be updated
	update := map[string]interface{}{}
	if req.ProductId != "" {
		update["product_id"] = req.ProductId
	}
	if req.Amount != 0 {
		update["amount"] = req.Amount
	}
	if req.Price != 0 {
		update["price"] = req.Price
	}
	if req.Status != "" {
		update["status"] = req.Status
	}
	
	if req.ContractId != "" {
		update["contract_id"] = req.ContractId
	}

	// Check if there's anything to update
	if len(update) == 0 {
		return &pb.ExchangeResponse{Message: "Nothing to update", Success: false}, nil
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
		UPDATE exchange
		SET %s, updated_at = now()
		WHERE id = $%d AND deleted_at = 0
	`, setQuery, argCount)

	// Append the id as the last argument
	args = append(args, req.Id)

	// Execute the query
	_, err := p.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.ExchangeResponse{
		Message: "Exchange successfully updated",
		Success: true,
	}, nil
}

func (p *ExchangeStorage) DeleteExchange(req *pb.ExchangeIdRequest) (*pb.ExchangeResponse, error) {
	query := `
		UPDATE exchange
		SET deleted_at = now()
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ExchangeResponse{
		Message: "Exchange successfully deleted",
		Success: true,
	}, nil
}

func (p *ExchangeStorage) ListExchanges(req *pb.GetAllExchangeRequest) (*pb.GetAllExchangeResponse, error) {
	exchanges := pb.GetAllExchangeResponse{}
	query := `
		SELECT id, product_id, amount, price, status, contract_id, created_at, updated_at, deleted_at
		FROM exchange
		WHERE deleted_at = 0
	`
	var args []interface{}
	count := 1

	if req.ProductId != "" {
		query += fmt.Sprintf(" AND product_id = $%d", count)
		args = append(args, req.ProductId)
		count++
	}
	if req.Status != "" {
		query += fmt.Sprintf(" AND status = $%d", count)
		args = append(args, req.Status)
		count++
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var exchange pb.GetExchangeResponse
		err := rows.Scan(
			&exchange.Id, &exchange.ProductId, &exchange.Amount, &exchange.Price, &exchange.Status,
			&exchange.ContractId, &exchange.CreatedAt, &exchange.UpdatedAt, &exchange.DeletedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		exchanges.AllExchanges = append(exchanges.AllExchanges, &exchange)
	}

	return &exchanges, nil
}

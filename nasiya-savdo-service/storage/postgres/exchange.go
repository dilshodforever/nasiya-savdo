package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

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
	id := uuid.NewString()
	if req.Status == `sell` {
		query := `
		SELECT SUM(
			CASE 
				WHEN status = 'buy' THEN amount 
				WHEN status = 'sell' THEN -amount 
				ELSE 0 
			END
		) 
		FROM exchange 
		WHERE deleted_at = 0 AND product_id = $1`
		var amount int
		err := p.db.QueryRow(query, req.ProductId).Scan(&amount)
		if err != nil {
			return nil, err
		}
		if amount < int(req.Amount) {
			return nil, fmt.Errorf("insufficient product quantity available")
		}
		query = `
			INSERT INTO exchange (id, product_id, amount, price, status, contract_id, created_at, deleted_at)
			VALUES ($1, $2, $3, $4, $5, $6, now(), 0)
		`
		_, err = p.db.Exec(query, id, req.ProductId, req.Amount, req.Price, req.Status, req.ContractId)
		if err != nil {
			return nil, err
		}
	} else {
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
	var contractId sql.NullString
	err := p.db.QueryRow(query, req.Id).Scan(
		&exchange.Id, &exchange.ProductId, &exchange.Amount, &exchange.Price, &exchange.Status,
		&contractId, &exchange.CreatedAt, &exchange.UpdatedAt, &exchange.DeletedAt,
	)
	// Handle the nullable field
	if contractId.Valid {
		exchange.ContractId = contractId.String
	} else {
		exchange.ContractId = "" // Or set it to a default value
	}
	if err != nil {
		return nil, err
	}

	return &exchange, nil
}

func (p *ExchangeStorage) UpdateExchange(req *pb.UpdateExchangeRequest) (*pb.ExchangeResponse, error) {

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
		SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, req.Id, time.Now().Unix())
	if err != nil {
		return nil, err
	}

	return &pb.ExchangeResponse{
		Message: "Exchange successfully deleted",
		Success: true,
	}, nil
}

func (p *ExchangeStorage) ListExchanges(req *pb.GetAllExchangeRequest) (*pb.GetAllExchangeResponse, error) {
	exchanges := pb.GetAllExchangeResponse{
		AllExchanges: []*pb.GetExchangeResponse{}, // Explicitly initialize to an empty slice
	}

	// SQL query with corrected JOIN condition
	query := `
		SELECT e.id, e.product_id, e.amount, e.price, e.status, e.contract_id, e.created_at, e.updated_at, e.deleted_at
		FROM exchange as e
		JOIN products as p ON p.id = e.product_id
		JOIN storage as s ON s.id = p.storage_id
		WHERE e.deleted_at = 0
	`

	var args []interface{}
	count := 1

	// Adding filtering conditions based on the request
	if req.ProductId != "" {
		query += fmt.Sprintf(" AND e.product_id = $%d", count)
		args = append(args, req.ProductId)
		count++
	}
	if req.Status != "" {
		query += fmt.Sprintf(" AND e.status = $%d", count)
		args = append(args, req.Status)
		count++
	}

	// Execute the query
	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan through the result set
	for rows.Next() {
		var exchange pb.GetExchangeResponse
		var contractId sql.NullString

		// Scanning data into fields
		err := rows.Scan(
			&exchange.Id, &exchange.ProductId, &exchange.Amount, &exchange.Price, &exchange.Status,
			&contractId, &exchange.CreatedAt, &exchange.UpdatedAt, &exchange.DeletedAt,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Handle the nullable contract ID
		if contractId.Valid {
			exchange.ContractId = contractId.String
		} else {
			exchange.ContractId = ""
		}

		// Append to the list of exchanges
		exchanges.AllExchanges = append(exchanges.AllExchanges, &exchange)
	}

	// Check for errors after iterating through rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Return the list of exchanges
	return &exchanges, nil
}

package postgres

import (
	"database/sql"
	"fmt"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
)

type ExchangeStorage struct {
	db *sql.DB
}

func NewExchangeStorage(db *sql.DB) *ExchangeStorage {
	return &ExchangeStorage{db: db}
}

func (p *ExchangeStorage) CreateExchange(req *pb.CreateExchangeRequest) (*pb.ExchangeResponse, error) {
	query := `
		INSERT INTO exchange (product_id, amount, price, status, contract_id, created_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, now(), 0)
		RETURNING id
	`
	var id string
	err := p.db.QueryRow(query, req.ProductId, req.Amount, req.Price, req.Status, req.ContractId).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.ExchangeResponse{
		Message: "Exchange successfully created with ID: " + id,
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
	query := `
		UPDATE exchange
		SET product_id = $1, amount = $2, price = $3, status = $4, contract_id = $5, updated_at = now()
		WHERE id = $6 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, req.ProductId, req.Amount, req.Price, req.Status, req.ContractId, req.Id)
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

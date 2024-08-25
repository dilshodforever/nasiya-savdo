package postgres

import (
	"database/sql"
	"fmt"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
)

type ContractStorage struct {
	db *sql.DB
}

func NewContractStorage(db *sql.DB) *ContractStorage {
	return &ContractStorage{db: db}
}

func (p *ContractStorage) CreateContract(req *pb.CreateContractRequest) (*pb.ContractResponse, error) {
	query := `
		INSERT INTO contract (consumer_name, consumer_passport_serial, consumer_address, passport_image, status, duration, created_at, deleted_at)
		VALUES ($1, $2, $3, $4, 'pending', $6, now(), 0)
		RETURNING id
	`
	var id string
	err := p.db.QueryRow(query, req.ConsumerName, req.ConsumerPassportSerial, req.ConsumerAddress, req.PassportImage, req.Status, req.Duration).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.ContractResponse{
		Message: "Contract successfully created with ID: " + id,
		Success: true,
	}, nil
}

func (p *ContractStorage) GetContract(req *pb.ContractIdRequest) (*pb.GetContractResponse, error) {
	query := `
		SELECT id, consumer_name, consumer_passport_serial, consumer_address, passport_image, status, duration, created_at, deleted_at
		FROM contract
		WHERE id = $1 AND deleted_at = 0
	`
	var contract pb.GetContractResponse
	err := p.db.QueryRow(query, req.Id).Scan(
		&contract.Id, &contract.ConsumerName, &contract.ConsumerPassportSerial, &contract.ConsumerAddress,
		&contract.PassportImage, &contract.Status, &contract.Duration, &contract.CreatedAt, &contract.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &contract, nil
}

func (p *ContractStorage) UpdateContract(req *pb.UpdateContractRequest) (*pb.ContractResponse, error) {
	query := `
		UPDATE contract
		SET consumer_name = $1, consumer_passport_serial = $2, consumer_address = $3, passport_image = $4,
		    status = $5, duration = $6, updated_at = now()
		WHERE id = $7 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, req.ConsumerName, req.ConsumerPassportSerial, req.ConsumerAddress, req.PassportImage, req.Status, req.Duration, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ContractResponse{
		Message: "Contract successfully updated",
		Success: true,
	}, nil
}

func (p *ContractStorage) DeleteContract(req *pb.ContractIdRequest) (*pb.ContractResponse, error) {
	query := `
		UPDATE contract
		SET deleted_at = now()
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ContractResponse{
		Message: "Contract successfully deleted",
		Success: true,
	}, nil
}

func (p *ContractStorage) ListContracts(req *pb.GetAllContractRequest) (*pb.GetAllContractResponse, error) {
	contracts := pb.GetAllContractResponse{}
	query := `
		SELECT id, consumer_name, consumer_passport_serial, consumer_address, passport_image, status, duration, created_at, deleted_at
		FROM contract
		WHERE deleted_at = 0
	`
	var args []interface{}
	count := 1

	if req.ConsumerName != "" {
		query += fmt.Sprintf(" AND consumer_name ILIKE $%d", count)
		args = append(args, "%"+req.ConsumerName+"%")
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
		var contract pb.GetContractResponse
		err := rows.Scan(
			&contract.Id, &contract.ConsumerName, &contract.ConsumerPassportSerial, &contract.ConsumerAddress,
			&contract.PassportImage, &contract.Status, &contract.Duration, &contract.CreatedAt, &contract.DeletedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		contracts.AllContracts = append(contracts.AllContracts, &contract)
	}

	return &contracts, nil
}

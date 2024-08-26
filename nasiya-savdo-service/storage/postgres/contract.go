package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
	update := map[string]interface{}{}
	if req.ConsumerName != "" {
		update["consumer_name"] = req.ConsumerName
	}
	if req.ConsumerPassportSerial != "" {
		update["consumer_passport_serial"] = req.ConsumerPassportSerial
	}
	if req.ConsumerAddress != "" {
		update["consumer_address"] = req.ConsumerAddress
	}
	if req.PassportImage != "" {
		update["passport_image"] = req.PassportImage
	}
	if req.Status != "" {
		update["status"] = req.Status
	}
	if req.Duration != 0 {
		update["duration"] = req.Duration
	}

	if len(update) == 0 {
		return &pb.ContractResponse{Message: "Nothing to update", Success: false}, nil
	}

	setClauses := []string{}
	args := []interface{}{}
	argCount := 1

	for column, value := range update {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, argCount))
		args = append(args, value)
		argCount++
	}

	setQuery := strings.Join(setClauses, ", ")

	query := fmt.Sprintf(`
		UPDATE contract
		SET %s, updated_at = now()
		WHERE id = $%d AND deleted_at = 0
	`, setQuery, argCount)

	args = append(args, req.Id)

	_, err := p.db.Exec(query, args...)
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

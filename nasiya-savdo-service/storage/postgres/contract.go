package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/google/uuid"
)

type ContractStorage struct {
	db *sql.DB
}

func NewContractStorage(db *sql.DB) *ContractStorage {
	return &ContractStorage{db: db}
}

func (p *ContractStorage) CreateContract(req *pb.CreateContractRequest) (*pb.ContractResponse, error) {
	id:=uuid.NewString()
	query := `
		INSERT INTO contract (id, consumer_name, consumer_passport_serial, consumer_address, consumer_phone_number, passport_image, status, duration, created_at)
		VALUES ($1, $2, $3, $4, $5, $6,'pending', $7, now())
	`
	_,err := p.db.Exec(query, id,req.ConsumerName, req.ConsumerPassportSerial, req.ConsumerAddress, req.PassportImage, req.Status, req.Duration)
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
		SELECT id, consumer_name, consumer_passport_serial, consumer_address, consumer_phone_number, passport_image, status, duration, created_at, deleted_at
		FROM contract
		WHERE id = $1
	`
	var contract pb.GetContractResponse
	var deletedAt sql.NullString
	err := p.db.QueryRow(query, req.Id).Scan(
		&contract.Id, &contract.ConsumerName, &contract.ConsumerPassportSerial, &contract.ConsumerAddress, &contract.ConsumerPhoneNumber,
		&contract.PassportImage, &contract.Status, &contract.Duration, &contract.CreatedAt, &deletedAt,
	)
	if deletedAt.Valid{
		contract.DeletedAt=deletedAt.String
	} else{
		contract.DeletedAt=""
	}
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
	if req.ConsumerPhoneNumber != ""{
		update["consumer_phone_number"] = req.ConsumerPhoneNumber
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
		WHERE id = $1 
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
		SELECT id, consumer_name, consumer_passport_serial, consumer_address, consumer_phone_number, passport_image, status, duration, created_at, deleted_at
		FROM contract
		WHERE true
	`
	var args []interface{}
	count := 1

	if req.ConsumerName != "" {
		query += fmt.Sprintf(" AND consumer_name ILIKE $%d", count)
		args = append(args, "%"+req.ConsumerName+"%")
		count++
	}

	if req.PasportSeria != "" { 
		query += fmt.Sprintf(" AND consumer_passport_serial ILIKE $%d", count)
		args = append(args, "%"+req.PasportSeria+"%")
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
		var deletedAt sql.NullString // Use sql.NullString for nullable fields

		err := rows.Scan(
			&contract.Id, &contract.ConsumerName, &contract.ConsumerPassportSerial, &contract.ConsumerAddress, &contract.ConsumerPhoneNumber,
			&contract.PassportImage, &contract.Status, &contract.Duration, &contract.CreatedAt, &deletedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Handle the nullable field
		if deletedAt.Valid {
			contract.DeletedAt = deletedAt.String
		} else {
			contract.DeletedAt = "" // Or set it to a default value
		}

		contracts.AllContracts = append(contracts.AllContracts, &contract)
	}

	return &contracts, nil
}



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

type ContractStorage struct {
	db *sql.DB
}

func NewContractStorage(db *sql.DB) *ContractStorage {
	return &ContractStorage{db: db}
}

func (p *ContractStorage) CreateContract(req *pb.CreateContractRequest) (*pb.ContractResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO contract (id, storage_id, consumer_name, consumer_passport_serial, consumer_address, consumer_phone_number, passport_image, status, duration, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, 'pending', $8, now())
	`
	_, err := p.db.Exec(query, id, req.StorageId, req.ConsumerName, req.ConsumerPassportSerial, req.ConsumerAddress, req.ConsumerPhoneNumber, req.PassportImage, req.Duration)
	if err != nil {
		return nil, err
	}

	return &pb.ContractResponse{
		Message: "Contract successfully created with ID: " + id,
		Success: true,
	}, nil
}

func (p *ContractStorage) GetContract(req *pb.ContractIdRequest) (*pb.GetContractResponse, error) {
	// Query for the contract details
	query := `
		SELECT id, consumer_name, consumer_passport_serial, consumer_address, consumer_phone_number, passport_image, status, duration, created_at, deleted_at
		FROM contract
		WHERE id = $1 and deleted_at=0
	`
	var contract pb.GetContractResponse
	var deletedAt sql.NullString

	err := p.db.QueryRow(query, req.Id).Scan(
		&contract.Id, &contract.ConsumerName, &contract.ConsumerPassportSerial, &contract.ConsumerAddress, &contract.ConsumerPhoneNumber,
		&contract.PassportImage, &contract.Status, &contract.Duration, &contract.CreatedAt, &deletedAt,
	)

	if err != nil {
		log.Printf("Error querying contract with ID %s: %v", req.Id, err)
		return nil, err
	}

	// Handle deleted_at field
	if deletedAt.Valid {
		contract.DeletedAt = deletedAt.String
	} else {
		contract.DeletedAt = ""
	}

	// Query for the exchange details
	var amount int32
	var price float64
	query = `
		SELECT amount, price 
		FROM exchange
		WHERE contract_id = $1 AND deleted_at = 0
	`
	rows, err := p.db.Query(query, req.Id) // Pass req.Id as the parameter
	if err != nil {
		log.Printf("Error querying exchange details for contract ID %s: %v", req.Id, err)
		return nil, err
	}
	defer rows.Close()

	// Calculate total price
	for rows.Next() {
		err := rows.Scan(&amount, &price)
		if err != nil {
			log.Printf("Error scanning exchange details for contract ID %s: %v", req.Id, err)
			return nil, err
		}
		contract.Price += float64(amount) * price
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating exchange rows for contract ID %s: %v", req.Id, err)
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
	if req.ConsumerPhoneNumber != "" {
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
		WHERE id = $%d 
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
		SET deleted_at = $2, status='canceled'
		WHERE id = $1 
	`
	_, err := p.db.Exec(query, req.Id, time.Now().Unix())
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
		FROM contract where deleted_at=0
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

	// Apply limit and page-based offset if Page is provided
	if req.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", count)
		args = append(args, req.Limit)
		count++

		// Only apply offset if Page is set
		if req.Page > 0 {
			offset := (req.Page - 1) * req.Limit
			query += fmt.Sprintf(" OFFSET $%d", count)
			args = append(args, offset)
		}
	}

	// Execute the contract query
	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contract pb.GetContractResponse
		var deletedAt sql.NullString

		err := rows.Scan(
			&contract.Id, &contract.ConsumerName, &contract.ConsumerPassportSerial, &contract.ConsumerAddress, &contract.ConsumerPhoneNumber,
			&contract.PassportImage, &contract.Status, &contract.Duration, &contract.CreatedAt, &deletedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if deletedAt.Valid {
			contract.DeletedAt = deletedAt.String
		} else {
			contract.DeletedAt = ""
		}

		// Query the exchange details for the contract
		exchangeQuery := `
			SELECT amount, price 
			FROM exchange
			WHERE contract_id = $1 AND deleted_at = 0
		`
		exchangeRows, err := p.db.Query(exchangeQuery, contract.Id)
		if err != nil {
			log.Printf("Error querying exchange details for contract ID %s: %v", contract.Id, err)
			return nil, err
		}
		defer exchangeRows.Close()

		// Calculate total price
		for exchangeRows.Next() {
			var amount int32
			var price float64
			err := exchangeRows.Scan(&amount, &price)
			if err != nil {
				log.Printf("Error scanning exchange details for contract ID %s: %v", contract.Id, err)
				return nil, err
			}
			contract.Price += float64(amount) * price
		}

		if err := exchangeRows.Err(); err != nil {
			log.Printf("Error iterating exchange rows for contract ID %s: %v", contract.Id, err)
			return nil, err
		}

		contracts.AllContracts = append(contracts.AllContracts, &contract)
	}

	countQuery := `SELECT COUNT(1) FROM contract where deleted_at=0`
	err = p.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, err
	}
	contracts.Count = int32(count)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &contracts, nil
}




// func (p *ContractStorage) ListContracts(req *pb.GetAllContractRequest) (*pb.GetAllContractResponse, error) {
// 	contracts := pb.GetAllContractResponse{}
// 	// query := `
// 	// 	SELECT id, consumer_name, consumer_passport_serial, consumer_address, consumer_phone_number, passport_image, status, duration, created_at, deleted_at
// 	// 	FROM contract
// 	// 	WHERE storage_id = $1
// 	// `
// 	query := `
// 		SELECT id, consumer_name, consumer_passport_serial, consumer_address, consumer_phone_number, passport_image, status, duration, created_at, deleted_at
// 		FROM contract 
// 	`
// 	var args []interface{}
// 	count := 1
// 	// args = append(args, req.StorageId)
// 	if req.ConsumerName != "" {
// 		query += fmt.Sprintf(" AND consumer_name ILIKE $%d", count)
// 		args = append(args, "%"+req.ConsumerName+"%")
// 		count++
// 	}

// 	if req.PasportSeria != "" {
// 		query += fmt.Sprintf(" AND consumer_passport_serial ILIKE $%d", count)
// 		args = append(args, "%"+req.PasportSeria+"%")
// 		count++
// 	}

// 	if req.Status != "" {
// 		query += fmt.Sprintf(" AND status = $%d", count)
// 		args = append(args, req.Status)
// 		count++
// 	}

// 	if req.Limit != 0 || req.Offset != 0 {
// 		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", count, count+1)
// 		args = append(args, req.Limit, req.Offset)
// 	}

// 	// Execute the contract query
// 	rows, err := p.db.Query(query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var contract pb.GetContractResponse
// 		var deletedAt sql.NullString // Use sql.NullString for nullable fields

// 		err := rows.Scan(
// 			&contract.Id, &contract.ConsumerName, &contract.ConsumerPassportSerial, &contract.ConsumerAddress, &contract.ConsumerPhoneNumber,
// 			&contract.PassportImage, &contract.Status, &contract.Duration, &contract.CreatedAt, &deletedAt,
// 		)
// 		if err != nil {
// 			log.Println(err)
// 			return nil, err
// 		}

// 		if deletedAt.Valid {
// 			contract.DeletedAt = deletedAt.String
// 		} else {
// 			contract.DeletedAt = "" // Or set it to a default value
// 		}

// 		// Query the exchange details for the contract
// 		exchangeQuery := `
// 			SELECT amount, price 
// 			FROM exchange
// 			WHERE contract_id = $1 AND deleted_at = 0
// 		`
// 		exchangeRows, err := p.db.Query(exchangeQuery, contract.Id) // New variable for exchange query rows
// 		if err != nil {
// 			log.Printf("Error querying exchange details for contract ID %s: %v", contract.Id, err)
// 			return nil, err
// 		}
// 		defer exchangeRows.Close()

// 		// Calculate total price
// 		for exchangeRows.Next() {
// 			var amount int32
// 			var price float64
// 			err := exchangeRows.Scan(&amount, &price)
// 			if err != nil {
// 				log.Printf("Error scanning exchange details for contract ID %s: %v", contract.Id, err)
// 				return nil, err
// 			}
// 			contract.Price += float64(amount) * price
// 		}

// 		// Check for errors during exchange row iteration
// 		if err := exchangeRows.Err(); err != nil {
// 			log.Printf("Error iterating exchange rows for contract ID %s: %v", contract.Id, err)
// 			return nil, err
// 		}

// 		// Append the contract to the response list
// 		contracts.AllContracts = append(contracts.AllContracts, &contract)
// 	}

// 	query = `SELECT COUNT(1) FROM contract`
// 	err = p.db.QueryRow(query).Scan(&count)
// 	if err != nil {
// 		return nil, err
// 	}
// 	contracts.Count = int32(count)

// 	// Check for errors during the contract row iteration
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return &contracts, nil
// }

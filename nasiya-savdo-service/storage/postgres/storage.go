package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/google/uuid"
)

type StorageStorage struct {
	db *sql.DB
}

func NewStorageStorage(db *sql.DB) *StorageStorage {
	return &StorageStorage{db: db}
}

func (p *StorageStorage) CreateStorage(req *pb.CreateStorageRequest) (*pb.StorageResponse, error) {
	id:=uuid.NewString()
	query := `
		INSERT INTO storage (id, name, user_id, created_at)
		VALUES ($1, $2, $3,now())
		RETURNING id
	`
	_,err := p.db.Exec(query, id, req.Name, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.StorageResponse{
		Message: "Storage successfully created with ID: " + id,
		Success: true,
	}, nil
}

func (p *StorageStorage) GetStorage(req *pb.StorageIdRequest) (*pb.GetStorageResponse, error) {
	query := `
		SELECT id, name, user_id
		FROM storage
		WHERE id = $1
	`
	var storage pb.GetStorageResponse
	err := p.db.QueryRow(query, req.Id).Scan(
		&storage.Id, &storage.Name, &storage.UserId,
	)
	if err != nil {
		return nil, err
	}

	return &storage, nil
}

func (p *StorageStorage) UpdateStorage(req *pb.UpdateStorageRequest) (*pb.StorageResponse, error) {
	// Initialize a map to hold the fields that need to be updated
	update := map[string]interface{}{}
	if req.Name != "" {
		update["name"] = req.Name
	}
	if req.UserId != "" {
		update["user_id"] = req.UserId
	}

	// Check if there's anything to update
	if len(update) == 0 {
		return &pb.StorageResponse{Message: "Nothing to update", Success: false}, nil
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
		UPDATE storage
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

	return &pb.StorageResponse{
		Message: "Storage successfully updated",
		Success: true,
	}, nil
}

func (p *StorageStorage) DeleteStorage(req *pb.StorageIdRequest) (*pb.StorageResponse, error) {
	query := `
		DELETE FROM storage
		WHERE id = $1
	`
	_, err := p.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.StorageResponse{
		Message: "Storage successfully deleted",
		Success: true,
	}, nil
}

func (p *StorageStorage) ListStorages(req *pb.GetAllStorageRequest) (*pb.GetAllStorageResponse, error) {
	storages := pb.GetAllStorageResponse{}
	query := `
		SELECT id, name, user_id
		FROM storage
		WHERE true
	`
	var args []interface{}
	count := 1

	if req.Name != "" {
		query += fmt.Sprintf(" AND name = $%d", count)
		args = append(args, req.Name)
		count++
	}
	if req.UserId != "" {
		query += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var storage pb.GetStorageResponse
		err := rows.Scan(
			&storage.Id, &storage.Name, &storage.UserId,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		storages.AllStorages = append(storages.AllStorages, &storage)
	}

	return &storages, nil
}

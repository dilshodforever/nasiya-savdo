package postgres

import (
	"database/sql"
	"fmt"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
)

type StorageStorage struct {
	db *sql.DB
}

func NewStorageStorage(db *sql.DB) *StorageStorage {
	return &StorageStorage{db: db}
}

func (p *StorageStorage) CreateStorage(req *pb.CreateStorageRequest) (*pb.StorageResponse, error) {
	query := `
		INSERT INTO storage (name, user_id, created_at)
		VALUES ($1, $2, now())
		RETURNING id
	`
	var id string
	err := p.db.QueryRow(query, req.Name, req.UserId).Scan(&id)
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
	query := `
		UPDATE storage
		SET name = $1, user_id = $2, updated_at = now()
		WHERE id = $3
	`
	_, err := p.db.Exec(query, req.Name, req.UserId, req.Id)
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

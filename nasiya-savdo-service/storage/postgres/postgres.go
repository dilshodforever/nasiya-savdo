package postgres

import (
	"database/sql"
	"fmt"

	"github.com/dilshodforever/nasiya-savdo/config"
	u "github.com/dilshodforever/nasiya-savdo/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db           *sql.DB
	Products     u.ProductStorage
	Contracts    u.ContractStorage
	Exchanges    u.ExchangeStorage
	Storages     u.StorageStorage
	Transactions u.TransactionStorage
}

func NewPostgresStorage() (u.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to database established!")
	return &Storage{Db: db}, nil
}

func (s *Storage) Product() u.ProductStorage {
	if s.Products == nil {
		s.Products = &ProductStorage{db: s.Db}
	}
	return s.Products
}

func (s *Storage) Contract() u.ContractStorage {
	if s.Contracts == nil {
		s.Contracts = &ContractStorage{db: s.Db}
	}
	return s.Contracts
}

func (s *Storage) Exchange() u.ExchangeStorage {
	if s.Exchanges == nil {
		s.Exchanges = &ExchangeStorage{db: s.Db}
	}
	return s.Exchanges
}

func (s *Storage) Storage() u.StorageStorage {
	if s.Storages == nil {
		s.Storages = &StorageStorage{db: s.Db}
	}
	return s.Storages
}

func (s *Storage) Transaction() u.TransactionStorage {
	if s.Transactions == nil {
		s.Transactions = &TransactionStorage{db: s.Db}
	}
	return s.Transactions
}

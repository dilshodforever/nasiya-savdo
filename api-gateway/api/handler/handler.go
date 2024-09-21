package handler

import (
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/minio/minio-go/v7"
)

type Handler struct {
	ContractService    pb.ContractServiceClient
	ExchangeService    pb.ExchangeServiceClient
	ProductService     pb.ProductServiceClient
	StorageService     pb.StorageServiceClient
	TransactionService pb.TransactionServiceClient
	MinIO              *minio.Client
	Notification       pb.NotificationtServiceClient
	Redis              InMemoryStorageI
}

func NewHandler(
	ContractService pb.ContractServiceClient,
	ExchangeService pb.ExchangeServiceClient,
	ProductService pb.ProductServiceClient,
	StorageService pb.StorageServiceClient,
	TransactionService pb.TransactionServiceClient,
	MinIO *minio.Client,
	Notification pb.NotificationtServiceClient,
	Redis InMemoryStorageI,
) *Handler {
	return &Handler{
		ContractService:    ContractService,
		ExchangeService:    ExchangeService,
		ProductService:     ProductService,
		StorageService:     StorageService,
		TransactionService: TransactionService,
		MinIO:              MinIO,
		Notification:       Notification,
		Redis:              Redis,
	}
}

package storage

import (
	"context"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
)

type InitRoot interface {
	Product() ProductStorage
	Contract() ContractStorage
	Exchange() ExchangeStorage
	Storage() StorageStorage
	Transaction() TransactionStorage
}

type MinIoRoot interface {
	MinIO() MinIOStorage
}

type ProductStorage interface {
	CreateProduct(req *pb.CreateProductRequest) (*pb.ProductResponse, error)
	GetProduct(req *pb.ProductIdRequest) (*pb.GetProductResponse, error)
	UpdateProduct(req *pb.UpdateProductRequest) (*pb.ProductResponse, error)
	DeleteProduct(req *pb.ProductIdRequest) (*pb.ProductResponse, error)
	ListProducts(req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
}

type ContractStorage interface {
	CreateContract(req *pb.CreateContractRequest) (*pb.ContractResponse, error)
	GetContract(req *pb.ContractIdRequest) (*pb.GetContractResponse, error)
	UpdateContract(req *pb.UpdateContractRequest) (*pb.ContractResponse, error)
	DeleteContract(req *pb.ContractIdRequest) (*pb.ContractResponse, error)
	ListContracts(req *pb.GetAllContractRequest) (*pb.GetAllContractResponse, error)
}

type ExchangeStorage interface {
	CreateExchange(req *pb.CreateExchangeRequest) (*pb.ExchangeResponse, error)
	GetExchange(req *pb.ExchangeIdRequest) (*pb.GetExchangeResponse, error)
	UpdateExchange(req *pb.UpdateExchangeRequest) (*pb.ExchangeResponse, error)
	DeleteExchange(req *pb.ExchangeIdRequest) (*pb.ExchangeResponse, error)
	ListExchanges(req *pb.GetAllExchangeRequest) (*pb.GetAllExchangeResponse, error)
}

type StorageStorage interface {
	CreateStorage(req *pb.CreateStorageRequest) (*pb.StorageResponse, error)
	GetStorage(req *pb.StorageIdRequest) (*pb.GetStorageResponse, error)
	UpdateStorage(req *pb.UpdateStorageRequest) (*pb.StorageResponse, error)
	DeleteStorage(req *pb.StorageIdRequest) (*pb.StorageResponse, error)
	ListStorages(req *pb.GetAllStorageRequest) (*pb.GetAllStorageResponse, error)
}

type TransactionStorage interface {
	CreateTransaction(req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error)
	GetTransaction(req *pb.TransactionIdRequest) (*pb.GetTransactionResponse, error)
	UpdateTransaction(req *pb.UpdateTransactionRequest) (*pb.TransactionResponse, error)
	DeleteTransaction(req *pb.TransactionIdRequest) (*pb.TransactionResponse, error)
	ListTransactions(req *pb.GetAllTransactionRequest) (*pb.GetAllTransactionResponse, error)
	CheckTransactions(req *pb.CheckRequest) (*pb.CheckResponse, error)
	TestNotification(req *pb.Testresponse) (*pb.Testrequest, error)
}

type MinIOStorage interface {
	UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error)
}

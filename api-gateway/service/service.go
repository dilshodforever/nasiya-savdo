package service

import (
	"fmt"
	"log"

	"github.com/dilshodforever/nasiya-savdo/api"
	"github.com/dilshodforever/nasiya-savdo/api/handler"
	"github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Connect() {
	// GRPC connection to different services
	NasiaConn, err := grpc.NewClient("nasiya-service:8087", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to ContractService: ", err.Error())
	}
	defer NasiaConn.Close()

	NotificationConn, err := grpc.NewClient("notification:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to NotificationService: ", err.Error())
	}
	defer NotificationConn.Close()

	// Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis_api:6379",
	})
	
	// MinIO connection
	minioClient, err := minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("Dior", "isakov05@", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal("Error while connecting to MinIO: ", err.Error())
	}

	// Creating gRPC clients for the services
	inmemory := handler.NewInMemoryStorage(rdb)
	notification := genprotos.NewNotificationtServiceClient(NotificationConn)
	accountClient := genprotos.NewContractServiceClient(NasiaConn)
	budgetClient := genprotos.NewExchangeServiceClient(NasiaConn)
	categoryClient := genprotos.NewProductServiceClient(NasiaConn)
	goalClient := genprotos.NewStorageServiceClient(NasiaConn)
	transactionClient := genprotos.NewTransactionServiceClient(NasiaConn)

	// Handler with MinIO and other services
	h := handler.NewHandler(accountClient, budgetClient, categoryClient, goalClient, transactionClient, minioClient, notification, inmemory)

	// Setting up the API with the handlers
	r := api.NewGin(h)
	fmt.Println("Server started on port:8080")

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}

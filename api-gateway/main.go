package main

import (
	"fmt"
	"log"

	"github.com/dilshodforever/nasiya-savdo/api"
	"github.com/dilshodforever/nasiya-savdo/api/handler"
	"github.com/dilshodforever/nasiya-savdo/genprotos" // Importing all protos from a single package

	//"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create GRPC clients for each service
	exchangeConn, err := grpc.NewClient(fmt.Sprintf("exchange%s", ":8087"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to ExchangeService: ", err.Error())
	}
	defer exchangeConn.Close()

	productConn, err := grpc.NewClient(fmt.Sprintf("product%s", ":8088"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to ProductService: ", err.Error())
	}
	defer productConn.Close()

	storageConn, err := grpc.NewClient(fmt.Sprintf("storage%s", ":8089"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to StorageService: ", err.Error())
	}
	defer storageConn.Close()

	transactionConn, err := grpc.NewClient(fmt.Sprintf("transaction%s", ":8090"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to TransactionService: ", err.Error())
	}
	defer transactionConn.Close()

	
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "redis_api:6379",
	// })


	accountClient := genprotos.NewContractServiceClient(storageConn)
	budgetClient := genprotos.NewExchangeServiceClient(storageConn)
	categoryClient := genprotos.NewProductServiceClient(storageConn)
	goalClient := genprotos.NewStorageServiceClient(storageConn)
	transactionClient := genprotos.NewTransactionServiceClient(transactionConn)
	// Create a new handler with the service clients
	h := handler.NewHandler(accountClient, budgetClient, categoryClient, goalClient, transactionClient)

	// Initialize Gin router
	r := api.NewGin(h)
	fmt.Println("Server started on port:8080")

	// Start the server
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}

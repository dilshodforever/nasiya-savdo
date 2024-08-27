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
	NasiaConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8087"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to ExchangeService: ", err.Error())
	}
	defer NasiaConn.Close()


	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "redis_api:6379",
	// })

	accountClient := genprotos.NewContractServiceClient(NasiaConn)
	budgetClient := genprotos.NewExchangeServiceClient(NasiaConn)
	categoryClient := genprotos.NewProductServiceClient(NasiaConn)
	goalClient := genprotos.NewStorageServiceClient(NasiaConn)
	transactionClient := genprotos.NewTransactionServiceClient(NasiaConn)
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

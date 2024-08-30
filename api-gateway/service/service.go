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
)

func Connect() {
	NasiaConn, err := grpc.NewClient(fmt.Sprintf("nasiya-service%s", ":8087"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to ExchangeService: ", err.Error())
	}
	defer NasiaConn.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis_api:6379",
	})
	inmemory := handler.NewInMemoryStorage(rdb)
	accountClient := genprotos.NewContractServiceClient(NasiaConn)
	budgetClient := genprotos.NewExchangeServiceClient(NasiaConn)
	categoryClient := genprotos.NewProductServiceClient(NasiaConn)
	goalClient := genprotos.NewStorageServiceClient(NasiaConn)
	transactionClient := genprotos.NewTransactionServiceClient(NasiaConn)
	minIOClient := genprotos.NewMediaServiceClient(NasiaConn)

	h := handler.NewHandler(accountClient, budgetClient, categoryClient, goalClient, transactionClient, minIOClient, inmemory)

	r := api.NewGin(h)
	fmt.Println("Server started on port:8080")

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}

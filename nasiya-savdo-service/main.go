package main

import (
	"log"
	"net"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/dilshodforever/nasiya-savdo/service"
	postgres "github.com/dilshodforever/nasiya-savdo/storage/postgres" // Update to use postgres storage package
	"google.golang.org/grpc"
)

func main() {
	// Initialize PostgreSQL connection
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	// Create a TCP listener on port 8087
	lis, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatal("Error while creating TCP listener: ", err.Error())
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register services with the gRPC server
	pb.RegisterProductServiceServer(s, service.NewProductService(db))
	pb.RegisterContractServiceServer(s, service.NewContractService(db))
	pb.RegisterExchangeServiceServer(s, service.NewExchangeService(db))
	pb.RegisterStorageServiceServer(s, service.NewStorageService(db))
	pb.RegisterTransactionServiceServer(s, service.NewTransactionService(db))

	log.Printf("Server listening at %v", lis.Addr())

	// Start the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

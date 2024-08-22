package main

import (
	"log"
	"net"

	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
	"github.com/dilshodforever/fooddalivary-food/service"
	postgres "github.com/dilshodforever/fooddalivary-food/storage/mongo"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewMongoConnecti0n()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, service.NewOrderService(db))
	pb.RegisterOrderItemServiceServer(s, service.NewOrderItemService(db))
	pb.RegisterProductServiceServer(s, service.NewProductService(db))
	
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

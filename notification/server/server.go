package server

import (
	"log"
	"net"

	pb "github.com/dilshodforever/5-oyimtixon/genprotos"
	"github.com/dilshodforever/5-oyimtixon/kafka"
	"github.com/dilshodforever/5-oyimtixon/service"
	postgres "github.com/dilshodforever/5-oyimtixon/storage/mongo"
	"google.golang.org/grpc"
)

func Connection() net.Listener {

	db, err := postgres.NewMongoConnection()
	if err != nil {
		log.Fatal("Error while connecting to DB: ", err.Error())
	}

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal("Error while starting TCP listener: ", err.Error())
	}
	notifications := service.NewNotificationService(db)
	s := grpc.NewServer()
	pb.RegisterNotificationtServiceServer(s, notifications)

	brokers := []string{"kafka:9092"}

	kcm := kafka.NewKafkaConsumerManager()
	appService := service.NewNotificationService(db)

	if err := kcm.RegisterConsumer(brokers, "NasiaSendNatification", "root", kafka.SendNotification(appService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'create-job_application' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return lis
}

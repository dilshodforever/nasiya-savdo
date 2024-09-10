package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dilshodforever/nasiya-savdo/api"
	"github.com/dilshodforever/nasiya-savdo/api/handler"
	"github.com/dilshodforever/nasiya-savdo/config"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/dilshodforever/nasiya-savdo/kafka"
	"github.com/dilshodforever/nasiya-savdo/service"
	"github.com/dilshodforever/nasiya-savdo/storage/postgres"
	r "github.com/dilshodforever/nasiya-savdo/storage/redis"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cnf := config.Load()
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	liss, err := net.Listen("tcp", cnf.GrpcPort)
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	udb := service.NewUserService(db)

	cus := kafka.NewKafkaConsumerManager()
	broker := []string{"kafka:9092"}
	err = cus.RegisterConsumer(broker, "user-create", "user", kafka.UserCreateHandler(udb))
	if err != nil {
		log.Fatal("Error while create user: ", err.Error())
	}
	err = cus.RegisterConsumer(broker, "user-update", "user", kafka.UserUpdateHandler(udb))
	if err != nil {
		log.Fatal("Error while update user: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service.NewUserService(db))
	log.Printf("server listening at %v", liss.Addr())
	go func() {
		if err := s.Serve(liss); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	userConn, err := grpc.NewClient(fmt.Sprintf("auth-service%s", cnf.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NewClient: ", err.Error())
	}
	defer userConn.Close()

	client := redis.NewClient(&redis.Options{
		Addr: cnf.RedisHost + cnf.RedisPort,
	})

	rdb := r.NewInMemoryStorage(client)
	pr, err := kafka.NewKafkaProducer(broker)
	if err != nil {
		log.Fatal("Error while producer: ", err)
	}

	us := pb.NewUserServiceClient(userConn)
	h := handler.NewHandler(us, rdb, pr)
	r := api.NewGin(h)

	fmt.Println("Server started on port:", cnf.HTTPPort)

	err = r.Run(cnf.HTTPPort)
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}

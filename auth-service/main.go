package main

import (
	"fmt"
	"log"
	"net"

	"github.com/redis/go-redis/v9"
	"gitlab.com/lingualeap/auth/api"
	"gitlab.com/lingualeap/auth/api/handler"
	"gitlab.com/lingualeap/auth/config"
	pb "gitlab.com/lingualeap/auth/genprotos/users"
	"gitlab.com/lingualeap/auth/kafka"
	"gitlab.com/lingualeap/auth/service"
	"gitlab.com/lingualeap/auth/storage/postgres"
	r "gitlab.com/lingualeap/auth/storage/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cnf := config.Load()
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	liss, err := net.Listen("tcp", cnf.GrpcUserPort)
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}
	udb := service.NewUserService(db)

	cus := kafka.NewKafkaConsumerManager()
	broker := []string{"kafka:9092"}
	cus.RegisterConsumer(broker, "user", "u", kafka.UserCreateHandler(udb))

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service.NewUserService(db))
	log.Printf("server listening at %v", liss.Addr())
	go func() {
		if err := s.Serve(liss); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	userConn, err := grpc.NewClient(fmt.Sprintf("auth-ll%s", cnf.GrpcUserPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
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

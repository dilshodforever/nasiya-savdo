package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/dilshodforever/5-oyimtixon/config"
	u "github.com/dilshodforever/5-oyimtixon/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	Db    *mongo.Database
	Accs  u.NotificationService
	
}

func NewMongoConnection() (u.InitRoot, error) {
	config:=config.Load()
	clientOptions := options.Client().ApplyURI(config.MongoDBConnection)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error: Couldn't connect to the database.", err)
	}

	fmt.Println("Connected to MongoDB!")

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("notifacitions")

	return &MongoStorage{Db: db}, err
}

func (s *MongoStorage) Notification() u.NotificationService {
	if s.Accs == nil {
		s.Accs = &AccountService{db: s.Db}
	}
	return s.Accs
}

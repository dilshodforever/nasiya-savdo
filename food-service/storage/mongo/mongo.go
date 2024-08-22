package postgres

import (
	"context"
	"fmt"
	"log"

	u "github.com/dilshodforever/fooddalivary-food/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	Db         *mongo.Database
	Products   u.ProductStorage
	Orders     u.OrderStorage
	OrderItems u.OrderItemStorage
}

func NewMongoConnecti0n() (u.InitRoot, error) {
	// 	uri := fmt.Sprintf("mongodb://%s:%d",
	//     "mongo",
	//     27017,
	//   )
	// 	clientOptions := options.Client().ApplyURI(uri).
	// 	SetAuth(options.Credential{Username: "dilshod", Password: "root"})

	// client, err := mongo.Connect(context.Background(), clientOptions)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	clientOptions := options.Client().ApplyURI("mongodb+srv://dilshod:2514@cluster0.klxv3df.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

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

	db := client.Database("fooddeleviry")

	return &MongoStorage{Db: db}, err
}

func (s *MongoStorage) Product() u.ProductStorage {
	if s.Products == nil {
		s.Products = &FoodStorage{s.Db}
	}
	return s.Products
}

func (s *MongoStorage) Order() u.OrderStorage {
	if s.Orders == nil {
		s.Orders = &FoodStorage{s.Db}
	}
	return s.Orders
}

func (s *MongoStorage) OrderItem() u.OrderItemStorage {
	if s.OrderItems == nil {
		s.OrderItems = &FoodStorage{s.Db}
	}
	return s.OrderItems
}

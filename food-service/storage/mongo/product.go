package postgres

import (
	ctx "context"
	"log"
	"time"

	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodStorage struct {
	db *mongo.Database
}

func NewProductStorage(db *mongo.Database) *FoodStorage {
	return &FoodStorage{db: db}
}

// CreateProduct inserts a new product into the database
func (f *FoodStorage) CreateProduct(req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	coll := f.db.Collection("products")
	id:=uuid.NewString()
	_, err := coll.InsertOne(ctx.Background(), bson.M{
		"ProductId":	id,
		"name":         req.Name,
		"description":  req.Description,
		"price":        req.Price,
		"ImageUrl":     req.ImageUrl,
		"CreatedAt":    time.Now().String(),
		"UpdatedAt":    time.Now().String(),
	})
	if err != nil {
		log.Printf("Failed to create product: %v", err)
		return &pb.ProductResponse{
			Message: "Failed to create product",
			Success: false,
		}, err
	}

	return &pb.ProductResponse{
		Message: "Product created successfully",
		Success: true,
	}, nil
}

// GetProduct retrieves a product by its ID
func (f *FoodStorage) GetProduct(req *pb.ProductIdRequest) (*pb.GetProductResponse, error) {
	coll := f.db.Collection("products")
	var product pb.GetProductResponse
	err := coll.FindOne(ctx.Background(), bson.M{"ProductId": req.ProductId}).Decode(&product)
	if err != nil {
		log.Printf("Failed to get product: %v", err)
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product
func (f *FoodStorage) UpdateProduct(req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	coll := f.db.Collection("products")
	_, err := coll.UpdateOne(ctx.Background(), bson.M{"ProductId": req.ProductId}, bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"description": req.Description,
			"price":       req.Price,
			"ImageUrl":    req.ImageUrl,
			"UpdatedAt":   time.Now(),
		},
	})
	if err != nil {
		log.Printf("Failed to update product: %v", err)
		return &pb.ProductResponse{
			Message: "Failed to update product",
			Success: false,
		}, err
	}

	return &pb.ProductResponse{
		Message: "Product updated successfully",
		Success: true,
	}, nil
}

// DeleteProduct removes a product by its ID
func (f *FoodStorage) DeleteProduct(req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	coll := f.db.Collection("products")
	_, err := coll.DeleteOne(ctx.Background(), bson.M{"ProductId": req.ProductId})
	if err != nil {
		log.Printf("Failed to delete product: %v", err)
		return &pb.ProductResponse{
			Message: "Failed to delete product",
			Success: false,
		}, err
	}

	return &pb.ProductResponse{
		Message: "Product deleted successfully",
		Success: true,
	}, nil
}

// ListProducts lists all products
func (f *FoodStorage) ListProducts(req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	coll := f.db.Collection("products")

	// Build the filter query based on the fields provided in the request
	filter := bson.M{}
	if req.ProductId != "" {
		filter["product_id"] = req.ProductId
	}
	if req.Name != "" {
		filter["name"] = req.Name
	}
	if req.Description != "" {
		filter["description"] = req.Description
	}
	if req.Price != 0 {
		filter["price"] = req.Price
	}

	// Execute the query with the filter
	cursor, err := coll.Find(ctx.Background(), filter)
	if err != nil {
		log.Printf("Failed to list products: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx.Background())

	// Check if any documents were found
	if cursor.RemainingBatchLength() == 0 {
		log.Print("No products found matching the filter criteria")
		return &pb.GetAllProductResponse{
			Allproducts: []*pb.GetProductResponse{},
			Message:     "No products found",
		}, nil
	}

	var products []*pb.GetProductResponse
	for cursor.Next(ctx.Background()) {
		var product pb.GetProductResponse
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Failed to decode product: %v", err)
			return nil, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &pb.GetAllProductResponse{
		Allproducts: products,
		Message:     "Products retrieved successfully",
	}, nil
}
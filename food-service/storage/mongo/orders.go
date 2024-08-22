package postgres

import (
	"context"
	"log"
	"time"

	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (g *FoodStorage) CreateOrder(req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	coll := g.db.Collection("orders")
	allitems, err:=g.ListOrderItems(&pb.GetByUseridItems{UserId: req.UserId})
	if err != nil {
		log.Printf("Failed to get OrderItems: %v", err)
		return nil, err
	}
	id := uuid.NewString()
	order := bson.M{
		"order_id":        id,
		"UserId":          req.UserId,
		"CourierId":       req.CourierId,
		"status":          "Tayyorlanmoqda",
		"TotalAmount":     allitems.TotalAmount,
		"DeliveryAddress": req.DeliveryAddress,
		"time":            req.Time,
		"products":		   allitems.Products,
		"created_at":      time.Now(),
		"deleted_at":      0,
		"updated_at":      time.Now(),
	}

	_, err = coll.InsertOne(context.Background(), order)
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return nil, err
	}
	return &pb.CreateOrderResponse{OrderId: id, Message: "Order created successfully"}, nil
}

func (g *FoodStorage) GetOrderById(req *pb.GetByIdRequest) (*pb.GetAllOrders, error) {
	coll := g.db.Collection("orders")
	var order pb.GetAllOrders
	err := coll.FindOne(context.Background(), bson.M{"order_id": req.OrderId}).Decode(&order)
	if err != nil {
		log.Printf("Failed to get order: %v", err)
		return nil, err
	}

	return &order, nil
}

func (g *FoodStorage) UpdateOrder(req *pb.UpdateOrderRequest) (*pb.UpdateStatusResponse, error) {
	coll := g.db.Collection("orders")
	_, err := coll.UpdateOne(context.Background(), bson.M{"order_id": req.OrderId}, bson.M{
		"$set": bson.M{
			"DeliveryAddress": req.DeliveryAddress,
		},
	})
	if err != nil {
		log.Printf("Failed to update order: %v", err)
		return &pb.UpdateStatusResponse{Message: "Error while update order", Succes: false}, err
	}
	return &pb.UpdateStatusResponse{Message: "Order updated successfully", Succes: true}, nil
}

func (g *FoodStorage) DeleteOrder(req *pb.DeleteOrdersByidRequest) (*pb.DeleteOrdersByidResponse, error) {
	coll := g.db.Collection("orders")
	_, err := coll.DeleteOne(context.Background(), bson.M{"order_id": req.OrderId})
	if err != nil {
		log.Printf("Failed to delete order: %v", err)
		return &pb.DeleteOrdersByidResponse{Message: "Error while delete order", Succes: false}, err
	}
	return &pb.DeleteOrdersByidResponse{Message: "Order deleted successfully", Succes: true}, nil
}

func (g *FoodStorage) ListOrders(req *pb.GetAllRequest) (*pb.GetAllOrdersList, error) {
	coll := g.db.Collection("orders")

	filter := bson.M{}
	if req.OrderId != "" {
		filter["OrderId"] = req.OrderId
	}
	if req.UserId != "" {
		filter["UserId"] = req.UserId
	}
	if req.CourierId != "" {
		filter["CourierId"] = req.CourierId
	}
	if req.Status != "" {
		filter["Status"] = req.Status
	}
	if req.TotalAmount != 0 {
		filter["TotalAmount"] = req.TotalAmount
	}
	if req.DeliveryAddress != "" {
		filter["DeliveryAddress"] = req.DeliveryAddress
	}

	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Failed to list orders: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders pb.GetAllOrdersList
	for cursor.Next(context.Background()) {
		var order pb.GetAllOrders
		if err := cursor.Decode(&order); err != nil {
			log.Printf("Failed to decode order: %v", err)
			return nil, err
		}
		orders.Orders = append(orders.Orders, &order)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &orders, nil
}


func (g *FoodStorage) UpdateStatus(req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	coll := g.db.Collection("orders")
	_, err := coll.UpdateOne(context.Background(), bson.M{"order_id": req.OrderId}, bson.M{
		"$set": bson.M{
			"status": req.Status,
		},
	})
	if err != nil {
		log.Printf("Failed to update order status: %v", err)
		return nil, err
	}

	return &pb.UpdateStatusResponse{Message: "Order status updated successfully"}, nil
}







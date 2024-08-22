package postgres

import (
	ctx "context"
	"log"

	protos "github.com/dilshodforever/fooddalivary-food/genprotos"
	"go.mongodb.org/mongo-driver/bson"
)

func (g *FoodStorage) AddItems(req *protos.AddItemsRequest) (*protos.AddItemsResponse, error) {

	product, err := g.GetProduct(&protos.ProductIdRequest{ProductId: req.ProductId})
	if err != nil {
		log.Printf("Failed to get product details: %v", err)
		return nil, err
	}

	res := &protos.AddItemsResponse{
		Products: []*protos.GetProductResponse{product},
	}

	coll := g.db.Collection("order_items")

	_, err = coll.UpdateOne(
		ctx.TODO(),
		bson.M{"UserId": req.UserId},
		bson.M{
			"$push": bson.M{"products": bson.M{"$each": res.Products}},
		},
	)
	if err != nil {
		log.Printf("Failed to append to order item products: %v", err)
		return nil, err
	}

	var updatedOrder protos.GetAllItems
	err = coll.FindOne(ctx.TODO(), bson.M{"UserId": req.UserId}).Decode(&updatedOrder)
	if err != nil {
		log.Printf("Failed to retrieve updated order item: %v", err)
		return nil, err
	}

	totalPrice := 0.0
	totalQuantity := int32(0)
	for _, prod := range updatedOrder.Products {
		totalPrice += prod.Price
		totalQuantity++
	}

	_, err = coll.UpdateOne(
		ctx.TODO(),
		bson.M{"UserId": req.UserId},
		bson.M{
			"$set": bson.M{"TotalAmount": totalPrice, "totalQuantity": totalQuantity},
		},
	)
	if err != nil {
		log.Printf("Failed when updated quantity and amount: %v", err)
		return nil, err
	}
	res.TotalPrice = totalPrice
	res.Quantity = totalQuantity

	return res, nil
}

func (g *FoodStorage) DeleteItems(req *protos.DeleItemsRequest) (*protos.DeleteProductResponse, error) {
	coll := g.db.Collection("order_items")
	_, err := coll.DeleteOne(ctx.Background(), bson.M{"products.productid": req.ProductId})
	if err != nil {
		log.Printf("Failed to delete order item: %v", err)
		return nil, err
	}

	return &protos.DeleteProductResponse{
		Message: "Item deleted successfully",
		Success: true,
	}, nil
}

func (g *FoodStorage) ListOrderItems(req *protos.GetByUseridItems) (*protos.GetAllItems, error) {
	coll := g.db.Collection("order_items")
	response := protos.GetAllItems{}
	err := coll.FindOne(ctx.Background(), bson.M{"UserId": req.UserId}).Decode(&response)
	if err != nil {
		log.Printf("Failed to list order items: %v", err)
		return nil, err
	}
	return &response, nil
}

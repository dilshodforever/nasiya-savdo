package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
	s "github.com/dilshodforever/fooddalivary-food/storage"
)

type OrderItemService struct {
	stg s.InitRoot
	pb.UnimplementedOrderItemServiceServer
}

func NewOrderItemService(stg s.InitRoot) *OrderItemService {
	return &OrderItemService{stg: stg}
}

// AddItems adds items to an order
func (s *OrderItemService) AddItems(ctx context.Context, req *pb.AddItemsRequest) (*pb.AddItemsResponse, error) {
	resp, err := s.stg.OrderItem().AddItems(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

// DeleteItems removes an item from an order
func (s *OrderItemService) DeleteItems(ctx context.Context, req *pb.DeleItemsRequest) (*pb.DeleteProductResponse, error) {
	resp, err := s.stg.OrderItem().DeleteItems(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

// ListOrderItems lists all items in an order by user ID
func (s *OrderItemService) ListOrderItems(ctx context.Context, req *pb.GetByUseridItems) (*pb.GetAllItems, error) {
	resp, err := s.stg.OrderItem().ListOrderItems(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	
	return resp, nil
}

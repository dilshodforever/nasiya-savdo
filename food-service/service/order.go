package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
	s "github.com/dilshodforever/fooddalivary-food/storage"
)

type OrderService struct {
	stg s.InitRoot
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(stg s.InitRoot) *OrderService {
	return &OrderService{stg: stg}
}

// CreateOrder handles the creation of a new order
func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	resp, err := s.stg.Order().CreateOrder(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

// GetOrderById retrieves an order by its ID

func (s *OrderService) GetOrderByid(ctx context.Context, req *pb.GetByIdRequest) (*pb.GetAllOrders, error) {
	resp, err := s.stg.Order().GetOrderById(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

// UpdateOrder updates an existing order

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateStatusResponse, error) {
	resp, err := s.stg.Order().UpdateOrder(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

// DeleteOrder removes an order by its ID
func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.DeleteOrdersByidRequest) (*pb.DeleteOrdersByidResponse, error) {
	resp, err := s.stg.Order().DeleteOrder(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

// ListOrders lists all orders
func (s *OrderService) ListOrders(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllOrdersList, error) {
	resp, err := s.stg.Order().ListOrders(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

// UpdateStatus updates the status of an order
func (s *OrderService) UpdateStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	resp, err := s.stg.Order().UpdateStatus(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

package storage

import (
	pb "github.com/dilshodforever/fooddalivary-food/genprotos"
)


type InitRoot interface {
	Product() ProductStorage
	Order() OrderStorage
	OrderItem() OrderItemStorage
}

type ProductStorage interface {
	CreateProduct(req *pb.CreateProductRequest) (*pb.ProductResponse, error)
	GetProduct(req *pb.ProductIdRequest) (*pb.GetProductResponse, error)
	UpdateProduct(req *pb.UpdateProductRequest) (*pb.ProductResponse, error)
	DeleteProduct(req *pb.ProductIdRequest) (*pb.ProductResponse, error)
	ListProducts(req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
}

type OrderStorage interface {
	CreateOrder(req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
	GetOrderById(req *pb.GetByIdRequest) (*pb.GetAllOrders, error)
	UpdateOrder(req *pb.UpdateOrderRequest) (*pb.UpdateStatusResponse, error)
	DeleteOrder(req *pb.DeleteOrdersByidRequest) (*pb.DeleteOrdersByidResponse, error)
	ListOrders(req *pb.GetAllRequest) (*pb.GetAllOrdersList, error)
	UpdateStatus(req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error)
}

type OrderItemStorage interface {
	AddItems(req *pb.AddItemsRequest) (*pb.AddItemsResponse, error)
	DeleteItems(req *pb.DeleItemsRequest) (*pb.DeleteProductResponse, error)
	ListOrderItems(req *pb.GetByUseridItems) (*pb.GetAllItems, error)
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: exchange.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ExchangeService_CreateExchange_FullMethodName = "/protos.ExchangeService/CreateExchange"
	ExchangeService_GetExchange_FullMethodName    = "/protos.ExchangeService/GetExchange"
	ExchangeService_UpdateExchange_FullMethodName = "/protos.ExchangeService/UpdateExchange"
	ExchangeService_DeleteExchange_FullMethodName = "/protos.ExchangeService/DeleteExchange"
	ExchangeService_ListExchanges_FullMethodName  = "/protos.ExchangeService/ListExchanges"
)

// ExchangeServiceClient is the client API for ExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangeServiceClient interface {
	CreateExchange(ctx context.Context, in *CreateExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error)
	GetExchange(ctx context.Context, in *ExchangeIdRequest, opts ...grpc.CallOption) (*GetExchangeResponse, error)
	UpdateExchange(ctx context.Context, in *UpdateExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error)
	DeleteExchange(ctx context.Context, in *ExchangeIdRequest, opts ...grpc.CallOption) (*ExchangeResponse, error)
	ListExchanges(ctx context.Context, in *GetAllExchangeRequest, opts ...grpc.CallOption) (*GetAllExchangeResponse, error)
}

type exchangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeServiceClient(cc grpc.ClientConnInterface) ExchangeServiceClient {
	return &exchangeServiceClient{cc}
}

func (c *exchangeServiceClient) CreateExchange(ctx context.Context, in *CreateExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeResponse)
	err := c.cc.Invoke(ctx, ExchangeService_CreateExchange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) GetExchange(ctx context.Context, in *ExchangeIdRequest, opts ...grpc.CallOption) (*GetExchangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetExchangeResponse)
	err := c.cc.Invoke(ctx, ExchangeService_GetExchange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) UpdateExchange(ctx context.Context, in *UpdateExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeResponse)
	err := c.cc.Invoke(ctx, ExchangeService_UpdateExchange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) DeleteExchange(ctx context.Context, in *ExchangeIdRequest, opts ...grpc.CallOption) (*ExchangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeResponse)
	err := c.cc.Invoke(ctx, ExchangeService_DeleteExchange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) ListExchanges(ctx context.Context, in *GetAllExchangeRequest, opts ...grpc.CallOption) (*GetAllExchangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllExchangeResponse)
	err := c.cc.Invoke(ctx, ExchangeService_ListExchanges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServiceServer is the server API for ExchangeService service.
// All implementations must embed UnimplementedExchangeServiceServer
// for forward compatibility
type ExchangeServiceServer interface {
	CreateExchange(context.Context, *CreateExchangeRequest) (*ExchangeResponse, error)
	GetExchange(context.Context, *ExchangeIdRequest) (*GetExchangeResponse, error)
	UpdateExchange(context.Context, *UpdateExchangeRequest) (*ExchangeResponse, error)
	DeleteExchange(context.Context, *ExchangeIdRequest) (*ExchangeResponse, error)
	ListExchanges(context.Context, *GetAllExchangeRequest) (*GetAllExchangeResponse, error)
	mustEmbedUnimplementedExchangeServiceServer()
}

// UnimplementedExchangeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExchangeServiceServer struct {
}

func (UnimplementedExchangeServiceServer) CreateExchange(context.Context, *CreateExchangeRequest) (*ExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExchange not implemented")
}
func (UnimplementedExchangeServiceServer) GetExchange(context.Context, *ExchangeIdRequest) (*GetExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchange not implemented")
}
func (UnimplementedExchangeServiceServer) UpdateExchange(context.Context, *UpdateExchangeRequest) (*ExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExchange not implemented")
}
func (UnimplementedExchangeServiceServer) DeleteExchange(context.Context, *ExchangeIdRequest) (*ExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExchange not implemented")
}
func (UnimplementedExchangeServiceServer) ListExchanges(context.Context, *GetAllExchangeRequest) (*GetAllExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExchanges not implemented")
}
func (UnimplementedExchangeServiceServer) mustEmbedUnimplementedExchangeServiceServer() {}

// UnsafeExchangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeServiceServer will
// result in compilation errors.
type UnsafeExchangeServiceServer interface {
	mustEmbedUnimplementedExchangeServiceServer()
}

func RegisterExchangeServiceServer(s grpc.ServiceRegistrar, srv ExchangeServiceServer) {
	s.RegisterService(&ExchangeService_ServiceDesc, srv)
}

func _ExchangeService_CreateExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).CreateExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_CreateExchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).CreateExchange(ctx, req.(*CreateExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_GetExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetExchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetExchange(ctx, req.(*ExchangeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_UpdateExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).UpdateExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_UpdateExchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).UpdateExchange(ctx, req.(*UpdateExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_DeleteExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).DeleteExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_DeleteExchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).DeleteExchange(ctx, req.(*ExchangeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_ListExchanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).ListExchanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_ListExchanges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).ListExchanges(ctx, req.(*GetAllExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExchangeService_ServiceDesc is the grpc.ServiceDesc for ExchangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExchangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ExchangeService",
	HandlerType: (*ExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExchange",
			Handler:    _ExchangeService_CreateExchange_Handler,
		},
		{
			MethodName: "GetExchange",
			Handler:    _ExchangeService_GetExchange_Handler,
		},
		{
			MethodName: "UpdateExchange",
			Handler:    _ExchangeService_UpdateExchange_Handler,
		},
		{
			MethodName: "DeleteExchange",
			Handler:    _ExchangeService_DeleteExchange_Handler,
		},
		{
			MethodName: "ListExchanges",
			Handler:    _ExchangeService_ListExchanges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exchange.proto",
}

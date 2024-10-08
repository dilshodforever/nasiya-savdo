// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: contract.proto

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
	ContractService_CreateContract_FullMethodName = "/protos.ContractService/CreateContract"
	ContractService_GetContract_FullMethodName    = "/protos.ContractService/GetContract"
	ContractService_UpdateContract_FullMethodName = "/protos.ContractService/UpdateContract"
	ContractService_DeleteContract_FullMethodName = "/protos.ContractService/DeleteContract"
	ContractService_ListContracts_FullMethodName  = "/protos.ContractService/ListContracts"
)

// ContractServiceClient is the client API for ContractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractServiceClient interface {
	CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*ContractResponse, error)
	GetContract(ctx context.Context, in *ContractIdRequest, opts ...grpc.CallOption) (*GetContractResponse, error)
	UpdateContract(ctx context.Context, in *UpdateContractRequest, opts ...grpc.CallOption) (*ContractResponse, error)
	DeleteContract(ctx context.Context, in *ContractIdRequest, opts ...grpc.CallOption) (*ContractResponse, error)
	ListContracts(ctx context.Context, in *GetAllContractRequest, opts ...grpc.CallOption) (*GetAllContractResponse, error)
}

type contractServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContractServiceClient(cc grpc.ClientConnInterface) ContractServiceClient {
	return &contractServiceClient{cc}
}

func (c *contractServiceClient) CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*ContractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContractResponse)
	err := c.cc.Invoke(ctx, ContractService_CreateContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) GetContract(ctx context.Context, in *ContractIdRequest, opts ...grpc.CallOption) (*GetContractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContractResponse)
	err := c.cc.Invoke(ctx, ContractService_GetContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) UpdateContract(ctx context.Context, in *UpdateContractRequest, opts ...grpc.CallOption) (*ContractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContractResponse)
	err := c.cc.Invoke(ctx, ContractService_UpdateContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) DeleteContract(ctx context.Context, in *ContractIdRequest, opts ...grpc.CallOption) (*ContractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContractResponse)
	err := c.cc.Invoke(ctx, ContractService_DeleteContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) ListContracts(ctx context.Context, in *GetAllContractRequest, opts ...grpc.CallOption) (*GetAllContractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllContractResponse)
	err := c.cc.Invoke(ctx, ContractService_ListContracts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractServiceServer is the server API for ContractService service.
// All implementations must embed UnimplementedContractServiceServer
// for forward compatibility
type ContractServiceServer interface {
	CreateContract(context.Context, *CreateContractRequest) (*ContractResponse, error)
	GetContract(context.Context, *ContractIdRequest) (*GetContractResponse, error)
	UpdateContract(context.Context, *UpdateContractRequest) (*ContractResponse, error)
	DeleteContract(context.Context, *ContractIdRequest) (*ContractResponse, error)
	ListContracts(context.Context, *GetAllContractRequest) (*GetAllContractResponse, error)
	mustEmbedUnimplementedContractServiceServer()
}

// UnimplementedContractServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContractServiceServer struct {
}

func (UnimplementedContractServiceServer) CreateContract(context.Context, *CreateContractRequest) (*ContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContract not implemented")
}
func (UnimplementedContractServiceServer) GetContract(context.Context, *ContractIdRequest) (*GetContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContract not implemented")
}
func (UnimplementedContractServiceServer) UpdateContract(context.Context, *UpdateContractRequest) (*ContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContract not implemented")
}
func (UnimplementedContractServiceServer) DeleteContract(context.Context, *ContractIdRequest) (*ContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContract not implemented")
}
func (UnimplementedContractServiceServer) ListContracts(context.Context, *GetAllContractRequest) (*GetAllContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContracts not implemented")
}
func (UnimplementedContractServiceServer) mustEmbedUnimplementedContractServiceServer() {}

// UnsafeContractServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractServiceServer will
// result in compilation errors.
type UnsafeContractServiceServer interface {
	mustEmbedUnimplementedContractServiceServer()
}

func RegisterContractServiceServer(s grpc.ServiceRegistrar, srv ContractServiceServer) {
	s.RegisterService(&ContractService_ServiceDesc, srv)
}

func _ContractService_CreateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).CreateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_CreateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).CreateContract(ctx, req.(*CreateContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_GetContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContractIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).GetContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_GetContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).GetContract(ctx, req.(*ContractIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_UpdateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).UpdateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_UpdateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).UpdateContract(ctx, req.(*UpdateContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_DeleteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContractIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).DeleteContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_DeleteContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).DeleteContract(ctx, req.(*ContractIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_ListContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).ListContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractService_ListContracts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).ListContracts(ctx, req.(*GetAllContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContractService_ServiceDesc is the grpc.ServiceDesc for ContractService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContractService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ContractService",
	HandlerType: (*ContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContract",
			Handler:    _ContractService_CreateContract_Handler,
		},
		{
			MethodName: "GetContract",
			Handler:    _ContractService_GetContract_Handler,
		},
		{
			MethodName: "UpdateContract",
			Handler:    _ContractService_UpdateContract_Handler,
		},
		{
			MethodName: "DeleteContract",
			Handler:    _ContractService_DeleteContract_Handler,
		},
		{
			MethodName: "ListContracts",
			Handler:    _ContractService_ListContracts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract.proto",
}

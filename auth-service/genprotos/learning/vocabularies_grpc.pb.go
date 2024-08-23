// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: vocabularies.proto

package learning

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VocabularyService_Create_FullMethodName  = "/vocabularies.VocabularyService/Create"
	VocabularyService_Update_FullMethodName  = "/vocabularies.VocabularyService/Update"
	VocabularyService_Delete_FullMethodName  = "/vocabularies.VocabularyService/Delete"
	VocabularyService_GetById_FullMethodName = "/vocabularies.VocabularyService/GetById"
	VocabularyService_GetAll_FullMethodName  = "/vocabularies.VocabularyService/GetAll"
)

// VocabularyServiceClient is the client API for VocabularyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VocabularyServiceClient interface {
	Create(ctx context.Context, in *VocabularyReq, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *Vocabulary, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Vocabulary, error)
	GetAll(ctx context.Context, in *VocabularyFilter, opts ...grpc.CallOption) (*AllVocabluaries, error)
}

type vocabularyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVocabularyServiceClient(cc grpc.ClientConnInterface) VocabularyServiceClient {
	return &vocabularyServiceClient{cc}
}

func (c *vocabularyServiceClient) Create(ctx context.Context, in *VocabularyReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, VocabularyService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) Update(ctx context.Context, in *Vocabulary, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, VocabularyService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, VocabularyService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Vocabulary, error) {
	out := new(Vocabulary)
	err := c.cc.Invoke(ctx, VocabularyService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) GetAll(ctx context.Context, in *VocabularyFilter, opts ...grpc.CallOption) (*AllVocabluaries, error) {
	out := new(AllVocabluaries)
	err := c.cc.Invoke(ctx, VocabularyService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VocabularyServiceServer is the server API for VocabularyService service.
// All implementations must embed UnimplementedVocabularyServiceServer
// for forward compatibility
type VocabularyServiceServer interface {
	Create(context.Context, *VocabularyReq) (*Void, error)
	Update(context.Context, *Vocabulary) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	GetById(context.Context, *ById) (*Vocabulary, error)
	GetAll(context.Context, *VocabularyFilter) (*AllVocabluaries, error)
	mustEmbedUnimplementedVocabularyServiceServer()
}

// UnimplementedVocabularyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVocabularyServiceServer struct {
}

func (UnimplementedVocabularyServiceServer) Create(context.Context, *VocabularyReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVocabularyServiceServer) Update(context.Context, *Vocabulary) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVocabularyServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedVocabularyServiceServer) GetById(context.Context, *ById) (*Vocabulary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedVocabularyServiceServer) GetAll(context.Context, *VocabularyFilter) (*AllVocabluaries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedVocabularyServiceServer) mustEmbedUnimplementedVocabularyServiceServer() {}

// UnsafeVocabularyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VocabularyServiceServer will
// result in compilation errors.
type UnsafeVocabularyServiceServer interface {
	mustEmbedUnimplementedVocabularyServiceServer()
}

func RegisterVocabularyServiceServer(s grpc.ServiceRegistrar, srv VocabularyServiceServer) {
	s.RegisterService(&VocabularyService_ServiceDesc, srv)
}

func _VocabularyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VocabularyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).Create(ctx, req.(*VocabularyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vocabulary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).Update(ctx, req.(*Vocabulary))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VocabularyFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).GetAll(ctx, req.(*VocabularyFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// VocabularyService_ServiceDesc is the grpc.ServiceDesc for VocabularyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VocabularyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vocabularies.VocabularyService",
	HandlerType: (*VocabularyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VocabularyService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VocabularyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VocabularyService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _VocabularyService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _VocabularyService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vocabularies.proto",
}
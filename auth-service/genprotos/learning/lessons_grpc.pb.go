// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: lessons.proto

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
	LessonService_Create_FullMethodName  = "/lessons.LessonService/Create"
	LessonService_Update_FullMethodName  = "/lessons.LessonService/Update"
	LessonService_Delete_FullMethodName  = "/lessons.LessonService/Delete"
	LessonService_GetById_FullMethodName = "/lessons.LessonService/GetById"
	LessonService_GetAll_FullMethodName  = "/lessons.LessonService/GetAll"
)

// LessonServiceClient is the client API for LessonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LessonServiceClient interface {
	Create(ctx context.Context, in *LessonReq, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *Lesson, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Lesson, error)
	GetAll(ctx context.Context, in *LessonFilter, opts ...grpc.CallOption) (*AllLessons, error)
}

type lessonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLessonServiceClient(cc grpc.ClientConnInterface) LessonServiceClient {
	return &lessonServiceClient{cc}
}

func (c *lessonServiceClient) Create(ctx context.Context, in *LessonReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, LessonService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonServiceClient) Update(ctx context.Context, in *Lesson, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, LessonService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, LessonService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Lesson, error) {
	out := new(Lesson)
	err := c.cc.Invoke(ctx, LessonService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonServiceClient) GetAll(ctx context.Context, in *LessonFilter, opts ...grpc.CallOption) (*AllLessons, error) {
	out := new(AllLessons)
	err := c.cc.Invoke(ctx, LessonService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LessonServiceServer is the server API for LessonService service.
// All implementations must embed UnimplementedLessonServiceServer
// for forward compatibility
type LessonServiceServer interface {
	Create(context.Context, *LessonReq) (*Void, error)
	Update(context.Context, *Lesson) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	GetById(context.Context, *ById) (*Lesson, error)
	GetAll(context.Context, *LessonFilter) (*AllLessons, error)
	mustEmbedUnimplementedLessonServiceServer()
}

// UnimplementedLessonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLessonServiceServer struct {
}

func (UnimplementedLessonServiceServer) Create(context.Context, *LessonReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLessonServiceServer) Update(context.Context, *Lesson) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLessonServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLessonServiceServer) GetById(context.Context, *ById) (*Lesson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedLessonServiceServer) GetAll(context.Context, *LessonFilter) (*AllLessons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedLessonServiceServer) mustEmbedUnimplementedLessonServiceServer() {}

// UnsafeLessonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LessonServiceServer will
// result in compilation errors.
type UnsafeLessonServiceServer interface {
	mustEmbedUnimplementedLessonServiceServer()
}

func RegisterLessonServiceServer(s grpc.ServiceRegistrar, srv LessonServiceServer) {
	s.RegisterService(&LessonService_ServiceDesc, srv)
}

func _LessonService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LessonService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).Create(ctx, req.(*LessonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lesson)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LessonService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).Update(ctx, req.(*Lesson))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LessonService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LessonService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LessonService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonServiceServer).GetAll(ctx, req.(*LessonFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// LessonService_ServiceDesc is the grpc.ServiceDesc for LessonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LessonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lessons.LessonService",
	HandlerType: (*LessonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LessonService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LessonService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LessonService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _LessonService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _LessonService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lessons.proto",
}
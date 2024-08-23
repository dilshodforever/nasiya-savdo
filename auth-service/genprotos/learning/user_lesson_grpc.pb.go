// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user_lesson.proto

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
	UserLessonService_Create_FullMethodName    = "/user_lessons.UserLessonService/Create"
	UserLessonService_Update_FullMethodName    = "/user_lessons.UserLessonService/Update"
	UserLessonService_Delete_FullMethodName    = "/user_lessons.UserLessonService/Delete"
	UserLessonService_GetById_FullMethodName   = "/user_lessons.UserLessonService/GetById"
	UserLessonService_GetAll_FullMethodName    = "/user_lessons.UserLessonService/GetAll"
	UserLessonService_Completed_FullMethodName = "/user_lessons.UserLessonService/Completed"
)

// UserLessonServiceClient is the client API for UserLessonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLessonServiceClient interface {
	Create(ctx context.Context, in *UserLessonReq, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *UserLesson, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*UserLesson, error)
	GetAll(ctx context.Context, in *UserLessonFilter, opts ...grpc.CallOption) (*AllUserLessons, error)
	Completed(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type userLessonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLessonServiceClient(cc grpc.ClientConnInterface) UserLessonServiceClient {
	return &userLessonServiceClient{cc}
}

func (c *userLessonServiceClient) Create(ctx context.Context, in *UserLessonReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, UserLessonService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) Update(ctx context.Context, in *UserLesson, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, UserLessonService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, UserLessonService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*UserLesson, error) {
	out := new(UserLesson)
	err := c.cc.Invoke(ctx, UserLessonService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) GetAll(ctx context.Context, in *UserLessonFilter, opts ...grpc.CallOption) (*AllUserLessons, error) {
	out := new(AllUserLessons)
	err := c.cc.Invoke(ctx, UserLessonService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) Completed(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, UserLessonService_Completed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLessonServiceServer is the server API for UserLessonService service.
// All implementations must embed UnimplementedUserLessonServiceServer
// for forward compatibility
type UserLessonServiceServer interface {
	Create(context.Context, *UserLessonReq) (*Void, error)
	Update(context.Context, *UserLesson) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	GetById(context.Context, *ById) (*UserLesson, error)
	GetAll(context.Context, *UserLessonFilter) (*AllUserLessons, error)
	Completed(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedUserLessonServiceServer()
}

// UnimplementedUserLessonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserLessonServiceServer struct {
}

func (UnimplementedUserLessonServiceServer) Create(context.Context, *UserLessonReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserLessonServiceServer) Update(context.Context, *UserLesson) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserLessonServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserLessonServiceServer) GetById(context.Context, *ById) (*UserLesson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserLessonServiceServer) GetAll(context.Context, *UserLessonFilter) (*AllUserLessons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserLessonServiceServer) Completed(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Completed not implemented")
}
func (UnimplementedUserLessonServiceServer) mustEmbedUnimplementedUserLessonServiceServer() {}

// UnsafeUserLessonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLessonServiceServer will
// result in compilation errors.
type UnsafeUserLessonServiceServer interface {
	mustEmbedUnimplementedUserLessonServiceServer()
}

func RegisterUserLessonServiceServer(s grpc.ServiceRegistrar, srv UserLessonServiceServer) {
	s.RegisterService(&UserLessonService_ServiceDesc, srv)
}

func _UserLessonService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLessonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).Create(ctx, req.(*UserLessonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLesson)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).Update(ctx, req.(*UserLesson))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLessonFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).GetAll(ctx, req.(*UserLessonFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_Completed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).Completed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_Completed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).Completed(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLessonService_ServiceDesc is the grpc.ServiceDesc for UserLessonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLessonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_lessons.UserLessonService",
	HandlerType: (*UserLessonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserLessonService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserLessonService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserLessonService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _UserLessonService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserLessonService_GetAll_Handler,
		},
		{
			MethodName: "Completed",
			Handler:    _UserLessonService_Completed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_lesson.proto",
}
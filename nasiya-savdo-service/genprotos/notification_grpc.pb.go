// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: notification.proto

package genprotos

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

// NotificationtServiceClient is the client API for NotificationtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationtServiceClient interface {
	GetNotification(ctx context.Context, in *GetNotificationByidRequest, opts ...grpc.CallOption) (*GetNotificationByidResponse, error)
	DeleteNotification(ctx context.Context, in *GetNotificationByidRequest, opts ...grpc.CallOption) (*NotificationsResponse, error)
	ListNotification(ctx context.Context, in *GetNotificationByidRequest, opts ...grpc.CallOption) (*ListNotificationResponse, error)
}

type notificationtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationtServiceClient(cc grpc.ClientConnInterface) NotificationtServiceClient {
	return &notificationtServiceClient{cc}
}

func (c *notificationtServiceClient) GetNotification(ctx context.Context, in *GetNotificationByidRequest, opts ...grpc.CallOption) (*GetNotificationByidResponse, error) {
	out := new(GetNotificationByidResponse)
	err := c.cc.Invoke(ctx, "/notifications.NotificationtService/GetNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationtServiceClient) DeleteNotification(ctx context.Context, in *GetNotificationByidRequest, opts ...grpc.CallOption) (*NotificationsResponse, error) {
	out := new(NotificationsResponse)
	err := c.cc.Invoke(ctx, "/notifications.NotificationtService/DeleteNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationtServiceClient) ListNotification(ctx context.Context, in *GetNotificationByidRequest, opts ...grpc.CallOption) (*ListNotificationResponse, error) {
	out := new(ListNotificationResponse)
	err := c.cc.Invoke(ctx, "/notifications.NotificationtService/ListNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationtServiceServer is the server API for NotificationtService service.
// All implementations must embed UnimplementedNotificationtServiceServer
// for forward compatibility
type NotificationtServiceServer interface {
	GetNotification(context.Context, *GetNotificationByidRequest) (*GetNotificationByidResponse, error)
	DeleteNotification(context.Context, *GetNotificationByidRequest) (*NotificationsResponse, error)
	ListNotification(context.Context, *GetNotificationByidRequest) (*ListNotificationResponse, error)
	mustEmbedUnimplementedNotificationtServiceServer()
}

// UnimplementedNotificationtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationtServiceServer struct {
}

func (UnimplementedNotificationtServiceServer) GetNotification(context.Context, *GetNotificationByidRequest) (*GetNotificationByidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedNotificationtServiceServer) DeleteNotification(context.Context, *GetNotificationByidRequest) (*NotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotification not implemented")
}
func (UnimplementedNotificationtServiceServer) ListNotification(context.Context, *GetNotificationByidRequest) (*ListNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotification not implemented")
}
func (UnimplementedNotificationtServiceServer) mustEmbedUnimplementedNotificationtServiceServer() {}

// UnsafeNotificationtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationtServiceServer will
// result in compilation errors.
type UnsafeNotificationtServiceServer interface {
	mustEmbedUnimplementedNotificationtServiceServer()
}

func RegisterNotificationtServiceServer(s grpc.ServiceRegistrar, srv NotificationtServiceServer) {
	s.RegisterService(&NotificationtService_ServiceDesc, srv)
}

func _NotificationtService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationByidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationtServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.NotificationtService/GetNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationtServiceServer).GetNotification(ctx, req.(*GetNotificationByidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationtService_DeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationByidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationtServiceServer).DeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.NotificationtService/DeleteNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationtServiceServer).DeleteNotification(ctx, req.(*GetNotificationByidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationtService_ListNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationByidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationtServiceServer).ListNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.NotificationtService/ListNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationtServiceServer).ListNotification(ctx, req.(*GetNotificationByidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationtService_ServiceDesc is the grpc.ServiceDesc for NotificationtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notifications.NotificationtService",
	HandlerType: (*NotificationtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotification",
			Handler:    _NotificationtService_GetNotification_Handler,
		},
		{
			MethodName: "DeleteNotification",
			Handler:    _NotificationtService_DeleteNotification_Handler,
		},
		{
			MethodName: "ListNotification",
			Handler:    _NotificationtService_ListNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}

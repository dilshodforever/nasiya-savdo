// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: reports.proto

package reports

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

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	GenerateSpendingReport(ctx context.Context, in *GenerateSpendingReportRequest, opts ...grpc.CallOption) (*GenerateSpendingReportResponse, error)
	GenerateIncomeReport(ctx context.Context, in *GenerateIncomeReportRequest, opts ...grpc.CallOption) (*GenerateIncomeReportResponse, error)
	GenerateBudgetPerformanceReport(ctx context.Context, in *GenerateBudgetPerformanceReportRequest, opts ...grpc.CallOption) (*GenerateBudgetPerformanceReportResponse, error)
	GenerateGoalProgressReport(ctx context.Context, in *GenerateGoalProgressReportRequest, opts ...grpc.CallOption) (*GenerateGoalProgressReportResponse, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) GenerateSpendingReport(ctx context.Context, in *GenerateSpendingReportRequest, opts ...grpc.CallOption) (*GenerateSpendingReportResponse, error) {
	out := new(GenerateSpendingReportResponse)
	err := c.cc.Invoke(ctx, "/budgets.ReportService/GenerateSpendingReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GenerateIncomeReport(ctx context.Context, in *GenerateIncomeReportRequest, opts ...grpc.CallOption) (*GenerateIncomeReportResponse, error) {
	out := new(GenerateIncomeReportResponse)
	err := c.cc.Invoke(ctx, "/budgets.ReportService/GenerateIncomeReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GenerateBudgetPerformanceReport(ctx context.Context, in *GenerateBudgetPerformanceReportRequest, opts ...grpc.CallOption) (*GenerateBudgetPerformanceReportResponse, error) {
	out := new(GenerateBudgetPerformanceReportResponse)
	err := c.cc.Invoke(ctx, "/budgets.ReportService/GenerateBudgetPerformanceReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GenerateGoalProgressReport(ctx context.Context, in *GenerateGoalProgressReportRequest, opts ...grpc.CallOption) (*GenerateGoalProgressReportResponse, error) {
	out := new(GenerateGoalProgressReportResponse)
	err := c.cc.Invoke(ctx, "/budgets.ReportService/GenerateGoalProgressReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations must embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	GenerateSpendingReport(context.Context, *GenerateSpendingReportRequest) (*GenerateSpendingReportResponse, error)
	GenerateIncomeReport(context.Context, *GenerateIncomeReportRequest) (*GenerateIncomeReportResponse, error)
	GenerateBudgetPerformanceReport(context.Context, *GenerateBudgetPerformanceReportRequest) (*GenerateBudgetPerformanceReportResponse, error)
	GenerateGoalProgressReport(context.Context, *GenerateGoalProgressReportRequest) (*GenerateGoalProgressReportResponse, error)
	mustEmbedUnimplementedReportServiceServer()
}

// UnimplementedReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (UnimplementedReportServiceServer) GenerateSpendingReport(context.Context, *GenerateSpendingReportRequest) (*GenerateSpendingReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSpendingReport not implemented")
}
func (UnimplementedReportServiceServer) GenerateIncomeReport(context.Context, *GenerateIncomeReportRequest) (*GenerateIncomeReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateIncomeReport not implemented")
}
func (UnimplementedReportServiceServer) GenerateBudgetPerformanceReport(context.Context, *GenerateBudgetPerformanceReportRequest) (*GenerateBudgetPerformanceReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateBudgetPerformanceReport not implemented")
}
func (UnimplementedReportServiceServer) GenerateGoalProgressReport(context.Context, *GenerateGoalProgressReportRequest) (*GenerateGoalProgressReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateGoalProgressReport not implemented")
}
func (UnimplementedReportServiceServer) mustEmbedUnimplementedReportServiceServer() {}

// UnsafeReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServiceServer will
// result in compilation errors.
type UnsafeReportServiceServer interface {
	mustEmbedUnimplementedReportServiceServer()
}

func RegisterReportServiceServer(s grpc.ServiceRegistrar, srv ReportServiceServer) {
	s.RegisterService(&ReportService_ServiceDesc, srv)
}

func _ReportService_GenerateSpendingReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateSpendingReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GenerateSpendingReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgets.ReportService/GenerateSpendingReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GenerateSpendingReport(ctx, req.(*GenerateSpendingReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_GenerateIncomeReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIncomeReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GenerateIncomeReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgets.ReportService/GenerateIncomeReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GenerateIncomeReport(ctx, req.(*GenerateIncomeReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_GenerateBudgetPerformanceReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateBudgetPerformanceReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GenerateBudgetPerformanceReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgets.ReportService/GenerateBudgetPerformanceReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GenerateBudgetPerformanceReport(ctx, req.(*GenerateBudgetPerformanceReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_GenerateGoalProgressReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateGoalProgressReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GenerateGoalProgressReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgets.ReportService/GenerateGoalProgressReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GenerateGoalProgressReport(ctx, req.(*GenerateGoalProgressReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportService_ServiceDesc is the grpc.ServiceDesc for ReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "budgets.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateSpendingReport",
			Handler:    _ReportService_GenerateSpendingReport_Handler,
		},
		{
			MethodName: "GenerateIncomeReport",
			Handler:    _ReportService_GenerateIncomeReport_Handler,
		},
		{
			MethodName: "GenerateBudgetPerformanceReport",
			Handler:    _ReportService_GenerateBudgetPerformanceReport_Handler,
		},
		{
			MethodName: "GenerateGoalProgressReport",
			Handler:    _ReportService_GenerateGoalProgressReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reports.proto",
}
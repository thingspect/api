// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.24.4
// source: api/thingspect_datapoint.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	DataPointService_PublishDataPoints_FullMethodName = "/thingspect.api.DataPointService/PublishDataPoints"
	DataPointService_ListDataPoints_FullMethodName    = "/thingspect.api.DataPointService/ListDataPoints"
	DataPointService_LatestDataPoints_FullMethodName  = "/thingspect.api.DataPointService/LatestDataPoints"
)

// DataPointServiceClient is the client API for DataPointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// DataPointService contains functions to create and query data points.
type DataPointServiceClient interface {
	// Publish data points via the API. Data points are generated by devices.
	PublishDataPoints(ctx context.Context, in *PublishDataPointsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List all data points for a device in an [end, start) time range, in descending timestamp order. Data points are generated by devices.
	ListDataPoints(ctx context.Context, in *ListDataPointsRequest, opts ...grpc.CallOption) (*ListDataPointsResponse, error)
	// List the latest data point for each of a device's attributes in a [now, start) time range. Data points are generated by devices.
	LatestDataPoints(ctx context.Context, in *LatestDataPointsRequest, opts ...grpc.CallOption) (*LatestDataPointsResponse, error)
}

type dataPointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataPointServiceClient(cc grpc.ClientConnInterface) DataPointServiceClient {
	return &dataPointServiceClient{cc}
}

func (c *dataPointServiceClient) PublishDataPoints(ctx context.Context, in *PublishDataPointsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DataPointService_PublishDataPoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPointServiceClient) ListDataPoints(ctx context.Context, in *ListDataPointsRequest, opts ...grpc.CallOption) (*ListDataPointsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDataPointsResponse)
	err := c.cc.Invoke(ctx, DataPointService_ListDataPoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPointServiceClient) LatestDataPoints(ctx context.Context, in *LatestDataPointsRequest, opts ...grpc.CallOption) (*LatestDataPointsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LatestDataPointsResponse)
	err := c.cc.Invoke(ctx, DataPointService_LatestDataPoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataPointServiceServer is the server API for DataPointService service.
// All implementations must embed UnimplementedDataPointServiceServer
// for forward compatibility
//
// DataPointService contains functions to create and query data points.
type DataPointServiceServer interface {
	// Publish data points via the API. Data points are generated by devices.
	PublishDataPoints(context.Context, *PublishDataPointsRequest) (*emptypb.Empty, error)
	// List all data points for a device in an [end, start) time range, in descending timestamp order. Data points are generated by devices.
	ListDataPoints(context.Context, *ListDataPointsRequest) (*ListDataPointsResponse, error)
	// List the latest data point for each of a device's attributes in a [now, start) time range. Data points are generated by devices.
	LatestDataPoints(context.Context, *LatestDataPointsRequest) (*LatestDataPointsResponse, error)
	mustEmbedUnimplementedDataPointServiceServer()
}

// UnimplementedDataPointServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataPointServiceServer struct {
}

func (UnimplementedDataPointServiceServer) PublishDataPoints(context.Context, *PublishDataPointsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishDataPoints not implemented")
}
func (UnimplementedDataPointServiceServer) ListDataPoints(context.Context, *ListDataPointsRequest) (*ListDataPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataPoints not implemented")
}
func (UnimplementedDataPointServiceServer) LatestDataPoints(context.Context, *LatestDataPointsRequest) (*LatestDataPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestDataPoints not implemented")
}
func (UnimplementedDataPointServiceServer) mustEmbedUnimplementedDataPointServiceServer() {}

// UnsafeDataPointServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataPointServiceServer will
// result in compilation errors.
type UnsafeDataPointServiceServer interface {
	mustEmbedUnimplementedDataPointServiceServer()
}

func RegisterDataPointServiceServer(s grpc.ServiceRegistrar, srv DataPointServiceServer) {
	s.RegisterService(&DataPointService_ServiceDesc, srv)
}

func _DataPointService_PublishDataPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishDataPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPointServiceServer).PublishDataPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataPointService_PublishDataPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPointServiceServer).PublishDataPoints(ctx, req.(*PublishDataPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPointService_ListDataPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPointServiceServer).ListDataPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataPointService_ListDataPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPointServiceServer).ListDataPoints(ctx, req.(*ListDataPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPointService_LatestDataPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestDataPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPointServiceServer).LatestDataPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataPointService_LatestDataPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPointServiceServer).LatestDataPoints(ctx, req.(*LatestDataPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataPointService_ServiceDesc is the grpc.ServiceDesc for DataPointService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataPointService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thingspect.api.DataPointService",
	HandlerType: (*DataPointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishDataPoints",
			Handler:    _DataPointService_PublishDataPoints_Handler,
		},
		{
			MethodName: "ListDataPoints",
			Handler:    _DataPointService_ListDataPoints_Handler,
		},
		{
			MethodName: "LatestDataPoints",
			Handler:    _DataPointService_LatestDataPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/thingspect_datapoint.proto",
}

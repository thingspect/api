// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: api/thingspect_session.proto

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SessionService_Login_FullMethodName     = "/thingspect.api.SessionService/Login"
	SessionService_CreateKey_FullMethodName = "/thingspect.api.SessionService/CreateKey"
	SessionService_DeleteKey_FullMethodName = "/thingspect.api.SessionService/DeleteKey"
	SessionService_ListKeys_FullMethodName  = "/thingspect.api.SessionService/ListKeys"
)

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SessionService contains functions to create sessions and keys.
type SessionServiceClient interface {
	// Log in a user. Login tokens are time-limited and accompanied by an expiration.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Create an API key. API keys are persistent and do not expire until revoked.
	CreateKey(ctx context.Context, in *CreateKeyRequest, opts ...grpc.CallOption) (*CreateKeyResponse, error)
	// Delete an API key by ID. API keys are persistent and do not expire until revoked.
	DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List all API keys. API keys are persistent and do not expire until revoked.
	ListKeys(ctx context.Context, in *ListKeysRequest, opts ...grpc.CallOption) (*ListKeysResponse, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SessionService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CreateKey(ctx context.Context, in *CreateKeyRequest, opts ...grpc.CallOption) (*CreateKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateKeyResponse)
	err := c.cc.Invoke(ctx, SessionService_CreateKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SessionService_DeleteKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) ListKeys(ctx context.Context, in *ListKeysRequest, opts ...grpc.CallOption) (*ListKeysResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListKeysResponse)
	err := c.cc.Invoke(ctx, SessionService_ListKeys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
// All implementations must embed UnimplementedSessionServiceServer
// for forward compatibility.
//
// SessionService contains functions to create sessions and keys.
type SessionServiceServer interface {
	// Log in a user. Login tokens are time-limited and accompanied by an expiration.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Create an API key. API keys are persistent and do not expire until revoked.
	CreateKey(context.Context, *CreateKeyRequest) (*CreateKeyResponse, error)
	// Delete an API key by ID. API keys are persistent and do not expire until revoked.
	DeleteKey(context.Context, *DeleteKeyRequest) (*emptypb.Empty, error)
	// List all API keys. API keys are persistent and do not expire until revoked.
	ListKeys(context.Context, *ListKeysRequest) (*ListKeysResponse, error)
	mustEmbedUnimplementedSessionServiceServer()
}

// UnimplementedSessionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSessionServiceServer struct{}

func (UnimplementedSessionServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSessionServiceServer) CreateKey(context.Context, *CreateKeyRequest) (*CreateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKey not implemented")
}
func (UnimplementedSessionServiceServer) DeleteKey(context.Context, *DeleteKeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (UnimplementedSessionServiceServer) ListKeys(context.Context, *ListKeysRequest) (*ListKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeys not implemented")
}
func (UnimplementedSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {}
func (UnimplementedSessionServiceServer) testEmbeddedByValue()                        {}

// UnsafeSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServiceServer will
// result in compilation errors.
type UnsafeSessionServiceServer interface {
	mustEmbedUnimplementedSessionServiceServer()
}

func RegisterSessionServiceServer(s grpc.ServiceRegistrar, srv SessionServiceServer) {
	// If the following call pancis, it indicates UnimplementedSessionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SessionService_ServiceDesc, srv)
}

func _SessionService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SessionService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CreateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SessionService_CreateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateKey(ctx, req.(*CreateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SessionService_DeleteKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteKey(ctx, req.(*DeleteKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_ListKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).ListKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SessionService_ListKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).ListKeys(ctx, req.(*ListKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionService_ServiceDesc is the grpc.ServiceDesc for SessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thingspect.api.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SessionService_Login_Handler,
		},
		{
			MethodName: "CreateKey",
			Handler:    _SessionService_CreateKey_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _SessionService_DeleteKey_Handler,
		},
		{
			MethodName: "ListKeys",
			Handler:    _SessionService_ListKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/thingspect_session.proto",
}

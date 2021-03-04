// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RuleServiceClient is the client API for RuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuleServiceClient interface {
	// Create a rule. Rules define how events are generated.
	CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*Rule, error)
	// Get a rule by ID. Rules define how events are generated.
	GetRule(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error)
	// Update a rule. Rules define how events are generated.
	UpdateRule(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*Rule, error)
	// Delete a rule by ID. Rules define how events are generated.
	DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List all rules. Rules define how events are generated.
	ListRules(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error)
	// Test a rule. Rules define how events are generated.
	TestRule(ctx context.Context, in *TestRuleRequest, opts ...grpc.CallOption) (*TestRuleResponse, error)
}

type ruleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuleServiceClient(cc grpc.ClientConnInterface) RuleServiceClient {
	return &ruleServiceClient{cc}
}

func (c *ruleServiceClient) CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, "/thingspect.api.RuleService/CreateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) GetRule(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, "/thingspect.api.RuleService/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) UpdateRule(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, "/thingspect.api.RuleService/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/thingspect.api.RuleService/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) ListRules(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error) {
	out := new(ListRulesResponse)
	err := c.cc.Invoke(ctx, "/thingspect.api.RuleService/ListRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) TestRule(ctx context.Context, in *TestRuleRequest, opts ...grpc.CallOption) (*TestRuleResponse, error) {
	out := new(TestRuleResponse)
	err := c.cc.Invoke(ctx, "/thingspect.api.RuleService/TestRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuleServiceServer is the server API for RuleService service.
// All implementations must embed UnimplementedRuleServiceServer
// for forward compatibility
type RuleServiceServer interface {
	// Create a rule. Rules define how events are generated.
	CreateRule(context.Context, *CreateRuleRequest) (*Rule, error)
	// Get a rule by ID. Rules define how events are generated.
	GetRule(context.Context, *GetRuleRequest) (*Rule, error)
	// Update a rule. Rules define how events are generated.
	UpdateRule(context.Context, *UpdateRuleRequest) (*Rule, error)
	// Delete a rule by ID. Rules define how events are generated.
	DeleteRule(context.Context, *DeleteRuleRequest) (*empty.Empty, error)
	// List all rules. Rules define how events are generated.
	ListRules(context.Context, *ListRulesRequest) (*ListRulesResponse, error)
	// Test a rule. Rules define how events are generated.
	TestRule(context.Context, *TestRuleRequest) (*TestRuleResponse, error)
	mustEmbedUnimplementedRuleServiceServer()
}

// UnimplementedRuleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRuleServiceServer struct {
}

func (UnimplementedRuleServiceServer) CreateRule(context.Context, *CreateRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRule not implemented")
}
func (UnimplementedRuleServiceServer) GetRule(context.Context, *GetRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRule not implemented")
}
func (UnimplementedRuleServiceServer) UpdateRule(context.Context, *UpdateRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRule not implemented")
}
func (UnimplementedRuleServiceServer) DeleteRule(context.Context, *DeleteRuleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (UnimplementedRuleServiceServer) ListRules(context.Context, *ListRulesRequest) (*ListRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRules not implemented")
}
func (UnimplementedRuleServiceServer) TestRule(context.Context, *TestRuleRequest) (*TestRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRule not implemented")
}
func (UnimplementedRuleServiceServer) mustEmbedUnimplementedRuleServiceServer() {}

// UnsafeRuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuleServiceServer will
// result in compilation errors.
type UnsafeRuleServiceServer interface {
	mustEmbedUnimplementedRuleServiceServer()
}

func RegisterRuleServiceServer(s grpc.ServiceRegistrar, srv RuleServiceServer) {
	s.RegisterService(&RuleService_ServiceDesc, srv)
}

func _RuleService_CreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).CreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thingspect.api.RuleService/CreateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).CreateRule(ctx, req.(*CreateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thingspect.api.RuleService/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).GetRule(ctx, req.(*GetRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thingspect.api.RuleService/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).UpdateRule(ctx, req.(*UpdateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thingspect.api.RuleService/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).DeleteRule(ctx, req.(*DeleteRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_ListRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).ListRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thingspect.api.RuleService/ListRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).ListRules(ctx, req.(*ListRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_TestRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).TestRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thingspect.api.RuleService/TestRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).TestRule(ctx, req.(*TestRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RuleService_ServiceDesc is the grpc.ServiceDesc for RuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thingspect.api.RuleService",
	HandlerType: (*RuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRule",
			Handler:    _RuleService_CreateRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _RuleService_GetRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _RuleService_UpdateRule_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _RuleService_DeleteRule_Handler,
		},
		{
			MethodName: "ListRules",
			Handler:    _RuleService_ListRules_Handler,
		},
		{
			MethodName: "TestRule",
			Handler:    _RuleService_TestRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/rule.proto",
}

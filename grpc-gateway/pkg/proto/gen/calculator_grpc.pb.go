// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: calculator.proto

package calculator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Calculator_Add_FullMethodName      = "/calculator.Calculator/Add"
	Calculator_Subtract_FullMethodName = "/calculator.Calculator/Subtract"
	Calculator_Multiply_FullMethodName = "/calculator.Calculator/Multiply"
	Calculator_Divide_FullMethodName   = "/calculator.Calculator/Divide"
)

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Calculator service definition
type CalculatorClient interface {
	// Unary RPC for addition
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Unary RPC for subtraction
	Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error)
	// Unary RPC for multiplication
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
	// Unary RPC for division
	Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, Calculator_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubtractResponse)
	err := c.cc.Invoke(ctx, Calculator_Subtract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultiplyResponse)
	err := c.cc.Invoke(ctx, Calculator_Multiply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DivideResponse)
	err := c.cc.Invoke(ctx, Calculator_Divide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations should embed UnimplementedCalculatorServer
// for forward compatibility.
//
// Calculator service definition
type CalculatorServer interface {
	// Unary RPC for addition
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Unary RPC for subtraction
	Subtract(context.Context, *SubtractRequest) (*SubtractResponse, error)
	// Unary RPC for multiplication
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
	// Unary RPC for division
	Divide(context.Context, *DivideRequest) (*DivideResponse, error)
}

// UnimplementedCalculatorServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalculatorServer struct{}

func (UnimplementedCalculatorServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServer) Subtract(context.Context, *SubtractRequest) (*SubtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedCalculatorServer) Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedCalculatorServer) Divide(context.Context, *DivideRequest) (*DivideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedCalculatorServer) testEmbeddedByValue() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	// If the following call pancis, it indicates UnimplementedCalculatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Subtract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Subtract(ctx, req.(*SubtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Divide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Divide(ctx, req.(*DivideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _Calculator_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Calculator_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Calculator_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}

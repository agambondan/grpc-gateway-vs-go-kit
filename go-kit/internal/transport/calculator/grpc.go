package calculator

import (
	"context"

	"git.bluebird.id/firman.agam/go-kit/internal/endpoint"
	"git.bluebird.id/firman.agam/go-kit/internal/transport/calculator/decoder"
	"git.bluebird.id/firman.agam/go-kit/internal/transport/calculator/encoder"
	"git.bluebird.id/firman.agam/go-kit/internal/transport/middleware"
	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
	"git.bluebird.id/promo/packages/zaplog"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
)

type grpcServer struct {
	add       grpcTransport.Handler
	divide    grpcTransport.Handler
	multiply  grpcTransport.Handler
	subtract  grpcTransport.Handler
	fibonacci grpcTransport.Handler
}

func NewGRPCCalculator(endpoint endpoint.Endpoints, Opts ...grpcTransport.ServerOption) calculator.CalculatorServer {
	// Add Endpoint
	addEndpoint := endpoint.AddEndpoint()
	addEndpoint = middleware.TransportLogging("add_endpoint")(addEndpoint)
	addEndpoint = middleware.APIKeyGRPCAuth()(addEndpoint)
	addRpc := kitGRPC.NewServer(
		addEndpoint,
		decoder.GRPCAdd,
		encoder.GRPCAdd,
		Opts...,
	)

	// Subtract Endpoint
	subtractEndpoint := endpoint.SubtractEndpoint()
	subtractEndpoint = middleware.TransportLogging("subtract_endpoint")(subtractEndpoint)
	subtractEndpoint = middleware.APIKeyGRPCAuth()(subtractEndpoint)
	subtractRpc := kitGRPC.NewServer(
		subtractEndpoint,
		decoder.GRPCSubtract,
		encoder.GRPCSubtract,
		Opts...,
	)

	// Multiply Endpoint
	multiplyEndpoint := endpoint.MultiplyEndpoint()
	multiplyEndpoint = middleware.TransportLogging("multiply_endpoint")(multiplyEndpoint)
	multiplyEndpoint = middleware.APIKeyGRPCAuth()(multiplyEndpoint)
	multiplyRpc := kitGRPC.NewServer(
		multiplyEndpoint,
		decoder.GRPCMultiply,
		encoder.GRPCMultiply,
		Opts...,
	)

	// Divide Endpoint
	divideEndpoint := endpoint.DivideEndpoint()
	divideEndpoint = middleware.TransportLogging("divide_endpoint")(divideEndpoint)
	divideEndpoint = middleware.APIKeyGRPCAuth()(divideEndpoint)
	divideRpc := kitGRPC.NewServer(
		divideEndpoint,
		decoder.GRPCDivide,
		encoder.GRPCDivide,
		Opts...,
	)

	// Fibonacci Endpoint
	fibonacciEndpoint := endpoint.FibonacciEndpoint()
	fibonacciEndpoint = middleware.TransportLogging("fibonacci_endpoint")(fibonacciEndpoint)
	fibonacciEndpoint = middleware.APIKeyGRPCAuth()(fibonacciEndpoint)
	fibonacciRpc := kitGRPC.NewServer(
		fibonacciEndpoint,
		decoder.GRPCFibonacci,
		encoder.GRPCFibonacci,
		Opts...,
	)
	return &grpcServer{
		add:       addRpc,
		divide:    divideRpc,
		multiply:  multiplyRpc,
		subtract:  subtractRpc,
		fibonacci: fibonacciRpc,
	}
}

// Add implements calculator.CalculatorServer.
func (g *grpcServer) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	_, resp, err := g.add.ServeGRPC(ctx, req)
	if err != nil {
		logger.Info("failed implement create subs code", zap.Error(err))
		return nil, err
	}
	return resp.(*calculator.AddResponse), nil
}

// Divide implements calculator.CalculatorServer.
func (g *grpcServer) Divide(ctx context.Context, req *calculator.DivideRequest) (*calculator.DivideResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	_, resp, err := g.divide.ServeGRPC(ctx, req)
	if err != nil {
		logger.Info("failed implement create subs code", zap.Error(err))
		return nil, err
	}
	return resp.(*calculator.DivideResponse), nil
}

// Multiply implements calculator.CalculatorServer.
func (g *grpcServer) Multiply(ctx context.Context, req *calculator.MultiplyRequest) (*calculator.MultiplyResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	_, resp, err := g.multiply.ServeGRPC(ctx, req)
	if err != nil {
		logger.Info("failed implement create subs code", zap.Error(err))
		return nil, err
	}
	return resp.(*calculator.MultiplyResponse), nil
}

// Subtract implements calculator.CalculatorServer.
func (g *grpcServer) Subtract(ctx context.Context, req *calculator.SubtractRequest) (*calculator.SubtractResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	_, resp, err := g.subtract.ServeGRPC(ctx, req)
	if err != nil {
		logger.Info("failed implement create subs code", zap.Error(err))
		return nil, err
	}
	return resp.(*calculator.SubtractResponse), nil
}

// Fibonacci implements calculator.CalculatorServer.
func (g *grpcServer) Fibonacci(ctx context.Context, req *calculator.FibonacciRequest) (*calculator.FibonacciResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	_, resp, err := g.fibonacci.ServeGRPC(ctx, req)
	if err != nil {
		logger.Info("failed implement create subs code", zap.Error(err))
		return nil, err
	}
	return resp.(*calculator.FibonacciResponse), nil
}

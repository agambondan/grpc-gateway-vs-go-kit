package calculatorgrpc

import (
	"context"

	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
)

type calculatorGRPCServer struct {
	svc calculator.CalculatorServer
}

// Add implements calculator.CalculatorServer.
func (c *calculatorGRPCServer) Add(context.Context, *calculator.AddRequest) (*calculator.AddResponse, error) {
	panic("unimplemented")
}

// Divide implements calculator.CalculatorServer.
func (c *calculatorGRPCServer) Divide(context.Context, *calculator.DivideRequest) (*calculator.DivideResponse, error) {
	panic("unimplemented")
}

// Multiply implements calculator.CalculatorServer.
func (c *calculatorGRPCServer) Multiply(context.Context, *calculator.MultiplyRequest) (*calculator.MultiplyResponse, error) {
	panic("unimplemented")
}

// Subtract implements calculator.CalculatorServer.
func (c *calculatorGRPCServer) Subtract(context.Context, *calculator.SubtractRequest) (*calculator.SubtractResponse, error) {
	panic("unimplemented")
}

func NewCalculatorGRPCServer(svc calculator.CalculatorServer) calculator.CalculatorServer {
	return &calculatorGRPCServer{
		svc: svc,
	}
}

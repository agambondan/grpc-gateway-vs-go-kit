package endpoint

import (
	"context"
	"fmt"

	"git.bluebird.id/firman.agam/go-kit/internal/service"
	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints interface {
	AddEndpoint() endpoint.Endpoint
	SubtractEndpoint() endpoint.Endpoint
	MultiplyEndpoint() endpoint.Endpoint
	DivideEndpoint() endpoint.Endpoint
	FibonacciEndpoint() endpoint.Endpoint
}

type endpoints struct {
	srv service.CalculatorService
}

func NewEndpoints(srv service.CalculatorService) Endpoints {
	return &endpoints{
		srv: srv,
	}
}

// AddEndpoint creates the Add endpoint
func (e *endpoints) AddEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.AddRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := e.srv.Add(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return calculator.AddResponse{Result: result}, nil
	}
}

// SubtractEndpoint creates the Subtract endpoint
func (e *endpoints) SubtractEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.SubtractRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := e.srv.Subtract(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return calculator.SubtractResponse{Result: result}, nil
	}
}

// MultiplyEndpoint creates the Multiply endpoint
func (e *endpoints) MultiplyEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.MultiplyRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := e.srv.Multiply(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return calculator.MultiplyResponse{Result: result}, nil
	}
}

// DivideEndpoint creates the Divide endpoint
func (e *endpoints) DivideEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.DivideRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := e.srv.Divide(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return calculator.DivideResponse{Result: result}, nil
	}
}

// FibonacciEndpoint creates the Fibonacci endpoint
func (e *endpoints) FibonacciEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.FibonacciRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := e.srv.Fibonacci(ctx, req.N)
		if err != nil {
			return nil, err
		}
		return calculator.FibonacciResponse{Result: result}, nil
	}
}

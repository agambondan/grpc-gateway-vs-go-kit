package endpoint

import (
	"context"
	"fmt"

	"git.bluebird.id/firman.agam/go-kit/internal/service"
	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
	"github.com/go-kit/kit/endpoint"
)

// MakeAddEndpoint creates the Add endpoint
func MakeAddEndpoint(srv service.CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.AddRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := srv.Add(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return calculator.AddResponse{Result: result}, nil
	}
}

// MakeSubtractEndpoint creates the Subtract endpoint
func MakeSubtractEndpoint(srv service.CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.SubtractRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := srv.Subtract(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return calculator.SubtractResponse{Result: result}, nil
	}
}

// MakeMultiplyEndpoint creates the Multiply endpoint
func MakeMultiplyEndpoint(srv service.CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.MultiplyRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := srv.Multiply(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return calculator.MultiplyResponse{Result: result}, nil
	}
}

// MakeDivideEndpoint creates the Divide endpoint
func MakeDivideEndpoint(srv service.CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.DivideRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := srv.Divide(ctx, req.A, req.B)
		if err != nil {
			return nil, err
		}
		return calculator.DivideResponse{Result: result}, nil
	}
}

// MakeFibonacciEndpoint creates the Fibonacci endpoint
func MakeFibonacciEndpoint(srv service.CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*calculator.FibonacciRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		result, err := srv.Fibonacci(ctx, req.N)
		if err != nil {
			return nil, err
		}
		return calculator.FibonacciResponse{Result: result}, nil
	}
}

// // MakeLoopingEndpoint creates the Looping endpoint
// func MakeLoopingEndpoint(srv service.CalculatorService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req, ok := request.(*calculator.LoopingRequest)
// 		if !ok {
// 			return nil, fmt.Errorf("invalid request type")
// 		}
// 		result, err := srv.Looping(ctx, req.N)
// 		if err != nil {
// 			return nil, err
// 		}
// 		// Add your logic here to perform looping and calculate time
// 		return calculator.LoopingResponse{}, nil
// 	}
// }

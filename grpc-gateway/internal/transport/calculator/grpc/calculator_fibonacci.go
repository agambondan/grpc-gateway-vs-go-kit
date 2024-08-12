package calculatorgrpc

import (
	"context"
	"time"

	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
	"git.bluebird.id/promo/packages/zaplog"
)

// Fibonacci implements calculator.CalculatorServer.
func (c *calculatorGRPCServer) Fibonacci(ctx context.Context, req *calculator.FibonacciRequest) (*calculator.FibonacciResponse, error) {
	start := time.Now()
	logger := zaplog.WithContext(ctx)

	result, err := c.svc.Fibonacci(ctx, req.N)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	executionTime := time.Since(start)

	return &calculator.FibonacciResponse{
		Result:         result,
		TimeMilisecond: int32(executionTime.Milliseconds()),
		TimeSecond:     int32(executionTime.Seconds()),
	}, nil
}

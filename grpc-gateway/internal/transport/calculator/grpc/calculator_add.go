package calculatorgrpc

import (
	"context"

	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
	"git.bluebird.id/promo/packages/zaplog"
)

// Add implements service.CalculatorService.
func (c *calculatorGRPCServer) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	logger := zaplog.WithContext(ctx)

	value, err := c.svc.Add(ctx, req.A, req.B)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &calculator.AddResponse{
		Result: value,
	}, nil
}

package calculatorgrpc

import (
	"context"

	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
	"git.bluebird.id/promo/packages/zaplog"
)

// Multiply implements service.CalculatorService.
func (c *calculatorGRPCServer) Multiply(ctx context.Context, req *calculator.MultiplyRequest) (*calculator.MultiplyResponse, error) {
	logger := zaplog.WithContext(ctx)

	value, err := c.svc.Multiply(ctx, req.A, req.B)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &calculator.MultiplyResponse{
		Result: value,
	}, nil
}

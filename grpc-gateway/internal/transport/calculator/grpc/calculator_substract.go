package calculatorgrpc

import (
	"context"

	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
	"git.bluebird.id/promo/packages/zaplog"
)

// Subtract implements service.CalculatorService.
func (c *calculatorGRPCServer) Subtract(ctx context.Context, req *calculator.SubtractRequest) (*calculator.SubtractResponse, error) {
	logger := zaplog.WithContext(ctx)

	value, err := c.svc.Subtract(ctx, req.A, req.B)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &calculator.SubtractResponse{
		Result: value,
	}, nil
}

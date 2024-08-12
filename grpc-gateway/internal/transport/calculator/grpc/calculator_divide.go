package calculatorgrpc

import (
	"context"

	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
	"git.bluebird.id/promo/packages/zaplog"
)

// Divide implements service.CalculatorService.
func (c *calculatorGRPCServer) Divide(ctx context.Context, req *calculator.DivideRequest) (*calculator.DivideResponse, error) {
	logger := zaplog.WithContext(ctx)

	value, err := c.svc.Divide(ctx, req.A, req.B)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &calculator.DivideResponse{
		Result: value,
	}, nil
}

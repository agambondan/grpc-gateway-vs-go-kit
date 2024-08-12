package calculatorgrpc

import (
	"context"

	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
)

// Looping implements calculator.CalculatorServer.
func (c *calculatorGRPCServer) Looping(ctx context.Context, req *calculator.LoopingRequest) (*calculator.LoopingResponse, error) {
	return nil, nil
}

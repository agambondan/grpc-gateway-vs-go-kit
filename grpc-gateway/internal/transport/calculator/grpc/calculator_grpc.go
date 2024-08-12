package calculatorgrpc

import (
	"git.bluebird.id/firman.agam/grpc-gateway/internal/service"
	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
)

type calculatorGRPCServer struct {
	svc service.CalculatorService
}

func NewCalculatorGRPCServer(svc service.CalculatorService) calculator.CalculatorServer {
	return &calculatorGRPCServer{
		svc: svc,
	}
}

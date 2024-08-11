package calculatorgrpc

import (
	"git.bluebird.id/firman.agam/grpc-gateway/internal/transport"
	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
	"google.golang.org/grpc"
)

type promotionGRPCServerRegistrar struct {
	promoServer calculator.CalculatorServer
}

func NewPromotionGRPCServerRegistrar(promoServer calculator.CalculatorServer) transport.GRPCService {
	return &promotionGRPCServerRegistrar{
		promoServer: promoServer,
	}
}

func (reg *promotionGRPCServerRegistrar) Register(server *grpc.Server) {
	calculator.RegisterCalculatorServer(server, reg.promoServer)
}

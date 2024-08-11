package calculatorhttp

import (
	"context"

	"git.bluebird.id/firman.agam/grpc-gateway/internal/transport"
	calculator "git.bluebird.id/firman.agam/grpc-gateway/pkg/proto/gen"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type promotionGRPCGatewayRegistrar struct{}

func NewPromotionGRPCGatewayRegistrar() transport.HTTPHandler {
	return &promotionGRPCGatewayRegistrar{}
}

func (reg *promotionGRPCGatewayRegistrar) Register(ctx context.Context, gwmux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return calculator.RegisterCalculatorHandler(ctx, gwmux, conn)
}

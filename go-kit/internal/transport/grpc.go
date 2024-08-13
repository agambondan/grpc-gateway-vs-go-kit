package transport

import (
	"context"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"git.bluebird.id/firman.agam/go-kit/internal/endpoint"
	calculatorTransport "git.bluebird.id/firman.agam/go-kit/internal/transport/calculator"
	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
	"git.bluebird.id/promo/packages/zaplog"
	kitJwt "github.com/go-kit/kit/auth/jwt"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func MakeGRPCHandler(server *grpc.Server, endpoint endpoint.Endpoints) {
	logger := zaplog.WithContext(context.Background())
	defer logger.Sync()

	Opts := []kitGRPC.ServerOption{
		kitGRPC.ServerBefore(kitJwt.GRPCToContext()),
	}

	// Register all RPCs with the gRPC server
	newGRPCCalculator := calculatorTransport.NewGRPCCalculator(endpoint, Opts...)

	// Health Check Server
	newGRPCHealth := health.NewServer()
	newGRPCHealth.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	newGRPCHealth.SetServingStatus("health", grpc_health_v1.HealthCheckResponse_SERVING)
	newGRPCHealth.SetServingStatus("calculator", grpc_health_v1.HealthCheckResponse_SERVING)

	// Register the servers with gRPC
	calculator.RegisterCalculatorServer(server, newGRPCCalculator)
	grpc_health_v1.RegisterHealthServer(server, newGRPCHealth)
}

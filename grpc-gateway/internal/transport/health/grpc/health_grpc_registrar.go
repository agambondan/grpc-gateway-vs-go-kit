package healthgrpc

import (
	"git.bluebird.id/firman.agam/grpc-gateway/internal/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type healthGRPCServerRegistrar struct {
	healthServer grpc_health_v1.HealthServer
}

func NewHealthGRPCServerRegistrar(healthServer grpc_health_v1.HealthServer) transport.GRPCService {
	return &healthGRPCServerRegistrar{
		healthServer: healthServer,
	}
}

func (reg *healthGRPCServerRegistrar) Register(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, reg.healthServer)
}

package healthgrpc

import (
	"context"
	"time"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// healthCheckServer implements the healthcheck proto interface
type healthCheckServer struct{}

func NewHealthCheckServer() *healthCheckServer {
	return &healthCheckServer{}
}

func (h *healthCheckServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (h *healthCheckServer) Watch(req *grpc_health_v1.HealthCheckRequest, srv grpc_health_v1.Health_WatchServer) error {
	status := grpc_health_v1.HealthCheckResponse_SERVING
	update := &grpc_health_v1.HealthCheckResponse{Status: status}

	for {
		select {
		case <-srv.Context().Done():
			return nil
		default:
			if err := srv.Send(update); err != nil {
				return err
			}
			time.Sleep(5 * time.Second)
		}
	}
}

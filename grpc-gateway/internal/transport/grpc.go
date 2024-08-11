package transport

import (
	"context"
	"net"

	"git.bluebird.id/firman.agam/grpc-gateway/internal/transport/middleware"
	"git.bluebird.id/promo/packages/zaplog"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCService represents a gRPC service that can be registered with the central server
type GRPCService interface {
	Register(server *grpc.Server)
}

// gRPCServer represents the central gRPC server
type gRPCServer struct {
	server *grpc.Server
	port   string
}

// GRPCServer
// represents interface abstraction every function who implement this can override the logic
type GRPCServer interface {
	RegisterService(service GRPCService)
	Start() error
	StopGracefully()
}

// NewGRPCServer creates a new instance of the central gRPC server
func NewGRPCServer(port string) GRPCServer {
	return &gRPCServer{
		server: grpc.NewServer(
			grpc.UnaryInterceptor(middleware.GRPCMiddleware),
		),
		port: port,
	}
}

// RegisterService registers a gRPC service with the central server
func (s *gRPCServer) RegisterService(service GRPCService) {
	service.Register(s.server)
}

// Start starts the central gRPC server
func (s *gRPCServer) Start() error {
	logger := zaplog.WithContext(context.Background())
	defer logger.Sync()

	logger.Info("Serving gRPC on port :" + s.port)

	lis, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	reflection.Register(s.server)

	return s.server.Serve(lis)
}

func (s *gRPCServer) StopGracefully() {
	s.server.GracefulStop()
}

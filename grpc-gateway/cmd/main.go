package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.bluebird.id/firman.agam/grpc-gateway/internal/transport"
	calculatorgrpc "git.bluebird.id/firman.agam/grpc-gateway/internal/transport/calculator/grpc"
	calculatorhttp "git.bluebird.id/firman.agam/grpc-gateway/internal/transport/calculator/http"
	healthgrpc "git.bluebird.id/firman.agam/grpc-gateway/internal/transport/health/grpc"
	healthhttp "git.bluebird.id/firman.agam/grpc-gateway/internal/transport/health/http"
	"git.bluebird.id/firman.agam/grpc-gateway/internal/utils/env"
	"git.bluebird.id/promo/packages/zaplog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc/health"
)

func main() {
	var err error
	logger := zaplog.WithContext(context.Background())

	healthGrpcServer := healthgrpc.NewHealthCheckServer()

	grpcServer := transport.NewGRPCServer("")

	go func() {
		grpcServer.RegisterService(calculatorgrpc.NewPromotionGRPCServerRegistrar(nil))
		grpcServer.RegisterService(healthgrpc.NewHealthGRPCServerRegistrar(healthGrpcServer))

		err = grpcServer.Start()
		if err != nil {
			logger.Info("failed to start grpc server", zap.Error(err))
		}
	}()

	httpServer := transport.NewHTTPServer(env.HTTPPort, env.GRPCPort)

	go func() {
		err = httpServer.RegisterGRPCGatewayHandler(calculatorhttp.NewPromotionGRPCGatewayRegistrar())
		if err != nil {
			logger.Fatal("failed to register promotion http server", zap.Error(err))
			return
		}

		healthServer := health.NewServer()
		httpServer.RegisterHTTPHandler("/health", healthhttp.NewHealthHandler(healthServer))
		httpServer.RegisterHTTPHandler("/metrics", promhttp.Handler())

		err = httpServer.Start()
		if err != nil {
			logger.Info("http server has been shutdown", zap.Error(err))
		}
	}()

	done := make(chan bool, 1)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		sig := <-sigCh
		logger.Info(fmt.Sprintf("Receive signal %v. Shutting down gracefully...", sig))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		grpcServer.StopGracefully()

		if err = httpServer.Stop(ctx); err != nil {
			logger.Info("failed to stop http server", zap.Error(err))
		}

		done <- true
	}()

	<-done
}

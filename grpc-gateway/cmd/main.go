package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	calculatorservice "git.bluebird.id/firman.agam/grpc-gateway/internal/service/calculator"
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
	"google.golang.org/grpc/health/grpc_health_v1"
)

func init() {
	env.LoadEnv()
}

func main() {
	var err error
	logger := zaplog.WithContext(context.Background())

	healthGrpcServer := healthgrpc.NewHealthCheckServer()

	calculatorService := calculatorservice.NewCalculatorService()
	calculatorGRPC := calculatorgrpc.NewCalculatorGRPCServer(calculatorService)

	grpcServer := transport.NewGRPCServer(env.GRPCPort)

	go func() {
		grpcServer.RegisterService(calculatorgrpc.NewPromotionGRPCServerRegistrar(calculatorGRPC))
		grpcServer.RegisterService(healthgrpc.NewHealthGRPCServerRegistrar(healthGrpcServer))

		err = grpcServer.Start()
		if err != nil {
			logger.Info("failed to start grpc server", zap.Error(err))
		}
	}()

	httpServer := transport.NewHTTPServer(env.HTTPPort, env.GRPCPort)

	go func() {
		httpServer.RegisterHTTPHandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		httpServer.RegisterHTTPHandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "pong from v2")
		})

		err = httpServer.RegisterGRPCGatewayHandler(calculatorhttp.NewPromotionGRPCGatewayRegistrar())
		if err != nil {
			logger.Fatal("failed to register promotion http server", zap.Error(err))
			return
		}

		healthServer := health.NewServer()
		healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
		healthServer.SetServingStatus("health", grpc_health_v1.HealthCheckResponse_SERVING)
		healthServer.SetServingStatus("calculator", grpc_health_v1.HealthCheckResponse_SERVING)
		
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

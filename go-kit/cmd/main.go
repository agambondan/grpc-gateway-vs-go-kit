package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.bluebird.id/firman.agam/go-kit/internal/endpoint"
	calculatorservice "git.bluebird.id/firman.agam/go-kit/internal/service/calculator"
	"git.bluebird.id/firman.agam/go-kit/internal/transport"
	"git.bluebird.id/firman.agam/go-kit/internal/utils/env"
	"git.bluebird.id/promo/packages/zaplog"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	env.LoadEnv()
}

func main() {
	// it shows your line code while print
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logger := zaplog.WithContext(context.Background())

	r := mux.NewRouter()

	// Apply CORS middleware
	allowedOrigins := "*" // Replace with specific origins if needed
	allowedMethods := "GET, POST, PUT, DELETE, OPTIONS"
	allowedHeaders := "Content-Type, Authorization"
	corsHandler := CORS(allowedOrigins, allowedMethods, allowedHeaders)

	httpHandler := corsHandler(r)

	httpServer := &http.Server{
		Handler: httpHandler,
	}

	newCalculatorService := calculatorservice.NewCalculatorService()

	newEndpoints := endpoint.NewEndpoints(newCalculatorService)

	go func() {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "pong")
		})

		v1 := r.PathPrefix("/api/v1").Subrouter()

		transport.MakeHTTPHandler(v1, newEndpoints)

		addr := fmt.Sprintf(":%s", env.HTTPPort)
		httpServer.Addr = addr

		logger.Info("http server running on " + addr)
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Panic("failed to listen and serve", zap.Error(err))
		}
	}()

	grpcServer := grpc.NewServer()

	go func() {
		transport.MakeGRPCHandler(grpcServer, newEndpoints)

		addr := fmt.Sprintf(":%s", env.GRPCPort)
		grpcListener, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Panic("error grpc listener", zap.Error(err))
		}

		logger.Info("grpc server running on " + addr)
		if err := grpcServer.Serve(grpcListener); err != nil && err != grpc.ErrServerStopped {
			logger.Panic("failed to serve grpc", zap.Error(err))
		}
	}()

	done := make(chan bool, 1)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		sig := <-sigCh
		logger.Info(fmt.Sprintf("received signal %v. shutting down gracefully...", sig))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Gracefully stop gRPC server
		grpcServer.GracefulStop()

		// Gracefully shutdown HTTP server
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Info("failed to stop HTTP server", zap.Error(err))
		}

		done <- true
	}()

	<-done
}

// CORS is a middleware function to handle CORS in Go-Kit services.
func CORS(allowedOrigins, allowedMethods, allowedHeaders string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

			// Handle preflight request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

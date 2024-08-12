package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"git.bluebird.id/firman.agam/grpc-gateway/internal/utils"
	"git.bluebird.id/promo/packages/zaplog"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func GRPCMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	ctx = utils.AddRequestIDToGRPCContext(ctx)

	logger := zaplog.WithContext(ctx)

	defer logger.Sync()

	response, err := handler(ctx, req)

	executionTime := time.Since(start).Milliseconds()

	logger.Info("GRPC Request",
		zap.Any("method", info.FullMethod),
		zap.Any("request", req),
		zap.Any("response", response),
		zap.Int64("duration", executionTime),
	)

	return response, err
}

func HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ctx := utils.AddRequestIDToHTTPHeader(context.Background(), r, w)
		logger := zaplog.WithContext(ctx)
		defer logger.Sync()

		r = r.WithContext(ctx)

		// Log the request body if it's not empty
		var requestMap map[string]interface{}
		body, _ := io.ReadAll(r.Body)
		if len(body) > 0 {
			err := json.Unmarshal(body, &requestMap)
			if err != nil {
				logger.Info("failed to unmarshal request", zap.Error(err))
			}

			// Restore the request body
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		// Create a response recorder to capture the response
		recorder := &responseRecorder{ResponseWriter: w, body: &bytes.Buffer{}, status: http.StatusOK}

		params := make(map[string]string)
		for k, v := range r.URL.Query() {
			params[k] = strings.Join(v, ",")
		}

		next.ServeHTTP(recorder, r)

		executionTime := time.Since(start)

		var responseMap map[string]interface{}
		if body := recorder.body.Bytes(); len(body) > 0 {
			_ = json.Unmarshal(body, &responseMap)
		}

		logger.Info("HTTP Request",
			zap.Any("params", params),
			zap.Any("request", requestMap),
			zap.Any("response", responseMap),
			zap.Int("status", recorder.status),
			zap.Int64("execution_time_microseconds", executionTime.Microseconds()),
			zap.Int64("execution_time_milliseconds", executionTime.Milliseconds()),
		)
	})
}

// responseRecorder is a wrapper to capture the response body and status code
type responseRecorder struct {
	http.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (rec *responseRecorder) WriteHeader(status int) {
	rec.status = status
	rec.ResponseWriter.WriteHeader(status)
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.body.Write(b)
	return rec.ResponseWriter.Write(b)
}

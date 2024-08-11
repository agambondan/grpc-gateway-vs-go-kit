package utils

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"git.bluebird.id/promo/packages/zaplog"
)

// getRequestID validates or generates a new request ID
func getRequestID(existingID string) string {
	if existingID == "" {
		return uuid.New().String()
	}
	if _, err := uuid.Parse(existingID); err != nil {
		return uuid.New().String()
	}
	return existingID
}

// addRequestIDToContext adds request ID to context
func addRequestIDToContext(ctx context.Context, requestID string) context.Context {
	return zaplog.NewContext(ctx, zap.String("request_id", requestID))
}

// AddRequestIDToHTTPHeader processes HTTP request, sets request ID, and returns updated context
func AddRequestIDToHTTPHeader(ctx context.Context, r *http.Request, w http.ResponseWriter) context.Context {
	requestID := getRequestID(r.Header.Get("X-Request-Id"))

	// Set request ID in headers
	r.Header.Set("X-Request-Id", requestID)
	w.Header().Add("X-Request-Id", requestID)

	// Add request ID to context
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = addRequestIDToContext(ctx, requestID)

	// Add additional context fields
	return zaplog.NewContext(ctx, zap.String("method", r.Method), zap.String("request_type", r.URL.Path))
}

// AddRequestIDToGRPCContext processes gRPC context, sets request ID, and returns updated context
func AddRequestIDToGRPCContext(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	requestID := ""
	if xRequestIds := md.Get("x-request-id"); len(xRequestIds) > 0 {
		requestID = getRequestID(xRequestIds[0])
	} else {
		requestID = uuid.New().String()
	}

	return addRequestIDToContext(ctx, requestID)
}

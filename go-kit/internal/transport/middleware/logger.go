package middleware

import (
	"context"
	"time"

	"git.bluebird.id/firman.agam/go-kit/internal/utils"
	"git.bluebird.id/promo/packages/zaplog"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

func TransportLogging(requestType string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			now := time.Now()
			requestID, _ := ctx.Value(http.ContextKeyRequestXRequestID).(string)
			requestID = utils.GenerateRequestID(requestID)

			ctx = zaplog.NewContext(ctx, zap.String("request_id", requestID), zap.String("request_type", requestType), zap.String("request_id", requestID))
			
			logger := zaplog.WithContext(ctx)

			resp, err = next(ctx, request)

			logger.Info("HTTP Request - "+requestType,
				zap.Any("request", request),
				zap.Any("response", resp),
				zap.Error(err),
				zap.String("execution_time", time.Since(now).String()),
			)
			return
		}
	}
}

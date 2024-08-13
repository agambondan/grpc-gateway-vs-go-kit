package middleware

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const apiKey = "Contoh API Key"

func APIKeyGRPCAuth() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "API key missing")
			}

			apiKeys := md["api-key"]
			if len(apiKeys) == 0 || apiKeys[0] != apiKey {
				return nil, status.Errorf(codes.PermissionDenied, "Invalid API key")
			}

			resp, err = next(ctx, request)
			return
		}
	}
}

package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.bluebird.id/promo/packages/zaplog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func CustomMetadata(ctx context.Context, r *http.Request) metadata.MD {
	md := make(metadata.MD)
	queryParams := r.URL.Query()
	for key, values := range queryParams {
		md[key] = values
	}
	return md
}

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	ctx = AddRequestIDToHTTPHeader(ctx, req, w)
	logger := zaplog.WithContext(ctx)

	defer logger.Sync()

	var response Error
	code := 400
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		logger.Info("Failed to extract ServerMetadata from context")
	}
	for k, vs := range md.HeaderMD {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	if se, ok := status.FromError(err); ok {
		code = runtime.HTTPStatusFromCode(se.Code())
		errorMessages := strings.Split(se.Message(), "|")
		response.Message = errorMessages[0]

		if len(errorMessages) > 1 {
			response.Message = errorMessages[1]
			response.LocalizedMessage.English = errorMessages[2]
			response.LocalizedMessage.Indonesia = errorMessages[3]
		}
	} else {
		response.Message = err.Error()
	}
	response.Code = code
	response.CodeStr = strconv.Itoa(code)
	jsonResp, _ := json.Marshal(response)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResp)
}

func CustomResponseHandler(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	w.Header().Del("Grpc-Metadata-Content-Type")
	return nil
}

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "X-Request-Id":
		return key, true
	case "X-User-Id":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

package decoder

import (
	"context"
	"encoding/json"
	"net/http"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func HTTPMultiply(_ context.Context, r *http.Request) (interface{}, error) {
	req := new(calculator.MultiplyRequest)
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func GRPCMultiply(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calculator.MultiplyRequest)
	return req, nil
}

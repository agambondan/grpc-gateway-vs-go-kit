package decoder

import (
	"context"
	"encoding/json"
	"net/http"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func HTTPDivide(_ context.Context, r *http.Request) (interface{}, error) {
	req := new(calculator.DivideRequest)
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func GRPCDivide(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calculator.DivideRequest)
	return req, nil
}

package decoder

import (
	"context"
	"encoding/json"
	"net/http"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func HTTPSubtract(_ context.Context, r *http.Request) (interface{}, error) {
	req := new(calculator.SubtractRequest)
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func GRPCSubtract(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calculator.SubtractRequest)
	return req, nil
}

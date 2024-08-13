package decoder

import (
	"context"
	"encoding/json"
	"net/http"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func HTTPAdd(_ context.Context, r *http.Request) (interface{}, error) {
	req := new(calculator.AddRequest)
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func GRPCAdd(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calculator.AddRequest)
	return req, nil
}
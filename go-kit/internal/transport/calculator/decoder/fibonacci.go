package decoder

import (
	"context"
	"net/http"
	"strconv"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
	"github.com/gorilla/mux"
)

func HTTPFibonacci(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])
	if err != nil {
		return nil, err
	}
	return &calculator.FibonacciRequest{N: int32(n)}, nil
}

func GRPCFibonacci(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calculator.FibonacciRequest)
	return req, nil
}

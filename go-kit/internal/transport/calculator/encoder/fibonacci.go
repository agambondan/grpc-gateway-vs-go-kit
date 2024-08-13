package encoder

import (
	"context"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func GRPCFibonacci(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*calculator.FibonacciResponse)
	return res, nil
}

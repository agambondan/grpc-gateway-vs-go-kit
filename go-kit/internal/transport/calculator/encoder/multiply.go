package encoder

import (
	"context"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func GRPCMultiply(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*calculator.MultiplyResponse)
	return res, nil
}

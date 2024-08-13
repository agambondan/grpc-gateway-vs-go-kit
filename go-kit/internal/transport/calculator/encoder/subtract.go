package encoder

import (
	"context"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func GRPCSubtract(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*calculator.SubtractResponse)
	return res, nil
}

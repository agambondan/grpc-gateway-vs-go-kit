package encoder

import (
	"context"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func GRPCDivide(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*calculator.DivideResponse)
	return res, nil
}

package encoder

import (
	"context"

	calculator "git.bluebird.id/firman.agam/go-kit/pkg/proto/gen"
)

func GRPCAdd(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*calculator.AddResponse)
	return res, nil
}

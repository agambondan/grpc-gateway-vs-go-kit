package middleware

import (
	"context"
	"encoding/base64"
	"strings"

	"git.bluebird.id/firman.agam/go-kit/internal/utils"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
)

var user, password = "admin", "admin"

func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}

	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}

	pair := strings.SplitN(string(c), ":", 2)
	if len(pair) != 2 {
		return
	}

	return pair[0], pair[1], true
}

func BasicAuth() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			auth, ok := ctx.Value(http.ContextKeyRequestAuthorization).(string)
			if !ok {
				return nil, utils.ErrAuth
			}

			givenUser, givenPassword, ok := parseBasicAuth(auth)
			if !ok {
				return nil, utils.ErrAuth
			} else if givenUser != user || givenPassword != password {
				return nil, utils.ErrAuth
			}

			resp, err = next(ctx, request)
			return
		}
	}
}

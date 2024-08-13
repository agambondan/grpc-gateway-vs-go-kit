package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"git.bluebird.id/firman.agam/go-kit/internal/utils"
	"git.bluebird.id/promo/packages/config"
	"github.com/go-kit/kit/endpoint"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/golang-jwt/jwt/v4"
)

var bbOneBaseUrl = config.GetString("BBONE_AUTH_BASE_URL", "https://devapi-bbone.bluebird.id")

func TokenAuth() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			auth, ok := ctx.Value(kitHttp.ContextKeyRequestAuthorization).(string)
			if !ok {
				return nil, utils.ErrAuth
			}
			token := strings.ReplaceAll(auth, "Bearer ", "")
			claims, ok := ExtractClaims(token)
			if !ok {
				return nil, utils.ErrAuth
			}

			if claims.Valid() != nil {
				return nil, utils.ErrTokenExpired
			}

			uid, ok := claims["uid"].(string)
			if !ok {
				return nil, utils.ErrAuth
			}

			err = ValidateToken(ctx, token, uid)
			if err != nil {
				return nil, utils.ErrAuth
			}

			resp, err = next(ctx, request)
			return
		}
	}
}

// ExtractClaims ...
func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenStr, nil)
	if err != nil && err.Error() != "no Keyfunc was provided." {
		return nil, false
	}
	if token != nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims, true
		}
	}

	return nil, false
}

func ValidateToken(ctx context.Context, token, uid string) error {
	s, _ := json.Marshal(map[string]interface{}{
		"user_id": uid,
		"token":   token,
	})

	req, err := http.NewRequest(http.MethodPost, bbOneBaseUrl+"/token/validate", bytes.NewReader(s))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	valid, ok := result["valid"].(bool)
	if !valid || !ok {
		return utils.ErrAuth
	}

	return nil
}

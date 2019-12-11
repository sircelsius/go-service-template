package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/sircelsius/go-service-template/internal/logging"
)

func (s *server) authenticationMiddleware(jwkUri string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO cache JWK instead of fetching them for every request
			set, err := jwk.Fetch(jwkUri)
			if err != nil {
				panic(err)
			}

			ctx := r.Context()
			token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (i interface{}, e error) {
				signingKey, err := set.Keys[0].Materialize()
				if err != nil {
					return nil, err
				}
				return signingKey, nil
			})
			if err != nil {
				if err == request.ErrNoTokenInRequest {
					logging.GetLogger(ctx).Info("No token found in request, proceeding")
				} else {
					logging.GetLogger(ctx).Info(fmt.Sprintf("Unable to parse token: %v", err.Error()))
				}
			}
			if token != nil && token.Valid {
				logging.GetLogger(ctx).Info("Valid token successfully extracted.")
				ctx = context.WithValue(ctx, "claims", token.Claims)
				ctx = context.WithValue(ctx, "authenticated", true)
				ctx = logging.ContextWithLogger(ctx, nil)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}


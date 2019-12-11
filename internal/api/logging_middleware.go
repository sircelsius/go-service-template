package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sircelsius/go-service-template/internal/logging"
	"go.uber.org/zap"
)

// loggingMiddleware decorates the request logger with mux route names and HTTP methods
// and decorates the request logger to include these two values.
func (s *server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logging.GetLogger(ctx)
		routeName := mux.CurrentRoute(r).GetName()
		method := r.Method
		requestLogger := logger.With(
			zap.String("route", routeName),
			zap.String("method", method),
		)

		ctx = logging.ContextWithLogger(ctx, requestLogger)
		ctx = context.WithValue(ctx, "route", routeName)
		ctx = context.WithValue(ctx, "method", method)

		logging.GetLogger(ctx).Info("request started")
		next.ServeHTTP(w, r.WithContext(ctx))
		logging.GetLogger(ctx).Info("request completed")
	})
}

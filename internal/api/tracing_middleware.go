package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sircelsius/go-service-template/internal/logging"
	"github.com/sircelsius/go-service-template/internal/tracing"
)

func (s *server) tracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var span opentracing.Span
		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			logging.GetLogger(r.Context()).Error(fmt.Sprintf("unable to extract trace information from wire, a new trace with no parent will be generated: %v", err))
		}

		appSpecificOperationName := mux.CurrentRoute(r).GetName()

		span = opentracing.StartSpan(
			appSpecificOperationName,
			ext.RPCServerOption(wireContext))
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(r.Context(), span)
		ctx = tracing.ContextWithSpanLogger(ctx, span)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

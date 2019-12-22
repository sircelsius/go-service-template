package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go/config"
)

func TestTracingMiddleware_WithoutOpentracingHeaders(t *testing.T) {
	c, _ := config.FromEnv()
	closer, _ := c.InitGlobalTracer("test",)
	defer closer.Close()

	s := server{router: mux.NewRouter()}
	h := func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			span := opentracing.SpanFromContext(r.Context())
			defer span.Finish()
			assert.NotEmpty(t, span)
		}
	}()

	s.router.Handle("/test", h).Name("test")
	s.router.Use(s.tracingMiddleware)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)
}

func TestTracingMiddleware_WithOpentracingHeaders(t *testing.T) {
	c, _ := config.FromEnv()
	closer, _ := c.InitGlobalTracer("test",)
	defer closer.Close()

	s := server{router: mux.NewRouter()}
	h := func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			span := opentracing.SpanFromContext(r.Context())
			defer span.Finish()
			assert.NotEmpty(t, span)
			assert.Equal(t, "bar", span.BaggageItem("foo"))
		}
	}()

	s.router.Handle("/test", h).Name("test")
	s.router.Use(s.tracingMiddleware)

	span := opentracing.StartSpan("foo")
	defer span.Finish()

	span.SetBaggageItem("foo", "bar")
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)

	_ = opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	rr := httptest.NewRecorder()

	s.router.ServeHTTP(rr, req)
}
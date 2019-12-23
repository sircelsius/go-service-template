package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *server) metricsMiddleware() func(next http.Handler) http.Handler {
	var summary = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_server",
		Help: "HTTP server summary",
	}, []string{"handler_name", "status_code", "http_method"})
	_ = prometheus.Register(summary)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				duration := time.Since(start)
				summary.WithLabelValues(mux.CurrentRoute(r).GetName(), "UNKNOWN", r.Method).Observe(duration.Seconds())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

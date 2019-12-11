package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *server) metricsHandler() http.Handler {
	return promhttp.Handler()
}

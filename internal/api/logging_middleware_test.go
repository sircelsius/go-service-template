package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware(t *testing.T) {
	s := server{router: mux.NewRouter()}
	h := func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			assert.NotEmpty(t, r.Context().Value("route"))
		}
	}()

	s.router.Handle("/test", h).Name("test")
	s.router.Use(s.loggingMiddleware)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)
}

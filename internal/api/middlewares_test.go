package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMiddlwares(t *testing.T) {
	s := server{router: mux.NewRouter()}
	s.middlewares()

	h := func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// here we only assert that the logging middleware has corectly been added.
			// middlewares themselves should be tested separately.
			assert.NotEmpty(t, r.Context().Value("route"))
		}
	}()

	s.router.Handle("/test", h).Name("test")

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)
}

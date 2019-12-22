package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	s := server{router: mux.NewRouter()}
	s.routes()

	req, _ := http.NewRequest(http.MethodGet, "/_system/health", nil)
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

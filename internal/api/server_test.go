package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	s := NewServer()

	req, _ := http.NewRequest(http.MethodGet, "/_system/health", nil)
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)

	// We only test that the health endpoint has been added here for the sake of simplicity.
	assert.Equal(t, http.StatusOK, rr.Code)
}

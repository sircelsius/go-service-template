package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/foo", nil)
	rr := httptest.NewRecorder()
	s := server{nil}
	s.healthHandler().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"status":"healthy!"}`
	actual := rr.Body.String()
	assert.Equal(t, expected, actual)
}

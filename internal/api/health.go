package api

import (
	"encoding/json"
	"net/http"
)

// healthHandler is a function that returns a handler func.
func (s *server) healthHandler() http.HandlerFunc {
	type healthResponse struct {
		Status string `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(healthResponse{"healthy!"})
		w.Write(b)
	}
}
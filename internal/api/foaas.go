package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sircelsius/go-service-template/internal/foaas"
	"github.com/sircelsius/go-service-template/internal/logging"
)

func (s *server) foaasHandler() http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	client := foaas.NewClient()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message, err := client.Cool(r.Context(), "sircelsius")
		if err != nil {
			logging.GetLogger(r.Context()).Error(fmt.Sprintf("error while getting foaas: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			b, _ := json.Marshal(response{"oops"})
			w.Write(b)
			return
		}

		b, _ := json.Marshal(response{message})
		w.Write(b)
	})
}

package api

import (
	"github.com/gorilla/mux"
)

// server is the struct that contains our HTTP server
type server struct{
	router *mux.Router
}

// NewServer returns a HTTP server that registers its routes and middlewares.
func NewServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}
	s.middlewares()
	s.routes()
	return s
}

func (s *server) GetRouter() *mux.Router {
	return s.router
}
package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	listenPort string
}

func NewServer(listenPort string) *Server {
	return &Server{listenPort}
}

func (s *Server) Start() error {
	r := chi.NewRouter()

	r.Mount("/api", apiResource{
		useCORS: true,
	}.Routes())

	return http.ListenAndServe(":"+s.listenPort, r)
}

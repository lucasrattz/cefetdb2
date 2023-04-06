package api

import (
	"cefetdb2api/storage"
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
	c := storage.NewDriveClient()

	storer := storage.NewDriveStorer(c)

	r.Mount("/api", apiResource{
		storer:  storer,
		useCORS: true,
	}.Routes())

	return http.ListenAndServe(":"+s.listenPort, r)
}

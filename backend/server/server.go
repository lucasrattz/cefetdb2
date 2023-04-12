package api

import (
	"cefetdb2api/config"
	"cefetdb2api/storage"
	"cefetdb2api/types"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	cfg         *config.Config
	credentials types.OAuthCredentials
}

func NewServer(cfg *config.Config, credentials types.OAuthCredentials) *Server {
	return &Server{cfg, credentials}
}

func (s *Server) Start() error {
	r := chi.NewRouter()
	c := storage.NewDriveClient(s.credentials)

	storer := storage.NewDriveStorer(c, s.cfg)

	r.Use(middleware.Heartbeat("/health"))

	r.Mount("/api", apiResource{
		storer: storer,
		cfg:    s.cfg,
	}.Routes())

	return http.ListenAndServe(":"+s.cfg.Server.Port, r)
}

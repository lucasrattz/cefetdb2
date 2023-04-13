package api

import (
	"context"
	"net/http"

	"cefetdb2api/config"
	"cefetdb2api/types"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type apiResource struct {
	storer types.Storer
	cfg    *config.Config
}

func (ar apiResource) Routes() chi.Router {
	r := chi.NewRouter()

	if ar.cfg.Server.UseCors {
		corsHandler := cors.New(cors.Options{
			AllowedOrigins: []string{ar.cfg.Server.Host + ":" + ar.cfg.Server.Port},
			AllowedMethods: []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Content-Type"},
			MaxAge:         300,
		}).Handler
		r.Use(corsHandler)
	}

	r.Get("/semesters", ar.getSemesters().ServeHTTP)
	r.Get("/disciplines", ar.getAllDisciplines().ServeHTTP)
	r.Get("/files", ar.getAllFiles().ServeHTTP)

	r.Route("/semesters/{semesterID}", func(r chi.Router) {
		r.Use(ar.semesterMiddleware)
		r.Get("/disciplines", ar.getDisciplinesBySemester().ServeHTTP)
		r.Route("/disciplines/{disciplineID}", func(r chi.Router) {
			r.Use(ar.disciplineMiddleware)
			r.Get("/files", ar.getFilesByDiscipline().ServeHTTP)
			r.Post("/files", ar.uploadFile().ServeHTTP)
		})
	})

	return r
}

var (
	contextKeySemesterID   = types.ContextKey("semesterID")
	contextKeyDisciplineID = types.ContextKey("disciplineID")
)

// Middleware to get semester ID from URL
func (ar apiResource) semesterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		semesterID := chi.URLParam(r, contextKeySemesterID.Value())
		if semesterID == "" {
			http.Error(w, "semester ID is required", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeySemesterID, semesterID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Middleware to get discipline ID from URL
func (ar apiResource) disciplineMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		disciplineID := chi.URLParam(r, contextKeyDisciplineID.Value())
		if disciplineID == "" {
			http.Error(w, "discipline ID is required", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyDisciplineID, disciplineID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

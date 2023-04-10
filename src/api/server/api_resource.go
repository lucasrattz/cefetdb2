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
			AllowedOrigins: []string{ar.cfg.Server.Host + ar.cfg.Server.Port},
			AllowedMethods: []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Content-Type"},
			MaxAge:         300,
		}).Handler
		r.Use(corsHandler)
	}

	r.Get("/semesters", ar.getSemesters)
	r.Get("/disciplines", ar.getAllDisciplines)
	r.Get("/files", ar.getAllFiles().ServeHTTP)

	r.Route("/semesters/{semesterID}", func(r chi.Router) {
		r.Use(ar.semesterMiddleware)
		r.Get("/disciplines", ar.getDisciplinesBySemester)
		r.Get("/{disciplineID}/files", ar.getFilesByDiscipline)
		r.Post("/{disciplineID}/files", ar.uploadFile)
	})

	return r
}

var (
	contextKeySemesterID = types.ContextKey("semesterID")
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

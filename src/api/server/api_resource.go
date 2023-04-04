package server

import (
	"context"
	"net/http"

	"cefetdb2api/types"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type apiResource struct {
	useCORS bool
}

func (ar apiResource) Routes() chi.Router {
	r := chi.NewRouter()

	if ar.useCORS {
		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost:3500"},
			AllowedMethods: []string{"GET", "POST"},
			MaxAge:         300,
		})
		r.Use(corsMiddleware.Handler)
	}

	r.Get("/semesters", ar.getSemesters)
	r.Get("/disciplines", ar.getAllDisciplines)
	r.Get("/files", ar.getAllFiles)

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

func (ar apiResource) getSemesters(w http.ResponseWriter, r *http.Request)      {}
func (ar apiResource) getAllDisciplines(w http.ResponseWriter, r *http.Request) {}
func (ar apiResource) getAllFiles(w http.ResponseWriter, r *http.Request)       {}

func (ar apiResource) getDisciplinesBySemester(w http.ResponseWriter, r *http.Request) {}
func (ar apiResource) getFilesByDiscipline(w http.ResponseWriter, r *http.Request)     {}
func (ar apiResource) uploadFile(w http.ResponseWriter, r *http.Request)               {}

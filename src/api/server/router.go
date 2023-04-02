package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type apiRouter struct{}

func (ar apiRouter) Routes() chi.Router {
	r := chi.NewRouter()

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

func (ar apiRouter) semesterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		semesterID := chi.URLParam(r, "semesterID")
		if semesterID == "" {
			http.Error(w, "semester ID is required", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "semesterID", semesterID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ar apiRouter) getSemesters(w http.ResponseWriter, r *http.Request)      {}
func (ar apiRouter) getAllDisciplines(w http.ResponseWriter, r *http.Request) {}
func (ar apiRouter) getAllFiles(w http.ResponseWriter, r *http.Request)       {}

func (ar apiRouter) getDisciplinesBySemester(w http.ResponseWriter, r *http.Request) {}
func (ar apiRouter) getFilesByDiscipline(w http.ResponseWriter, r *http.Request)     {}
func (ar apiRouter) uploadFile(w http.ResponseWriter, r *http.Request)               {}

package api

import (
	"context"
	"encoding/json"
	"net/http"

	"cefetdb2api/types"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type apiResource struct {
	storer  types.Storer
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

//
// Handlers
//

// GET /semesters
func (ar apiResource) getSemesters(w http.ResponseWriter, r *http.Request) {}

// GET /disciplines
func (ar apiResource) getAllDisciplines(w http.ResponseWriter, r *http.Request) {}

// GET /files
func (ar apiResource) getAllFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	files, err := ar.storer.GetAllFiles(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(files)
}

// GET /semesters/{semesterID}/disciplines
func (ar apiResource) getDisciplinesBySemester(w http.ResponseWriter, r *http.Request) {}

// GET /semesters/{semesterID}/{disciplineID}/files
func (ar apiResource) getFilesByDiscipline(w http.ResponseWriter, r *http.Request) {}

// POST /semesters/{semesterID}/{disciplineID}/files
func (ar apiResource) uploadFile(w http.ResponseWriter, r *http.Request) {}

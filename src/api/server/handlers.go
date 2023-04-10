package api

import (
	"cefetdb2api/types"
	"encoding/json"
	"net/http"
)

// GET /semesters
func (ar apiResource) getSemesters(w http.ResponseWriter, r *http.Request) {}

// GET /disciplines
func (ar apiResource) getAllDisciplines(w http.ResponseWriter, r *http.Request) {}

// GET /files
func (ar apiResource) getAllFiles() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := r.Context()
		defer r.Body.Close()

		files, err := ar.storer.GetAllFiles(ctx)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(files)
	}
	return http.HandlerFunc(fn)
}

// GET /semesters/{semesterID}/disciplines
func (ar apiResource) getDisciplinesBySemester(w http.ResponseWriter, r *http.Request) {}

// GET /semesters/{semesterID}/{disciplineID}/files
func (ar apiResource) getFilesByDiscipline(w http.ResponseWriter, r *http.Request) {}

// POST /semesters/{semesterID}/{disciplineID}/files
func (ar apiResource) uploadFile(w http.ResponseWriter, r *http.Request) {}

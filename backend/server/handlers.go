package api

import (
	"cefetdb2api/types"
	"encoding/json"
	"errors"
	"net/http"
)

// GET /semesters
func (ar apiResource) getSemesters() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := r.Context()
		defer r.Body.Close()

		semesters, err := ar.storer.GetSemesters(ctx)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(semesters)
	}
	return http.HandlerFunc(fn)
}

// GET /disciplines
func (ar apiResource) getAllDisciplines() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := r.Context()
		defer r.Body.Close()

		disciplines, err := ar.storer.GetAllDisciplines(ctx)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(disciplines)
	}
	return http.HandlerFunc(fn)
}

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
func (ar apiResource) getDisciplinesBySemester() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := r.Context()
		defer r.Body.Close()

		semesterID := ctx.Value(contextKeySemesterID).(string)

		disciplines, err := ar.storer.GetDisciplinesBySemester(ctx, semesterID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(disciplines)
	}
	return http.HandlerFunc(fn)
}

// GET /semesters/{semesterID}/{disciplineID}/files
func (ar apiResource) getFilesByDiscipline() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := r.Context()
		defer r.Body.Close()

		disciplineID := ctx.Value(contextKeyDisciplineID).(string)

		files, err := ar.storer.GetFilesByDiscipline(ctx, disciplineID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(files)
	}
	return http.HandlerFunc(fn)
}

// POST /semesters/{semesterID}/{disciplineID}/files
func (ar apiResource) uploadFile() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx := r.Context()

		disciplineID := ctx.Value(contextKeyDisciplineID).(string)

		if err := r.ParseMultipartForm(5 << 20); err != nil { // Store up to 5MB in memory
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		file, mHeader, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		header := make([]byte, 512)

		if _, err := file.Read(header); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		if _, err := file.Seek(0, 0); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}

		fileMimeType := http.DetectContentType(header)

		f := types.NewUploadableFile(file, mHeader.Filename, mHeader.Size, fileMimeType, disciplineID)

		if err := f.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(types.ErrorResponse{
				ErrorMessage: errors.Join(err...).Error(),
			})
			return
		}

		defer file.Close()
		res, err := ar.storer.UploadFile(ctx, f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(types.ErrorResponse{ErrorMessage: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(res)
	}

	return http.HandlerFunc(fn)
}

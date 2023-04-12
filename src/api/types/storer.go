package types

import (
	"context"
)

type Storer interface {
	GetSemesters(context.Context) ([]FolderResponse, error)
	GetAllDisciplines(context.Context) ([]FolderResponse, error)
	GetAllFiles(context.Context) ([]FileResponse, error)
	GetDisciplinesBySemester(context.Context, string) ([]FolderResponse, error)
	GetFilesByDiscipline(context.Context, string) ([]FileResponse, error)
	UploadFile(context.Context, UploadableFile) (FileResponse, error)
}

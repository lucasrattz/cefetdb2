package types

import "context"

type Storer interface {
	//GetAllSemesters(context.Context) ([]FolderResponse, error)
	//GetAllDisciplines(context.Context) ([]FolderResponse, error)
	GetAllFiles(context.Context) ([]FileResponse, error)
	//GetDisciplinesBySemester(context.Context) ([]FolderResponse, error)
	//GetDisciplineFiles(context.Context) ([]FileResponse, error)
	//Upload(context.Context) (FileResponse, error)
}

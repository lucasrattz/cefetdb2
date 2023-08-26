package storage

import (
	"cefetdb2api/config"
	"cefetdb2api/types"
	"context"

	"google.golang.org/api/drive/v3"
)

type DriveStorer struct {
	driveClient *DriveClient
	cfg         *config.Config
}

func NewDriveStorer(driveClient *DriveClient, cfg *config.Config) *DriveStorer {
	return &DriveStorer{driveClient: driveClient, cfg: cfg}
}

func (d DriveStorer) GetAllFiles(ctx context.Context) ([]types.FileResponse, error) {
	srv, err := d.driveClient.GetDriveService(ctx, d.cfg)
	if err != nil {
		return nil, err
	}

	r, err := srv.Files.List().Q("mimeType!=" + types.FolderMimeType).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return nil, err
	}

	var files []types.FileResponse

	for _, i := range r.Files {
		files = append(files, types.FileResponse{ID: i.Id, Name: i.Name})
	}

	return files, nil
}

func (d DriveStorer) GetSemesters(ctx context.Context) ([]types.FolderResponse, error) {
	srv, err := d.driveClient.GetDriveService(ctx, d.cfg)
	if err != nil {
		return nil, err
	}

	r, err := srv.Files.List().
		Q("properties has {key='type' and value='Semester'}").Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return nil, err
	}

	var files []types.FolderResponse

	for _, i := range r.Files {
		files = append(files, types.FolderResponse{ID: i.Id, Name: i.Name, Type: "Semester"})
	}

	return files, nil
}

func (d DriveStorer) GetAllDisciplines(ctx context.Context) ([]types.FolderResponse, error) {
	srv, err := d.driveClient.GetDriveService(ctx, d.cfg)
	if err != nil {
		return nil, err
	}

	r, err := srv.Files.List().
		Q("properties has {key='type' and value='Discipline'}").Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return nil, err
	}

	var files []types.FolderResponse

	for _, i := range r.Files {
		files = append(files, types.FolderResponse{ID: i.Id, Name: i.Name, Type: "Discipline"})
	}

	return files, nil
}

func (d DriveStorer) GetDisciplinesBySemester(ctx context.Context, semesterID string) ([]types.FolderResponse, error) {
	srv, err := d.driveClient.GetDriveService(ctx, d.cfg)
	if err != nil {
		return nil, err
	}

	var query = "properties has {key='type' and value='Discipline'} and properties has {key='semester' and value='" + semesterID + "'}"
	r, err := srv.Files.List().Q(query).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return nil, err
	}

	var files []types.FolderResponse

	for _, i := range r.Files {
		files = append(files, types.FolderResponse{ID: i.Id, Name: i.Name, Type: "Discipline"})
	}

	return files, nil
}

func (d DriveStorer) GetFilesByDiscipline(ctx context.Context, disciplineID string) ([]types.FileResponse, error) {
	srv, err := d.driveClient.GetDriveService(ctx, d.cfg)
	if err != nil {
		return nil, err
	}

	var query = "properties has {key='type' and value='Discipline'} and name='" + disciplineID + "'"
	f, err := srv.Files.List().Q(query).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return nil, err
	}

	parentFolderID := f.Files[0].Id
	r, err := srv.Files.List().
		Q("'" + parentFolderID + "' in parents").Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return nil, err
	}

	var files []types.FileResponse

	for _, i := range r.Files {
		files = append(files, types.FileResponse{ID: i.Id, Name: i.Name})
	}

	return files, nil
}

func (d DriveStorer) UploadFile(ctx context.Context, f types.UploadableFile) (types.FileResponse, error) {
	srv, err := d.driveClient.GetDriveService(ctx, d.cfg)
	if err != nil {
		return types.FileResponse{}, err
	}

	var query = "properties has {key='type' and value='Discipline'} and name='" + f.ParentName + "'"
	p, err := srv.Files.List().Q(query).Fields("nextPageToken, files(id)").Do()
	if err != nil {
		return types.FileResponse{}, err
	}

	r, err := srv.Files.Create(&drive.File{
		Name:    f.Name,
		Parents: []string{p.Files[0].Id},
	}).ResumableMedia(ctx, f.File, f.Size, f.MimeType).Do()
	if err != nil {
		return types.FileResponse{}, err
	}

	return types.FileResponse{ID: r.Id, Name: r.Name}, nil
}

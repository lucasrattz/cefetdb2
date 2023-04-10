package storage

import (
	"cefetdb2api/config"
	"cefetdb2api/types"
	"context"
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

// func (d DriveStorer) GetSemesters(ctx context.Context) ([]types.FolderResponse, error) {
// 	srv, err := d.driveClient.GetDriveService(ctx, d.cfg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r, err := srv.Files.List().Q("mimeType!=" + types.FolderMimeType).Fields("nextPageToken, files(id, name)").Do()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var files []types.FolderResponse

// 	for _, i := range r.Files {
// 		files = append(files, types.FolderResponse{ID: i.Id, Name: i.Name})
// 	}

// 	return files, nil
// }

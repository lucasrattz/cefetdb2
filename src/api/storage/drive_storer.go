package storage

import (
	"cefetdb2api/types"
	"context"
)

type DriveStorer struct {
	driveClient *DriveClient
}

func NewDriveStorer(driveClient *DriveClient) *DriveStorer {
	return &DriveStorer{}
}

func (d DriveStorer) GetAllFiles(ctx context.Context) ([]types.FileResponse, error) {
	srv, err := d.driveClient.GetDriveService(ctx)
	if err != nil {
		return nil, err
	}

	r, err := srv.Files.List().Fields("nextPageToken, files(id, name, mimeType)").Do()
	if err != nil {
		return nil, err
	}

	var files []types.FileResponse

	for _, i := range r.Files {
		if !types.MimeType(i.MimeType).IsFolder() {
			files = append(files, types.FileResponse{ID: i.Id, Name: i.Name})
		}
	}

	return files, nil
}

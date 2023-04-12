package types

import (
	"errors"
	"io"
)

type UploadableFile struct {
	File       io.ReaderAt
	Name       string
	Size       int64
	MimeType   string
	ParentName string
}

var permittedMimeTypes = []string{
	"application/pdf",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/msword",
	"application/vnd.oasis.opendocument.text",
	"application/rtf",
}

func NewUploadableFile(file io.ReaderAt, name string, size int64, mimeType string, parent string) UploadableFile {
	return UploadableFile{
		File:       file,
		Name:       name,
		Size:       size,
		MimeType:   mimeType,
		ParentName: parent,
	}
}

func isPermittedMimeType(mimeType string) bool {
	for _, i := range permittedMimeTypes {
		if i == mimeType {
			return true
		}
	}

	return false
}

func (f UploadableFile) Validate() []error {
	var errs []error

	if f.File == nil {
		errs = append(errs, errors.New("file cannot be nil"))
	}

	if f.Name == "" {
		errs = append(errs, errors.New("file name cannot be blank"))
	}

	if f.Size > 5<<20 {
		errs = append(errs, errors.New("file size above the 5MB limit"))
	}

	if !isPermittedMimeType(f.MimeType) {
		errs = append(errs, errors.New("file type not permitted"))
	}

	if f.ParentName == "" {
		errs = append(errs, errors.New("parent name cannot be blank"))
	}

	return errs
}

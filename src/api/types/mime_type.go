package types

type MimeType string

func (m MimeType) IsFolder() bool {

	folderMimeType := "application/vnd.google-apps.folder"

	return string(m) == folderMimeType
}

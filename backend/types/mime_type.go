package types

type MimeType string

const (
	FolderMimeType string = "'application/vnd.google-apps.folder'"
	FileMimeType   string = "'application/vnd.google-apps.file'"
)

func (m MimeType) IsFolder() bool {
	return string(m) == FolderMimeType
}

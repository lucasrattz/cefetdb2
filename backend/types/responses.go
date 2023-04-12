package types

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

type FileResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FolderResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

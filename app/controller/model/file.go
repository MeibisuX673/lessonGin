package model

type CreateFile struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type FileResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

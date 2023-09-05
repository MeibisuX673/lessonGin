package model

type AlbumCreate struct {
	Title    string  `json:"title" validate:"required"`
	ArtistID uint    `json:"artistID" validate:"required,number"`
	Price    float64 `json:"price" validate:"required,number"`
	FileId   *uint   `json:"file_id"`
}

type AlbumResponse struct {
	Id       uint          `json:"id"`
	Title    string        `json:"title"`
	Price    float64       `json:"price"`
	ArtistID uint          `json:"artistID"`
	File     *FileResponse `json:"file,omitempty"`
}

type UpdateAlbum struct {
	Title    *string  `json:"title"`
	Price    *float64 `json:"price"`
	ArtistID *uint    `json:"artist_id"`
	FileID   *uint    `json:"file_id"`
}

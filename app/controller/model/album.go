package model

type CreateAlbum struct {
	Title    string  `json:"title" binding:"required"`
	ArtistID uint    `json:"artistID" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	FileId   *uint   `json:"file_id"`
}

type ResponseAlbum struct {
	Id    uint          `json:"id"`
	Title string        `json:"title"`
	Price float64       `json:"price"`
	File  *ResponseFile `json:"file"`
}

type UpdateAlbum struct {
	Title    *string  `json:"title"`
	Price    *float64 `json:"price"`
	ArtistID *uint    `json:"artist_id"`
	FileID   *uint    `json:"file_id"`
}

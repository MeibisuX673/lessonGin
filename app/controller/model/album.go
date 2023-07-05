package model

type CreateAlbum struct {
	Title    string  `json:"title" binding:"required"`
	ArtistID int     `json:"artistID" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
}

type ResponseAlbum struct {
	Id     int            `json:"id"`
	Title  string         `json:"title"`
	Artist ResponseArtist `json:"artist"`
	Price  float64        `json:"price"`
}

type UpdateAlbum struct {
	Title    *string  `json:"title"`
	Price    *float64 `json:"price"`
	ArtistID *int     `json:"artist_id"`
}

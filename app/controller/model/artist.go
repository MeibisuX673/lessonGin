package model

type ArtistCreate struct {
	Name     string `json:"name" validate:"required"`
	Age      uint   `json:"age" validate:"min=18,max=120,number,required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password"`
	FileIds  []uint `json:"fileIds"`
}

type ArtistResponse struct {
	ID     uint            `json:"id"`
	Name   string          `json:"name"`
	Age    uint            `json:"age"`
	Files  []FileResponse  `json:"files"`
	Albums []AlbumResponse `json:"albums"`
}

type UpdateArtist struct {
	Name *string `json:"name"`
	Age  *uint   `json:"age"`
}

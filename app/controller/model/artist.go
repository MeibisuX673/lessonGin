package model

type CreateArtist struct {
	Name     string `json:"name" validate:"required"`
	Age      uint   `json:"age" validate:"min=18,max=120,number,required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password"`
	FileIds  []uint `json:"fileIds"`
}

type ResponseArtist struct {
	ID     uint            `json:"id"`
	Name   string          `json:"name"`
	Age    uint            `json:"age"`
	Files  []ResponseFile  `json:"files"`
	Albums []ResponseAlbum `json:"albums"`
}

type UpdateArtist struct {
	Name    *string `json:"name"`
	Age     *uint   `json:"age"`
	FileIds []uint  `json:"fileIds"`
}

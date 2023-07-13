package model

type CreateArtist struct {
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	FileIds []uint `json:"fileIds"`
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

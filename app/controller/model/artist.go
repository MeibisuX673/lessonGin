package model

type CreateArtist struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ResponseArtist struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateArtist struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

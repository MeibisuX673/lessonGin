package model

type CreateArtist struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ResponseArtist struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

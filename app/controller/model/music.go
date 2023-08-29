package model

// todo Нужно ли подробности о пользователе и альбома
type MusicResponse struct {
	ID       uint         `json:"ID"`
	Name     string       `json:"name"`
	ArtistID uint         `json:"artistID"`
	AlbumID  uint         `json:"albumID"`
	File     FileResponse `json:"musicFile"`
}

type MusicCreate struct {
	Name     string `json:"name" validate:"required"`
	ArtistID uint   `json:"artistID" validate:"required"`
	AlbumID  uint   `json:"albumID" validate:"required"`
	FileID   uint   `json:"musicFileID" validate:"required"`
}

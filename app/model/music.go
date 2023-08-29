package model

type Music struct {
	BaseModel `gorm:"embedded"`
	Name      string `gorm:"varchar(255)"`
	ArtisID   uint
	AlbumID   uint
	File      File
}

func (a Music) TableName() string {
	return "musics"
}

package model

type Music struct {
	BaseModel `gorm:"embedded"`
	Name      string `gorm:"varchar(255)"`
	ArtistID  uint
	AlbumID   uint
	File      File
}

func (a Music) TableName() string {
	return "musics"
}

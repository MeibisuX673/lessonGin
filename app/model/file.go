package model

type File struct {
	BaseModel `gorm:"embedded"`
	Path      string `gorm:"varchar(255)"`
	Name      string `gorm:"varchar(255)"`
	ArtistID  *uint
	AlbumID   *uint
}

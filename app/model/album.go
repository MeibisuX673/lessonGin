package model

type Album struct {
	BaseModel `gorm:"embedded"`
	Title     string `gorm:"varchar(255)" json:"title"`
	ArtistID  uint
	Price     float64
	FileID    *uint   `gorm:"unique"`
	File      *File   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Musics    []Music `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (a Album) TableName() string {
	return "albums"
}

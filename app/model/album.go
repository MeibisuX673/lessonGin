package model

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title    string `gorm:"varchar(255)" json:"title"`
	ArtistID uint
	Price    float64
	FileID   *uint
	File     *File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (a Album) TableName() string {
	return "albums"
}

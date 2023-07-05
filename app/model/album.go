package model

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title    string `gorm:"varchar(255)" json:"title"`
	ArtistID int
	Artist   Artist
	Price    float64
}

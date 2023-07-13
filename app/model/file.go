package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Path     string `gorm:"varchar(255)"`
	Name     string `gorm:"varchar(255)"`
	ArtistID *uint
}

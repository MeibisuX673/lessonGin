package model

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	Name string `gorm:"varchar(255)"`
	Age  int
}

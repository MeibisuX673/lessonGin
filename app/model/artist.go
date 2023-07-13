package model

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	Name   string `gorm:"varchar(255)"`
	Age    uint
	Files  []File  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Albums []Album `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (a Artist) TableName() string {
	return "artists"
}

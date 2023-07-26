package model

type Artist struct {
	BaseModel `gorm:"embedded"`
	Name      string `gorm:"varchar(255)"`
	Age       uint
	Email     string `gorm:"unique"`
	Password  string
	Files     []File  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Albums    []Album `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (a Artist) TableName() string {
	return "artists"
}

package model


type Artist struct {
	ID int	`gorm:"primaryKey"`
	Name string `gorm:"varchar(255)"`
	Age  int
}

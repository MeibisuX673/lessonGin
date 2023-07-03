package database

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var AppDatabase Database

type Database struct {
	BD *gorm.DB
}

func (db *Database) Init() (*Database, error) {

	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	err := db.dbMysqlInit()
	if err != nil {
		return nil, err
	}

	db.migrations()

	return db, nil

}

func (db *Database) dbMysqlInit() error {

	var err error

	db.BD, err = gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) migrations() {

	db.BD.AutoMigrate(&model.Artist{})
	db.BD.AutoMigrate(&model.Album{})

}

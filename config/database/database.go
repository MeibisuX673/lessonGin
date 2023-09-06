package database

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var AppDatabase Database

type Database struct {
	BD *gorm.DB
}

func (db *Database) Init() (*Database, error) {

	if environment.Env.GetEnv("MODE") == "test" {
		err := db.dbSqliteInit()
		if err != nil {
			return nil, err
		}
		db.migrations()
		return db, nil
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

	db.BD, err = gorm.Open(mysql.Open(environment.Env.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil

}

func (db *Database) dbSqliteInit() error {

	var err error

	db.BD, err = gorm.Open(sqlite.Open("./store.db"))
	if err != nil {
		return err
	}

	return nil

}

func (db *Database) migrations() {

	if err := db.BD.AutoMigrate(&model.Artist{}); err != nil {
		panic(err.Error())
	}

	if err := db.BD.AutoMigrate(&model.Album{}); err != nil {
		panic(err.Error())
	}

	if err := db.BD.AutoMigrate(&model.File{}); err != nil {
		panic(err.Error())
	}

	if err := db.BD.AutoMigrate(&model.Music{}); err != nil {
		panic(err.Error())
	}

}

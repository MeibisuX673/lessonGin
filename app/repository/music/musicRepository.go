package music

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository/helper"
	"github.com/MeibisuX673/lessonGin/config/database"
	"gorm.io/gorm/clause"
	"net/http"
)

type MusicRepository struct {
}

func (mr MusicRepository) Create(musicCreate *dto.MusicCreate) (*model.Music, dto.ErrorInterface) {

	file, err := helper.DefinedAssociationFile(musicCreate.FileID)
	if err != nil {
		return nil, err
	}

	db := database.AppDatabase.BD

	var music model.Music = model.Music{
		Name:     musicCreate.Name,
		ArtistID: musicCreate.ArtistID,
		AlbumID:  musicCreate.AlbumID,
		File:     *file,
	}

	result := db.Create(&music)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return &music, nil

}

func (mr MusicRepository) Update(id uint, musicUpdate map[string]interface{}) (*model.Music, dto.ErrorInterface) {

	var music model.Music

	db := database.AppDatabase.BD

	if count := db.Preload(clause.Associations).First(&music, id).RowsAffected; count == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Трек не найден",
		}
	}

	if err := db.Model(&music).Updates(musicUpdate).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &music, nil
}

func (mr MusicRepository) GetAll(query model.Query) ([]model.Music, dto.ErrorInterface) {

	var musics []model.Music

	db := database.AppDatabase.BD

	result := db.Preload(clause.Associations)

	helper.ConfigurationDbQuery(result, query)

	result.Find(&musics)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return musics, nil

}

func (mr MusicRepository) GetById(id uint) (*model.Music, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var music model.Music

	if err := db.Preload(clause.Associations).First(&music, id).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Тррек не найден",
		}
	}

	return &music, nil

}

func (mr MusicRepository) FindOneBy(m map[string]interface{}) (*model.Music, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var music model.Music

	if err := db.Model(&music).Find(m).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &music, nil
}

func (mr MusicRepository) FindBy(m map[string]interface{}) ([]model.Music, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var music []model.Music

	if err := db.Model(&music).Find(m).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return music, nil

}

func (mr MusicRepository) Delete(id uint) dto.ErrorInterface {

	db := database.AppDatabase.BD

	var music model.Music

	if count := db.First(&music, id).RowsAffected; count == 0 {
		return &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Файл не найден",
		}
	}

	if err := db.Unscoped().Delete(&music).Error; err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

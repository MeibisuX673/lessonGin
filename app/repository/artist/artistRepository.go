package artist

import (
	"context"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository/helper"
	"github.com/MeibisuX673/lessonGin/config/database"
	"gorm.io/gorm/clause"
	"net/http"
)

type ArtistRepository struct {
}

func (ar *ArtistRepository) Create(createArtist *dto.ArtistCreate) (*model.Artist, dto.ErrorInterface) {

	db := database.AppDatabase.BD
	artist := model.Artist{
		Name:     createArtist.Name,
		Age:      createArtist.Age,
		Email:    createArtist.Email,
		Password: createArtist.Password,
	}

	// todo definedAssociationFile сделать этот метод для всех репозиториев
	files, err := ar.definedAssociationFile(*createArtist)
	if err != nil {
		return nil, err
	}
	if files != nil {
		artist.Files = files
	}

	if result := db.Create(&artist); result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return &artist, nil

}

func (ar *ArtistRepository) GetAll(query model.Query) ([]model.Artist, dto.ErrorInterface) {

	var artists []model.Artist

	db := database.AppDatabase.BD

	result := db.Preload(clause.Associations)

	helper.ConfigurationDbQuery(result, query)

	result.Find(&artists)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return artists, nil

}

func (ar *ArtistRepository) Update(id uint, artistUpdate map[string]interface{}) (*model.Artist, dto.ErrorInterface) {

	var artist model.Artist

	db := database.AppDatabase.BD

	if count := db.Preload(clause.Associations).First(&artist, id).RowsAffected; count == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		}
	}

	if err := db.Model(&artist).Updates(artistUpdate).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &artist, nil

}

func (ar *ArtistRepository) GetById(id uint) (*model.Artist, dto.ErrorInterface) {

	var artist model.Artist

	db := database.AppDatabase.BD

	err := db.Preload(clause.Associations).First(&artist, id).Error

	if artist.ID == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		}
	}

	if err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &artist, nil

}

func (ar *ArtistRepository) FindOneBy(m map[string]interface{}) (*model.Artist, dto.ErrorInterface) {

	var artist model.Artist

	db := database.AppDatabase.BD

	result := db.Model(&artist).First(m)
	if result.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		}
	}

	return &artist, nil

}

func (ar *ArtistRepository) FindBy(m map[string]interface{}) ([]model.Artist, dto.ErrorInterface) {

	var artists []model.Artist

	db := database.AppDatabase.BD

	result := db.Model(&artists).Find(m)
	if result.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		}
	}

	return artists, nil

}

func (ar *ArtistRepository) Delete(id uint) dto.ErrorInterface {

	var db = database.AppDatabase.BD

	if err := ar.clearAssociations(id); err != nil {
		return err
	}

	err := db.Delete(id).Error

	if err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil

}

func (ar *ArtistRepository) clearAssociations(id uint) dto.ErrorInterface {

	db := database.AppDatabase.BD

	artist, _ := ar.GetById(id)

	//TODO Посмотреть удаление сущьностей файла связанные с альбомом
	if err := db.Unscoped().Model(artist).Association("Albums").Unscoped().Clear(); err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	tx := db.WithContext(context.Background())

	if err := tx.Unscoped().Model(artist).Association("Files").Unscoped().Clear(); err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil

}

func (ar *ArtistRepository) definedAssociationFile(createArtist dto.ArtistCreate) ([]model.File, dto.ErrorInterface) {

	var files []model.File

	db := database.AppDatabase.BD

	if createArtist.FileIds != nil {
		err := db.Find(&files, createArtist.FileIds).Error
		if err != nil {
			return nil, &dto.Error{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}

	}

	for _, file := range files {
		if file.AlbumID != nil || file.ArtistID != nil {
			return nil, nil
		}
	}

	return files, nil
}

package album

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository/helper"
	"github.com/MeibisuX673/lessonGin/config/database"
	"gorm.io/gorm/clause"
	"net/http"
)

type AlbumRepository struct {
}

func (ar *AlbumRepository) Create(albumCreate dto.AlbumCreate) (*model.Album, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	file, err := helper.DefinedAssociationFile(*albumCreate.FileId)
	if err != nil {
		return nil, err
	}
	if file == nil {
		return nil, &dto.Error{
			Status:  http.StatusConflict,
			Message: "File in use",
		}
	}

	album := model.Album{
		Title:    albumCreate.Title,
		ArtistID: albumCreate.ArtistID,
		Price:    albumCreate.Price,
		File:     file,
	}

	result := db.Create(&album)
	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return &album, nil

}

func (ar *AlbumRepository) Update(id uint, albumUpdate map[string]interface{}) (*model.Album, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var album model.Album

	if count := db.First(&album, id).RowsAffected; count == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Альбои не найден",
		}
	}

	result := db.Model(&album).Updates(albumUpdate)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return &album, nil

}

func (ar *AlbumRepository) GetAll(query model.Query) ([]model.Album, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var albums []model.Album

	result := db.Preload(clause.Associations)

	helper.ConfigurationDbQuery(result, query)

	result.Find(&albums)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return albums, nil

}

func (ar *AlbumRepository) GetById(id uint) (*model.Album, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var album model.Album

	result := db.Preload(clause.Associations).First(&album, id)

	if result.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Альбом не найден",
		}
	}

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return &album, nil

}

func (ar *AlbumRepository) FindBy(m map[string]interface{}) ([]model.Album, dto.ErrorInterface) {

	var albums []model.Album

	db := database.AppDatabase.BD

	result := db.Model(&albums).Find(m)
	if result.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Альбом не найден",
		}
	}

	return albums, nil

}

func (ar *AlbumRepository) FindOneBy(arguments map[string]interface{}) (*model.Album, dto.ErrorInterface) {

	albums, err := ar.FindOneBy(arguments)
	if err != nil {
		return nil, err
	}

	return albums, nil

}

func (ar *AlbumRepository) Delete(id uint) dto.ErrorInterface {

	db := database.AppDatabase.BD

	var album model.Album

	if count := db.First(&album, id).RowsAffected; count == 0 {
		return &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Album not found",
		}
	}

	result := db.Unscoped().Delete(&model.Album{}, id)

	if result.Error != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return nil

}

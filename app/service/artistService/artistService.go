package artistService

import (
	"context"
	"encoding/json"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/MeibisuX673/lessonGin/app/service/hashPasswordService"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
	"github.com/MeibisuX673/lessonGin/config/database"
	"gorm.io/gorm/clause"
	"net/http"
)

func CreateArtist(artistRequest *dto.CreateArtist) (*dto.ResponseArtist, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	artist := model.Artist{
		Name:  artistRequest.Name,
		Age:   artistRequest.Age,
		Email: artistRequest.Email,
	}

	hashedPassword, err := hashPasswordService.HashPassword(artistRequest.Password)
	if err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	artist.Password = hashedPassword

	if artistRequest.FileIds != nil {
		for _, fileId := range artistRequest.FileIds {

			_, err := fileService.GetFileById(fileId)

			if err != nil {
				return nil, err
			}

		}
	}

	if result := db.Create(&artist); result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	response := ConvertToOneArtistResponse(artist)

	return &response, nil

}

func GetCollectionArtist(query model.Query) ([]dto.ResponseArtist, dto.ErrorInterface) {

	var artists []model.Artist

	db := database.AppDatabase.BD

	result := db.Preload(clause.Associations)

	queryService.ConfigurationDbQuery(result, query)

	result.Find(&artists)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	response := convertToArtistResponseCollection(artists)

	return response, nil

}

func GetArtistById(id uint) (*dto.ResponseArtist, dto.ErrorInterface) {

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

	response := ConvertToOneArtistResponse(artist)

	return &response, nil

}

func UpdateArtist(id int, updateArtist dto.UpdateArtist) (*dto.ResponseArtist, dto.ErrorInterface) {

	var artistUpdateMap map[string]interface{}

	updateArtistByte, _ := json.Marshal(updateArtist)

	if err := json.Unmarshal(updateArtistByte, &artistUpdateMap); err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	sortMap := checkNil(artistUpdateMap)

	var artist model.Artist

	db := database.AppDatabase.BD

	if count := db.First(&artist, id).RowsAffected; count == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		}
	}

	if err := db.Model(&artist).Updates(sortMap).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	response := ConvertToOneArtistResponse(artist)

	return &response, nil

}

func DeleteArtist(id uint) dto.ErrorInterface {

	db := database.AppDatabase.BD

	var artist model.Artist

	if count := db.Preload(clause.Associations).First(&artist, id).RowsAffected; count == 0 {
		return &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		}
	}

	if err := deleteFiles(artist); err != nil {
		return err
	}

	if err := clearAssociations(&artist); err != nil {
		return err
	}

	tx := db.WithContext(context.Background())

	err := tx.Delete(&artist).Error

	if err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil

}

func checkNil(args map[string]interface{}) map[string]interface{} {

	sortNil := make(map[string]interface{})

	for key, value := range args {
		if value != nil {
			sortNil[key] = value
		}
	}

	return sortNil

}

func clearAssociations(artist *model.Artist) dto.ErrorInterface {

	db := database.AppDatabase.BD

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

func deleteFiles(artist model.Artist) dto.ErrorInterface {

	var artistFiles []model.File
	var albumFiles []model.File

	db := database.AppDatabase.BD

	if artist.Files != nil {
		db := database.AppDatabase.BD
		db.Where("artist_id = ?", artist.ID).Find(&artistFiles)
		if err := fileService.DeleteFileFromDisk(artist.Files); err != nil {
			return err
		}
	}

	if artist.Albums != nil {
		for _, album := range artist.Albums {
			if album.FileID != nil {
				var filesModel []model.File
				db.Find(&filesModel, album.FileID)
				albumFiles = append(albumFiles, filesModel...)
			}

		}
		if err := fileService.DeleteFileFromDisk(albumFiles); err != nil {
			return err
		}
	}

	return nil

}

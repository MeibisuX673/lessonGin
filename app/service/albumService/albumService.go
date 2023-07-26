package albumService

import (
	"encoding/json"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
	"github.com/MeibisuX673/lessonGin/config/database"
	"gorm.io/gorm/clause"
	"net/http"
)

func CreateAlbum(albumRequest dto.CreateAlbum) (*dto.ResponseAlbum, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	album := model.Album{
		Title:    albumRequest.Title,
		ArtistID: albumRequest.ArtistID,
		Price:    albumRequest.Price,
	}

	if albumRequest.FileId != nil {

		file, err := fileService.GetFileById(*albumRequest.FileId)
		if err != nil {
			return nil, err
		}

		fileId := file.ID
		album.FileID = &fileId

	}

	result := db.Create(&album)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	response := convertToOneAlbumResponse(album)

	return &response, nil

}

func GetCollectionAlbum(query model.Query) ([]dto.ResponseAlbum, *dto.Error) {

	db := database.AppDatabase.BD

	var albums []model.Album

	result := db.Preload(clause.Associations)

	queryService.ConfigurationDbQuery(result, query)

	result.Find(&albums)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	response := convertToAlbumResponseCollection(albums)

	return response, nil

}

func GetAlbumById(id int) (*dto.ResponseAlbum, dto.ErrorInterface) {

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

	response := convertToOneAlbumResponse(album)

	return &response, nil

}

func UpdateAlbum(id int, albumUpdate dto.UpdateAlbum) (*dto.ResponseAlbum, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var albumUpdateMap map[string]interface{}

	updateAlbumByte, _ := json.Marshal(albumUpdate)

	if err := json.Unmarshal(updateAlbumByte, &albumUpdateMap); err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	sortMap := checkNil(albumUpdateMap)

	var album model.Album

	tx := db.First(&album, id)

	if tx.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "альбом не найден",
		}
	}

	if albumUpdate.ArtistID != nil {
		if _, err := artistService.GetArtistById(*albumUpdate.ArtistID); err != nil {
			return nil, err
		}
	}

	result := db.Model(&album).Updates(sortMap)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	response := convertToOneAlbumResponse(album)

	return &response, nil

}

func DeleteAlbum(id int) dto.ErrorInterface {

	db := database.AppDatabase.BD

	var album model.Album

	if count := db.First(&album, id).RowsAffected; count == 0 {
		return &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Album not found",
		}
	}

	if album.File != nil {

		if err := fileService.DeleteFileFromDisk([]model.File{*album.File}); err != nil {
			return err
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

func checkNil(args map[string]interface{}) map[string]interface{} {

	sortNil := make(map[string]interface{})

	for key, value := range args {
		if value != nil {
			sortNil[key] = value
		}
	}

	return sortNil

}

package albumService

import (
	"encoding/json"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/config/database"
)

func CreateAlbum(albumRequest dto.CreateAlbum) (*model.Album, error) {

	db := database.AppDatabase.BD

	album := model.Album{
		Title:    albumRequest.Title,
		ArtistID: albumRequest.ArtistID,
		Price:    albumRequest.Price,
	}

	result := db.Create(&album)

	if result.Error != nil {
		return nil, result.Error
	}

	return &album, nil

}

func GetCollectionArtist() ([]model.Album, error) {

	db := database.AppDatabase.BD

	var albums []model.Album

	result := db.Find(&albums)

	if result.Error != nil {
		return nil, result.Error
	}

	return albums, nil

}

func GetAlbumById(id int) (*model.Album, error) {

	db := database.AppDatabase.BD

	var album model.Album

	result := db.Find(&album, id)

	if result.RowsAffected == 0 {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &album, nil

}

func UpdateAlbum(id int, albumUpdate dto.UpdateAlbum) (*model.Album, error) {

	db := database.AppDatabase.BD

	var albumUpdateMap map[string]interface{}

	updateAlbumByte, _ := json.Marshal(albumUpdate)

	if err := json.Unmarshal(updateAlbumByte, &albumUpdateMap); err != nil {
		return nil, err
	}

	sortMap := checkNil(albumUpdateMap)

	var album model.Album

	db.First(&album, id)
	result := db.Model(&album).Updates(sortMap)

	if result.Error != nil {
		return nil, result.Error
	}

	return &album, nil

}

func DeleteAlbum(id int) error {

	db := database.AppDatabase.BD

	result := db.Delete(&model.Album{}, id)

	if result.Error != nil {
		return result.Error
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

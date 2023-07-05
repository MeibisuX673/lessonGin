package artistService

import (
	"encoding/json"

	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/config/database"
)

func CreateArtist(artistRequest *dto.CreateArtist) (*model.Artist, error) {

	db := database.AppDatabase.BD

	artist := model.Artist{
		Name: artistRequest.Name,
		Age:  artistRequest.Age,
	}

	if result := db.Create(&artist); result.Error != nil {
		return nil, result.Error
	}

	return &artist, nil

}

func GetCollectionArtist() ([]model.Artist, error) {

	var artists []model.Artist

	db := database.AppDatabase.BD

	result := db.Find(&artists)

	if result.Error != nil {
		return nil, result.Error
	}

	return artists, nil

}

func GetArtistById(id int) (*model.Artist, error) {

	var artist model.Artist

	db := database.AppDatabase.BD

	result := db.First(&artist, id)

	if result.RowsAffected == 0 {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &artist, nil

}

func UpdateArtist(id int, updateArtist dto.UpdateArtist) (*model.Artist, error) {

	var artistUpdateMap map[string]interface{}

	updateArtistByte, _ := json.Marshal(updateArtist)

	if err := json.Unmarshal(updateArtistByte, &artistUpdateMap); err != nil {
		return nil, err
	}

	sortMap := checkNil(artistUpdateMap)

	var artist model.Artist

	db := database.AppDatabase.BD

	db.First(&artist, id)
	db.Model(&artist).Updates(sortMap)

	return &artist, nil

}

func DeleteArtist(id int) error {

	db := database.AppDatabase.BD

	result := db.Delete(&model.Artist{}, id)

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

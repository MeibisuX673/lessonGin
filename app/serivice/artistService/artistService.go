package artistService

import (
	model2 "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/config/database"
)

func CreateArtist(artistRequest *model2.CreateArtist) (*model.Artist, error) {

	db := database.AppDatabase.BD

	artist := model.Artist{
		Name: artistRequest.Name,
		Age:  uint16(artistRequest.Age),
	}

	if result := db.Create(&artist); result.Error != nil {
		return nil, result.Error
	}

	return &artist, nil

}

func GetCollectionArtist() ([]*model.Artist, error) {

	var artists []*model.Artist

	db := database.AppDatabase.BD

	if result := db.Find(artists); result.Error != nil {
		return nil, result.Error
	}

	return artists, nil

}

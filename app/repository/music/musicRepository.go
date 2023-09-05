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

func (m MusicRepository) Create(musicCreate *dto.MusicCreate) (*model.Music, dto.ErrorInterface) {

	file, err := helper.DefinedAssociationFile(musicCreate.FileID)
	if err != nil {
		return nil, err
	}

	db := database.AppDatabase.BD

	var music model.Music = model.Music{
		Name:    musicCreate.Name,
		ArtisID: musicCreate.ArtistID,
		AlbumID: musicCreate.AlbumID,
		File:    *file,
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

func (m MusicRepository) Update(id uint, artistUpdate map[string]interface{}) (*model.Music, dto.ErrorInterface) {
	//TODO implement me
	panic("implement me")
}

func (m MusicRepository) GetAll(query model.Query) ([]model.Music, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var musics []model.Music

	if err := db.Preload(clause.Associations).Find(&musics).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return musics, nil

}

func (m MusicRepository) GetById(id uint) (*model.Music, dto.ErrorInterface) {
	//TODO implement me
	panic("implement me")
}

func (m MusicRepository) FindOneBy(m2 map[string]interface{}) (*model.Music, dto.ErrorInterface) {
	//TODO implement me
	panic("implement me")
}

func (m MusicRepository) FindBy(m2 map[string]interface{}) ([]model.Music, dto.ErrorInterface) {
	//TODO implement me
	panic("implement me")
}

func (m MusicRepository) Delete(id uint) dto.ErrorInterface {
	//TODO implement me
	panic("implement me")
}

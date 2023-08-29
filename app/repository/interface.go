package repository

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

type ArtistRepositoryInterface interface {
	Create(artist *dto.ArtistCreate) (*model.Artist, dto.ErrorInterface)
	Update(id uint, artistUpdate map[string]interface{}) (*model.Artist, dto.ErrorInterface)
	GetAll(query model.Query) ([]model.Artist, dto.ErrorInterface)
	GetById(id uint) (*model.Artist, dto.ErrorInterface)
	FindOneBy(map[string]interface{}) (*model.Artist, dto.ErrorInterface)
	FindBy(map[string]interface{}) ([]model.Artist, dto.ErrorInterface)
	Delete(id uint) dto.ErrorInterface
}

type AlbumRepositoryInterface interface {
	Create(album dto.AlbumCreate) (*model.Album, dto.ErrorInterface)
	Update(id uint, albumUpdate map[string]interface{}) (*model.Album, dto.ErrorInterface)
	GetAll(query model.Query) ([]model.Album, dto.ErrorInterface)
	GetById(id uint) (*model.Album, dto.ErrorInterface)
	FindBy(map[string]interface{}) ([]model.Album, dto.ErrorInterface)
	FindOneBy(map[string]interface{}) (*model.Album, dto.ErrorInterface)
	Delete(id uint) dto.ErrorInterface
}

type FileRepositoryInterface interface {
	Create(file dto.CreateFile) (*model.File, dto.ErrorInterface)
	GetAll(query model.Query) ([]model.File, dto.ErrorInterface)
	GetById(id uint) (*model.File, dto.ErrorInterface)
	FindOneBy(map[string]interface{}) (*model.File, dto.ErrorInterface)
	FindBy(map[string]interface{}) ([]model.File, dto.ErrorInterface)
	Delete(id uint) dto.ErrorInterface
}

type MusicRepositoryInterface interface {
	Create(music *dto.MusicCreate) (*model.Music, dto.ErrorInterface)
	Update(id uint, artistUpdate map[string]interface{}) (*model.Music, dto.ErrorInterface)
	GetAll(query model.Query) ([]model.Music, dto.ErrorInterface)
	GetById(id uint) (*model.Music, dto.ErrorInterface)
	FindOneBy(map[string]interface{}) (*model.Music, dto.ErrorInterface)
	FindBy(map[string]interface{}) ([]model.Music, dto.ErrorInterface)
	Delete(id uint) dto.ErrorInterface
}

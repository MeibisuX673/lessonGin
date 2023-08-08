package albumService

import (
	"encoding/json"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository"
	"net/http"
)

type AlbumService struct {
	AlbumRepository repository.AlbumRepositoryInterface
}

func New(albumRepository repository.AlbumRepositoryInterface) *AlbumService {

	return &AlbumService{AlbumRepository: albumRepository}
}

func (as *AlbumService) CreateAlbum(albumRequest dto.AlbumCreate) (*dto.AlbumResponse, dto.ErrorInterface) {

	album, err := as.AlbumRepository.Create(albumRequest)
	if err != nil {
		return nil, err
	}

	response := convertToOneAlbumResponse(*album)

	return &response, nil

}

func (as *AlbumService) GetCollectionAlbum(query model.Query) ([]dto.AlbumResponse, dto.ErrorInterface) {

	albums, err := as.AlbumRepository.GetAll(query)
	if err != nil {
		return nil, err
	}

	response := convertToAlbumResponseCollection(albums)

	return response, nil

}

func (as *AlbumService) GetAlbumById(id uint) (*dto.AlbumResponse, dto.ErrorInterface) {

	album, err := as.AlbumRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := convertToOneAlbumResponse(*album)

	return &response, nil

}

func (as *AlbumService) UpdateAlbum(id uint, albumUpdate dto.UpdateAlbum) (*dto.AlbumResponse, dto.ErrorInterface) {

	var albumUpdateMap map[string]interface{}

	updateAlbumByte, _ := json.Marshal(albumUpdate)

	if err := json.Unmarshal(updateAlbumByte, &albumUpdateMap); err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	sortMap := checkNil(albumUpdateMap)

	album, errUpdate := as.AlbumRepository.Update(id, sortMap)
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := convertToOneAlbumResponse(*album)

	return &response, nil

}

func (as *AlbumService) GetBy(arguments map[string]interface{}) ([]dto.AlbumResponse, dto.ErrorInterface) {

	albums, err := as.AlbumRepository.FindBy(arguments)
	if err != nil {
		return nil, err
	}

	response := convertToAlbumResponseCollection(albums)

	return response, nil

}

func (as *AlbumService) DeleteAlbum(id uint) dto.ErrorInterface {

	_, err := as.AlbumRepository.GetById(id)
	if err != nil {
		return err
	}

	if err := as.AlbumRepository.Delete(id); err != nil {
		return err
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

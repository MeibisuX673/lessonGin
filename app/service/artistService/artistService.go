package artistService

import (
	"encoding/json"
	"github.com/MeibisuX673/lessonGin/app/Helper"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository"
	"github.com/MeibisuX673/lessonGin/app/service/helper"
	"net/http"
)

type ArtistService struct {
	ArtistRepository repository.ArtistRepositoryInterface
}

func New(
	artistRepository repository.ArtistRepositoryInterface,
) *ArtistService {

	return &ArtistService{
		ArtistRepository: artistRepository,
	}
}

func (as *ArtistService) CreateArtist(artistRequest *dto.ArtistCreate) (*dto.ArtistResponse, dto.ErrorInterface) {

	hashedPassword, err := Helper.HashPassword(artistRequest.Password)
	if err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	artistRequest.Password = hashedPassword

	artist, errCreate := as.ArtistRepository.Create(artistRequest)
	if errCreate != nil {
		return nil, errCreate
	}

	response := ConvertToOneArtistResponse(*artist)

	return &response, nil

}

func (as *ArtistService) GetCollectionArtist(query model.Query) ([]dto.ArtistResponse, dto.ErrorInterface) {

	artists, err := as.ArtistRepository.GetAll(query)
	if err != nil {
		return nil, err
	}

	response := convertToArtistResponseCollection(artists)

	return response, nil

}

func (as *ArtistService) GetArtistById(id uint) (*dto.ArtistResponse, dto.ErrorInterface) {

	artist, err := as.ArtistRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := ConvertToOneArtistResponse(*artist)

	return &response, nil

}

func (as *ArtistService) GetOneBy(arguments map[string]interface{}) (*dto.ArtistResponse, dto.ErrorInterface) {

	artist, err := as.ArtistRepository.FindOneBy(arguments)
	if err != nil {
		return nil, err
	}

	response := ConvertToOneArtistResponse(*artist)

	return &response, nil

}

func (as *ArtistService) UpdateArtist(id uint, updateArtist dto.UpdateArtist) (*dto.ArtistResponse, dto.ErrorInterface) {

	var artistUpdateMap map[string]interface{}

	updateArtistByte, _ := json.Marshal(updateArtist)

	if err := json.Unmarshal(updateArtistByte, &artistUpdateMap); err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	sortMap := helper.CheckNil(artistUpdateMap)

	artist, err := as.ArtistRepository.Update(id, sortMap)
	if err != nil {
		return nil, err
	}

	response := ConvertToOneArtistResponse(*artist)

	return &response, nil

}

func (as *ArtistService) DeleteArtist(id uint) dto.ErrorInterface {

	_, err := as.ArtistRepository.GetById(id)
	if err != nil {
		return err
	}

	err = as.ArtistRepository.Delete(id)
	if err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil

}

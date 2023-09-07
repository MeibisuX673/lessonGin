package audioService

import (
	"bufio"
	"encoding/json"
	"fmt"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository"
	"github.com/MeibisuX673/lessonGin/app/service/helper"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"github.com/Vernacular-ai/godub"
	audio "github.com/Vernacular-ai/godub/converter"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
)

type AudioService struct {
	MusicRepository repository.MusicRepositoryInterface
}

func New(musicRepository repository.MusicRepositoryInterface) *AudioService {
	return &AudioService{MusicRepository: musicRepository}
}

func Mp3ConvertToWav(name string) string {

	filePath := path.Join(environment.Env.GetEnv("DIR_MUSIC"), name)

	regex := regexp.MustCompile("\\.[a-z0-9]+")
	newName := regex.ReplaceAllString(name, ".wav")

	toFilePath := path.Join(environment.Env.GetEnv("DIR_MUSIC"), newName)

	w, err := os.Create(toFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = audio.NewConverter(w).WithBitRate(audio.M4ABitRateEconomy).WithDstFormat("wav").Convert(filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = godub.NewLoader().Load(toFilePath)
	if err != nil {
		log.Fatal(err)
	}

	return newName

}

func (as *AudioService) GetBytes(fileName string) []byte {

	file, _ := os.Open(path.Join(environment.Env.GetEnv("DIR_MUSIC"), fileName))
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)

	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)

	}

	return bs

}

func (as *AudioService) CreateMusic(musicCreate dto.MusicCreate) (*dto.MusicResponse, dto.ErrorInterface) {

	music, err := as.MusicRepository.Create(&musicCreate)
	if err != nil {
		return nil, err
	}

	//todo converter
	result := &dto.MusicResponse{
		Name:     music.Name,
		ArtistID: music.ArtistID,
		AlbumID:  music.AlbumID,
		File: dto.FileResponse{
			ID:   music.File.ID,
			Name: music.File.Name,
			Path: music.File.Path,
		},
	}

	return result, nil

}

func (as *AudioService) GetCollectionMusic(query model.Query) ([]dto.MusicResponse, dto.ErrorInterface) {

	musics, err := as.MusicRepository.GetAll(query)
	if err != nil {
		return nil, err
	}

	response := convertToMusicResponseCollection(musics)

	return response, nil

}

func (as *AudioService) GetMusicById(id uint) (*dto.MusicResponse, dto.ErrorInterface) {

	music, err := as.MusicRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	//todo converter
	response := dto.MusicResponse{
		ID:       music.ID,
		Name:     music.Name,
		ArtistID: music.ArtistID,
		AlbumID:  music.AlbumID,
		File: dto.FileResponse{
			ID:   music.File.ID,
			Name: music.File.Name,
			Path: music.File.Path,
		},
	}

	return &response, nil

}

func (as *AudioService) UpdateMusic(id uint, updateMusic dto.MusicUpdate) (*dto.MusicResponse, dto.ErrorInterface) {

	var musicUpdateMap map[string]interface{}

	updateArtistByte, _ := json.Marshal(updateMusic)

	if err := json.Unmarshal(updateArtistByte, &musicUpdateMap); err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	sortMap := helper.CheckNil(musicUpdateMap)

	music, err := as.MusicRepository.Update(id, sortMap)
	if err != nil {
		return nil, err
	}

	response := dto.MusicResponse{
		Name:     music.Name,
		ArtistID: music.ArtistID,
		AlbumID:  music.AlbumID,
		File: dto.FileResponse{
			ID:   music.File.ID,
			Name: music.File.Name,
			Path: music.File.Path,
		},
	}

	return &response, nil

}

func (as *AudioService) DeleteMusic(id uint) dto.ErrorInterface {

	err := as.MusicRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil

}

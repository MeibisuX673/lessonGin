package audioService

import (
	"bufio"
	"fmt"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"github.com/Vernacular-ai/godub"
	audio "github.com/Vernacular-ai/godub/converter"
	"io"
	"log"
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
		ArtistID: music.ArtisID,
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

//func (as *AudioService) GetMusicById(id uint) (*dto.MusicResponse, dto.ErrorInterface) {
//
//	db := database.AppDatabase.BD
//
//	var musics model.File
//
//	result := db.Preload(clause.Associations).First(&musics, id)
//	if result.RowsAffected == 0 {
//		return nil, &dto.Error{
//			Status:  http.StatusNotFound,
//			Message: "Файл не найден",
//		}
//	}
//
//	bytes := as.getBytes(musics.Name)
//	//todo converter
//	response := dto.MusicResponse{
//		Name:        musics.Name,
//		Bytes:       bytes,
//		ArtistID:    musics.ArtistID,
//		AlbumID:     musics.AlbumID,
//		MusicFileID: musics.ID,
//	}
//
//	return &response, nil
//
//}

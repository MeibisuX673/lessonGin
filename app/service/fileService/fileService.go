package fileService

import (
	"fmt"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository"
	"github.com/MeibisuX673/lessonGin/app/service/audioService"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
)

type FileService struct {
	FileRepository repository.FileRepositoryInterface
}

func New(fileRepository repository.FileRepositoryInterface) *FileService {
	return &FileService{FileRepository: fileRepository}
}

func init() {
	InitSaveByExt()
}

var converts = map[string]func(string) string{
	".mp3": audioService.Mp3ConvertToWav,
}

var extensionsImage []string = []string{
	".png",
	".jpg",
	".svg",
}

var extensionsMusic []string = []string{
	".mp3",
	".wav",
}

func GetMusicExtension() []string {
	return extensionsMusic
}

var saveByExt = map[string][2]string{}

func InitSaveByExt() {
	for _, s := range extensionsImage {
		saveByExt[s] = [2]string{
			environment.Env.GetEnv("DIR_IMAGES"), environment.Env.GetEnv("IMAGE_URL"),
		}
	}
	for _, s := range extensionsMusic {
		saveByExt[s] = [2]string{
			environment.Env.GetEnv("DIR_MUSIC"), environment.Env.GetEnv("MUSIC_URL"),
		}
	}
}

func (fs *FileService) UploadFile(c *gin.Context) (*model.File, dto.ErrorInterface) {

	file, err := c.FormFile("file")

	if err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}
	}

	extension := filepath.Ext(file.Filename)

	key, ok := saveByExt[extension]

	if !ok {
		return nil, &dto.Error{
			Status:  http.StatusBadRequest,
			Message: "invalid format",
		}
	}

	newFileName := uuid.New().String() + extension

	dir := key[0] + "/" + newFileName
	if err := c.SaveUploadedFile(file, dir); err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: "Failed to save file",
		}
	}

	if fun, ok := converts[extension]; ok {
		newFileName = fun(newFileName)
	}

	newFile := model.File{
		Name: newFileName,
		Path: fmt.Sprintf(key[1], newFileName),
	}

	return &newFile, nil

}

func (fs *FileService) CreateFileInDatabase(createFile dto.CreateFile) (*dto.FileResponse, dto.ErrorInterface) {

	file, err := fs.FileRepository.Create(createFile)
	if err != nil {
		return nil, err
	}

	response := convertToOneFileResponse(*file)

	return &response, nil

}

func (fs *FileService) GetFileById(id uint) (*dto.FileResponse, dto.ErrorInterface) {

	file, err := fs.FileRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := convertToOneFileResponse(*file)

	return &response, nil

}

func (fs *FileService) GetFileCollection(query model.Query) ([]dto.FileResponse, dto.ErrorInterface) {

	files, err := fs.FileRepository.GetAll(query)
	if err != nil {
		return nil, err
	}

	response := convertToFileResponseCollection(files)

	return response, nil

}

func (fs *FileService) DeleteFile(id uint) dto.ErrorInterface {

	file, err := fs.FileRepository.GetById(id)
	if err != nil {
		return err
	}

	if err := os.Remove(fmt.Sprintf(os.Getenv("DIR_IMAGES")+"/%s", file.Name)); err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := fs.FileRepository.Delete(id); err != nil {
		return err
	}

	return nil

}

func (fs *FileService) GetBy(arguments map[string]interface{}) ([]dto.FileResponse, dto.ErrorInterface) {

	files, err := fs.FileRepository.FindBy(arguments)
	if err != nil {
		return nil, err
	}

	response := convertToFileResponseCollection(files)

	return response, nil

}

func (fs *FileService) DeleteFileFromDisk(files []dto.FileResponse) dto.ErrorInterface {

	for _, file := range files {

		if err := os.Remove(fmt.Sprintf(os.Getenv("DIR_IMAGES")+"/%s", file.Name)); err != nil {
			return &dto.Error{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	return nil

}

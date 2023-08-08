package fileService

import (
	"errors"
	"fmt"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository"
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

var extensions []string = []string{
	".png",
	".jpg",
	".svg",
}

func (fs *FileService) UploadFile(c *gin.Context) (*model.File, error) {

	file, err := c.FormFile("file")

	if err != nil {
		return nil, err
	}

	extension := filepath.Ext(file.Filename)

	if !checkExtension(extension) {
		return nil, errors.New("Invalid file extension")
	}

	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "./assets/images/"+newFileName); err != nil {
		return nil, err
	}

	newFile := model.File{
		Name: newFileName,
		Path: fmt.Sprintf(os.Getenv("IMAGE_URL"), newFileName),
	}

	if err != nil {
		return nil, err
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

func checkExtension(extension string) bool {

	for _, value := range extensions {
		if value == extension {
			return true
		}
	}

	return false

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

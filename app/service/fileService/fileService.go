package fileService

import (
	"errors"
	"fmt"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"net/http"
	"os"
	"path/filepath"
)

var extensions []string = []string{
	".png",
	".jpg",
	".svg",
}

func UploadFile(c *gin.Context) (*model.File, error) {

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

func CreateFileInDatabase(createFile dto.CreateFile) (*dto.ResponseFile, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var file model.File = model.File{
		Name: createFile.Name,
		Path: createFile.Path,
	}

	result := db.Create(&file)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	response := convertToOneFileResponse(file)

	return &response, nil

}

func GetFileById(id uint) (*dto.ResponseFile, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var file model.File

	result := db.Preload(clause.Associations).First(&file, id)

	if result.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Файл не найден",
		}
	}

	response := convertToOneFileResponse(file)

	return &response, nil

}

func GetFileCollection() ([]dto.ResponseFile, dto.ErrorInterface) {

	var files []model.File

	db := database.AppDatabase.BD

	if err := db.Model(model.File{}).Find(&files).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	response := convertToFileResponseCollection(files)

	return response, nil

}

func DeleteFile(id uint) dto.ErrorInterface {

	db := database.AppDatabase.BD

	var file model.File

	if count := db.Model(model.File{}).First(&file, id).RowsAffected; count == 0 {
		return &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Файл не найден",
		}
	}

	if err := os.Remove(fmt.Sprintf(os.Getenv("DIR_IMAGES")+"/%s", file.Name)); err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := db.Model(model.File{}).Unscoped().Delete(&file).Error; err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil

}

func checkExtension(extension string) bool {

	for _, value := range extensions {
		if value == extension {
			return true
		}
	}

	return false

}

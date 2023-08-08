package file

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository/mySql/helper"
	"github.com/MeibisuX673/lessonGin/config/database"
	"gorm.io/gorm/clause"
	"net/http"
)

type FileRepository struct {
}

func (f FileRepository) FindOneBy(m map[string]interface{}) (*model.File, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var file model.File

	if err := db.Model(&file).Find(m).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &file, nil
}

func (f FileRepository) Create(fileCreate dto.CreateFile) (*model.File, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var file model.File = model.File{
		Name: fileCreate.Name,
		Path: fileCreate.Path,
	}

	result := db.Create(&file)

	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return &file, nil
}

func (f FileRepository) GetAll(query model.Query) ([]model.File, dto.ErrorInterface) {

	var files []model.File

	db := database.AppDatabase.BD

	result := db.Preload(clause.Associations)

	helper.ConfigurationDbQuery(result, query)

	result.Find(&files)
	if result.Error != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: result.Error.Error(),
		}
	}

	return files, nil
}

func (f FileRepository) GetById(id uint) (*model.File, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var file model.File

	result := db.Preload(clause.Associations).First(&file, id)
	if result.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Файл не найден",
		}
	}

	return &file, nil

}

func (f FileRepository) FindBy(m map[string]interface{}) ([]model.File, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var files []model.File

	if err := db.Model(&files).Find(m).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return files, nil

}

func (f FileRepository) Delete(id uint) dto.ErrorInterface {

	db := database.AppDatabase.BD

	var file model.File

	if count := db.First(&file, id).RowsAffected; count == 0 {
		return &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Файл не найден",
		}
	}

	if err := db.Unscoped().Delete(&file).Error; err != nil {
		return &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil

}

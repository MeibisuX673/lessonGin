package helper

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/config/database"
	"gorm.io/gorm"
	"net/http"
)

func ConfigurationDbQuery(db *gorm.DB, query model.Query) {

	if query.Filters != nil {
		for _, value := range query.Filters {
			db.Where(value)
		}
	}
	if query.RangeFilters != nil {
		for _, value := range query.RangeFilters {
			db.Where(value)
		}
	}

	if query.Orders != nil {
		for _, value := range query.Orders {
			db.Order(value)
		}
	}

	db.Offset(int(query.Page*query.Limit - query.Limit))

	db.Limit(int(query.Limit))

}

func DefinedAssociationFile(id uint) (*model.File, dto.ErrorInterface) {

	db := database.AppDatabase.BD

	var file model.File

	if err := db.First(&file, id).Error; err != nil {
		return nil, &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if file.AlbumID != nil || file.ArtistID != nil || file.MusicID != nil {
		return nil, nil
	}

	return &file, nil
}

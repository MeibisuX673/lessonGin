package authService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/service/hashPasswordService"
	"github.com/MeibisuX673/lessonGin/config/database"
	"net/http"
)

func CheckUser(auth dto.Auth) (*model.Artist, dto.ErrorInterface) {

	var artist model.Artist

	db := database.AppDatabase.BD

	result := db.First(&artist, "email = ?", auth.Email)

	if result.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		}
	}

	if hashPasswordService.CheckHashPassword(auth.Password, artist.Password) {
		return nil, &dto.Error{
			Status:  http.StatusBadRequest,
			Message: "Не правильный пароль",
		}
	}

	return &artist, nil
}

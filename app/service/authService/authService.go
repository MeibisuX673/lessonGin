package authService

import (
	"github.com/MeibisuX673/lessonGin/app/Helper"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/repository"
	"github.com/MeibisuX673/lessonGin/config/database"
	"net/http"
)

type AuthService struct {
	ArtistRepository repository.ArtistRepositoryInterface
}

func New(artistRepository repository.ArtistRepositoryInterface) *AuthService {
	return &AuthService{ArtistRepository: artistRepository}
}

func (authS *AuthService) CheckUser(auth dto.Auth) (*model.Artist, dto.ErrorInterface) {

	var artist model.Artist

	db := database.AppDatabase.BD

	result := db.First(&artist, "email = ?", auth.Email)

	if result.RowsAffected == 0 {
		return nil, &dto.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		}
	}

	if Helper.CheckHashPassword(auth.Password, artist.Password) {
		return nil, &dto.Error{
			Status:  http.StatusBadRequest,
			Message: "Неправильный пароль",
		}
	}

	return &artist, nil
}

package jwtService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

type JWTService struct {
}

func (jwtS *JWTService) CreateJwtToken(artist model.Artist) (string, dto.ErrorInterface) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["sub"] = artist.ID

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return tokenStr, nil

}

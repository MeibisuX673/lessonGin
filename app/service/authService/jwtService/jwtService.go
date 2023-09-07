package jwtService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type JWTService struct {
}

func (jwtS *JWTService) CreateJwtToken(artist model.Artist) (string, dto.ErrorInterface) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["sub"] = artist.ID

	tokenStr, err := token.SignedString([]byte(environment.Env.GetEnv("SECRET")))
	if err != nil {
		return "", &dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return tokenStr, nil

}

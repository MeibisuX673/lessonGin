package securityService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"strings"
)

func GetCurrentUser(c *gin.Context) dto.ArtistResponse {

	tokenString := c.GetHeader("Authorization")

	tokenData := strings.Fields(tokenString)
	token, _ := jwt.Parse(tokenData[1], func(token *jwt.Token) (interface{}, error) {

		return []byte(environment.Env.GetEnv("SECRET")), nil
	})

	db := database.AppDatabase.BD

	var artist model.Artist

	db.Preload(clause.Associations).First(&artist, token.Claims.(jwt.MapClaims)["sub"])

	responseArtist := artistService.ConvertToOneArtistResponse(artist)

	return responseArtist

}

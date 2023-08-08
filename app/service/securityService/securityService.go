package securityService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func GetCurrentUser(c *gin.Context) dto.ArtistResponse {

	tokenString := c.GetHeader("Authorization")

	tokenData := strings.Fields(tokenString)

	token, _ := jwt.Parse(tokenData[1], func(token *jwt.Token) (interface{}, error) {

		return []byte(os.Getenv("SECRET")), nil
	})

	db := database.AppDatabase.BD

	var artist model.Artist

	db.First(&artist, token.Claims.(jwt.MapClaims)["sub"])

	responseArtist := artistService.ConvertToOneArtistResponse(artist)

	return responseArtist

}

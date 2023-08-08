package middleware

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

func JwtMiddleware(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	if len(tokenString) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{
			Status:  http.StatusUnauthorized,
			Message: "Not Unauthorized",
		})
		return
	}

	tokenData := strings.Fields(tokenString)

	if tokenData[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{
			Status:  http.StatusUnauthorized,
			Message: "Invalid type token",
		})
		return
	}

	token, _ := jwt.Parse(tokenData[1], func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{
				Status:  http.StatusUnauthorized,
				Message: "Not Unauthorized",
			})
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{
				Status:  http.StatusUnauthorized,
				Message: "Время токена истекло",
			})

		}

		var artist model.Artist

		db := database.AppDatabase.BD

		result := db.First(&artist, claims["sub"])
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{
				Status:  http.StatusUnauthorized,
				Message: "Not Unauthorized",
			})
		}

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{
			Status:  http.StatusUnauthorized,
			Message: "Not Unauthorized",
		})

	}

	c.Next()

}

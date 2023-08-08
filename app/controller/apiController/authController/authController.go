package authController

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/service/authService"
	"github.com/MeibisuX673/lessonGin/app/service/authService/jwtService"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type AuthController struct {
	AuthService *authService.AuthService
	JWTService  *jwtService.JWTService
}

// Auth
//
//	 @Summary		Auth
//		@Description	auth
//		@Tags			auth
//		@Accept			json
//		@Produce		json
//	 @Param 	body body dto.Auth true "body"
//		@Success		200	{object}    dto.JwtToken
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/auth [post]
func (au *AuthController) Auth(c *gin.Context) {

	var auth dto.Auth

	if err := c.BindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(auth); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	artist, err := au.AuthService.CheckUser(auth)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	jwtToken, err := au.JWTService.CreateJwtToken(*artist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.JwtToken{
		Token: jwtToken,
		Type:  "Bearer",
	})

}

package router

import (
	"github.com/MeibisuX673/lessonGin/app/controller"
	"github.com/MeibisuX673/lessonGin/app/controller/artistController"
	"github.com/gin-gonic/gin"
)

var controllers controller.Controller

func initApiRouter(ge *gin.Engine) {

	controllers = initializationController()

	initArtistRoutes(ge.Group("/api"))

}

func initArtistRoutes(rg *gin.RouterGroup) {

	albums := rg.Group("artists")
	{

		albums.POST("", controllers.ArtistController.POSTArtist)
		albums.GET("", controllers.ArtistController.GETCollectionArtist)

	}

}

func initializationController() controller.Controller {

	return controller.Controller{
		ArtistController: artistController.ArtistController{},
	}
}

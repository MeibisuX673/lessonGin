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

	artists := rg.Group("artists")
	{

		artists.POST("", controllers.ArtistController.POSTArtist)
		artists.GET("", controllers.ArtistController.GETCollectionArtist)
		artists.GET("/:id", controllers.ArtistController.GETArtistById)
		artists.PUT("/:id", controllers.ArtistController.PUTArtist)
		

	}

}

func initializationController() controller.Controller {

	return controller.Controller{
		ArtistController: artistController.ArtistController{},
	}
}

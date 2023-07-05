package router

import (
	"github.com/MeibisuX673/lessonGin/app/controller"
	"github.com/MeibisuX673/lessonGin/app/controller/albumController"
	"github.com/MeibisuX673/lessonGin/app/controller/artistController"
	"github.com/gin-gonic/gin"
)

var controllers controller.Controller

func initApiRouter(ge *gin.Engine) {

	controllers = initializationController()

	groupApi := ge.Group("/api")

	initArtistRoutes(groupApi)
	initAlbumRoutes(groupApi)

}

func initArtistRoutes(rg *gin.RouterGroup) {

	artists := rg.Group("artists")
	{

		artists.POST("", controllers.ArtistController.POSTArtist)
		artists.GET("", controllers.ArtistController.GETCollectionArtist)
		artists.GET("/:id", controllers.ArtistController.GETArtistById)
		artists.PUT("/:id", controllers.ArtistController.PUTArtist)
		artists.DELETE("/:id", controllers.ArtistController.DELETEArtist)

	}

}

func initAlbumRoutes(rg *gin.RouterGroup) {

	albums := rg.Group("/albums")
	{

		albums.POST("", controllers.AlbumController.POSTAlbum)
		albums.GET("", controllers.AlbumController.GETCollectionAlbum)
		albums.PUT("/:id", controllers.AlbumController.PUTAlbum)
		albums.GET("/:id", controllers.AlbumController.GETAlbumById)
		albums.DELETE("/:id", controllers.AlbumController.DELETEAlbum)

	}

}

func initializationController() controller.Controller {

	return controller.Controller{
		ArtistController: artistController.ArtistController{},
		AlbumController:  albumController.AlbumController{},
	}
}

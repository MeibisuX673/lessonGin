package router

import (
	"github.com/MeibisuX673/lessonGin/app/controller"
	"github.com/MeibisuX673/lessonGin/app/controller/albumController"
	"github.com/MeibisuX673/lessonGin/app/controller/artistController"
	"github.com/MeibisuX673/lessonGin/app/controller/authController"
	"github.com/MeibisuX673/lessonGin/app/controller/fileController"
	"github.com/MeibisuX673/lessonGin/app/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var controllers controller.Controller

func initApiRouter(ge *gin.Engine) {

	controllers = initializationController()

	groupSwagger := ge.Group("/swagger/*any")
	groupApi := ge.Group("/api")

	initSwagger(groupSwagger)
	initAuthRoutes(groupApi)
	initArtistRoutes(groupApi)
	initAlbumRoutes(groupApi)
	initFileRoutes(groupApi)

}

func initAuthRoutes(rg *gin.RouterGroup) {

	auth := rg.Group("auth")
	{
		auth.POST("", controllers.AuthController.Auth)
	}

}

func initArtistRoutes(rg *gin.RouterGroup) {

	artists := rg.Group("artists")
	{

		artists.POST("", controllers.ArtistController.POSTArtist)
		artists.GET("", controllers.ArtistController.GETCollectionArtist)
		artists.GET("/:id", controllers.ArtistController.GETArtistById)
		artists.PUT("/:id", middleware.JwtMiddleware, controllers.ArtistController.PUTArtist)
		artists.DELETE("/:id", middleware.JwtMiddleware, controllers.ArtistController.DELETEArtist)

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

func initFileRoutes(rg *gin.RouterGroup) {

	files := rg.Group("/files")
	{
		files.POST("", controllers.FileController.POSTFile)
		files.GET("/:id", controllers.FileController.GETFileById)
		files.GET("", controllers.FileController.GETFileCollection)
		files.DELETE("/:id", controllers.FileController.DELETEFile)
	}
}

func initializationController() controller.Controller {

	return controller.Controller{
		ArtistController: artistController.ArtistController{},
		AlbumController:  albumController.AlbumController{},
		FileController:   fileController.FileController{},
		AuthController:   authController.AuthController{},
	}
}

func initSwagger(rg *gin.RouterGroup) {

	rg.GET("", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

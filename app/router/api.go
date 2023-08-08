package router

import (
	"github.com/MeibisuX673/lessonGin/app/controller"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/albumController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/artistController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/authController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/fileController"
	"github.com/MeibisuX673/lessonGin/app/middleware"
	"github.com/MeibisuX673/lessonGin/app/repository/mySql/album"
	"github.com/MeibisuX673/lessonGin/app/repository/mySql/artist"
	"github.com/MeibisuX673/lessonGin/app/repository/mySql/file"
	"github.com/MeibisuX673/lessonGin/app/service/albumService"
	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/MeibisuX673/lessonGin/app/service/authService"
	"github.com/MeibisuX673/lessonGin/app/service/authService/jwtService"
	"github.com/MeibisuX673/lessonGin/app/service/emailService"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
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

		albums.POST("", middleware.JwtMiddleware, controllers.AlbumController.POSTAlbum)
		albums.GET("", controllers.AlbumController.GETCollectionAlbum)
		albums.PUT("/:id", middleware.JwtMiddleware, controllers.AlbumController.PUTAlbum)
		albums.GET("/:id", controllers.AlbumController.GETAlbumById)
		albums.DELETE("/:id", middleware.JwtMiddleware, controllers.AlbumController.DELETEAlbum)

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

		ArtistController: artistController.ArtistController{
			ArtistService: artistService.New(&artist.ArtistRepository{}),
			QueryService:  queryService.New(),
			FileService:   fileService.New(&file.FileRepository{}),
			AlbumService:  albumService.New(&album.AlbumRepository{}),
			EmailService:  emailService.New(),
		},

		AlbumController: albumController.AlbumController{
			AlbumService:  albumService.New(&album.AlbumRepository{}),
			ArtistService: artistService.New(&artist.ArtistRepository{}),
			QueryService:  queryService.New(),
			FileService:   fileService.New(&file.FileRepository{}),
		},

		FileController: fileController.FileController{
			FileService:  fileService.New(&file.FileRepository{}),
			QueryService: queryService.New(),
		},

		AuthController: authController.AuthController{
			AuthService: authService.New(&artist.ArtistRepository{}),
			JWTService:  &jwtService.JWTService{},
		},
	}
}

func initSwagger(rg *gin.RouterGroup) {

	rg.GET("", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

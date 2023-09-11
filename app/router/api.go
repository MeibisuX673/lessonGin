package router

import (
	"github.com/MeibisuX673/lessonGin/app/controller"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/albumController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/artistController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/authController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/fileController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/musicController"
	"github.com/MeibisuX673/lessonGin/app/middleware"
	"github.com/MeibisuX673/lessonGin/app/repository/album"
	"github.com/MeibisuX673/lessonGin/app/repository/artist"
	"github.com/MeibisuX673/lessonGin/app/repository/file"
	"github.com/MeibisuX673/lessonGin/app/repository/music"
	"github.com/MeibisuX673/lessonGin/app/service/albumService"
	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/MeibisuX673/lessonGin/app/service/audioService"
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
	initMusicRoutes(groupApi)

}

func initAuthRoutes(rg *gin.RouterGroup) {

	auth := rg.Group("auth")
	{
		auth.POST("", controllers.AuthController.Auth)
	}

}

func initMusicRoutes(rg *gin.RouterGroup) {

	musics := rg.Group("musics")
	{
		musics.POST("", middleware.JwtMiddleware, controllers.MusicController.PostMusic)
		musics.GET("", controllers.MusicController.GetCollection)
		musics.GET("/:id", controllers.MusicController.GetMusicById)
		musics.PUT("/:id", middleware.JwtMiddleware, controllers.MusicController.PUTMusic)
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
		files.GET("/music-bytes/:id", controllers.FileController.GetBytesMusicById)
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
			AudioService: audioService.New(&music.MusicRepository{}),
		},

		AuthController: authController.AuthController{
			AuthService: authService.New(&artist.ArtistRepository{}),
			JWTService:  &jwtService.JWTService{},
		},

		MusicController: musicController.MusicController{
			AudioService: audioService.New(&music.MusicRepository{}),
			QueryService: queryService.New(),
			AlbumService: albumService.New(&album.AlbumRepository{}),
			FileService:  fileService.New(&file.FileRepository{}),
		},
	}
}

func initSwagger(rg *gin.RouterGroup) {

	rg.GET("", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

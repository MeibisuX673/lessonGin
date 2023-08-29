package controller

import (
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/albumController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/artistController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/authController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/fileController"
	"github.com/MeibisuX673/lessonGin/app/controller/apiController/musicController"
)

type Controller struct {
	ArtistController artistController.ArtistController
	AlbumController  albumController.AlbumController
	FileController   fileController.FileController
	AuthController   authController.AuthController
	MusicController  musicController.MusicController
}

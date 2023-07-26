package controller

import (
	"github.com/MeibisuX673/lessonGin/app/controller/albumController"
	"github.com/MeibisuX673/lessonGin/app/controller/artistController"
	"github.com/MeibisuX673/lessonGin/app/controller/authController"
	"github.com/MeibisuX673/lessonGin/app/controller/fileController"
)

type Controller struct {
	ArtistController artistController.ArtistController
	AlbumController  albumController.AlbumController
	FileController   fileController.FileController
	AuthController   authController.AuthController
}

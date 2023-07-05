package controller

import (
	"github.com/MeibisuX673/lessonGin/app/controller/albumController"
	"github.com/MeibisuX673/lessonGin/app/controller/artistController"
)

type Controller struct {
	ArtistController artistController.ArtistController
	AlbumController  albumController.AlbumController
}

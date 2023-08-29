package service

import (
	"github.com/MeibisuX673/lessonGin/app/service/albumService"
	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/MeibisuX673/lessonGin/app/service/audioService"
	"github.com/MeibisuX673/lessonGin/app/service/authService"
	"github.com/MeibisuX673/lessonGin/app/service/emailService"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
)

type Services struct {
	AlbumService  *albumService.AlbumService
	ArtistService *artistService.ArtistService
	AuthService   *authService.AuthService
	EmailService  *emailService.EmailService
	FileService   *fileService.FileService
	QueryService  *queryService.QueryService
	AudioService  *audioService.AudioService
}

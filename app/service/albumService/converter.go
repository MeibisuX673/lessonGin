package albumService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func convertToAlbumResponseCollection(albums []model.Album) (responseAlbum []dto.ResponseAlbum) {

	var file *dto.ResponseFile

	for _, album := range albums {

		if album.File != nil {
			file = &dto.ResponseFile{
				ID:   album.File.ID,
				Name: album.File.Name,
				Path: album.File.Path,
			}
		}

		responseAlbum = append(responseAlbum, dto.ResponseAlbum{
			Id:       album.ID,
			Title:    album.Title,
			Price:    album.Price,
			File:     file,
			ArtistID: album.ArtistID,
		})
	}

	return responseAlbum

}

func convertToOneAlbumResponse(album model.Album) (responseAlbum dto.ResponseAlbum) {

	var file *dto.ResponseFile

	if album.File != nil {
		file = &dto.ResponseFile{
			ID:   album.File.ID,
			Name: album.File.Name,
			Path: album.File.Path,
		}
	}

	responseAlbum = dto.ResponseAlbum{
		Id:       album.ID,
		Title:    album.Title,
		Price:    album.Price,
		File:     file,
		ArtistID: album.ArtistID,
	}

	return responseAlbum

}

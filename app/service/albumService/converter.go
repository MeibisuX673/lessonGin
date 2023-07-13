package albumService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func convertToAlbumResponseCollection(albums []model.Album) (responseAlbum []dto.ResponseAlbum) {

	for _, album := range albums {
		responseAlbum = append(responseAlbum, dto.ResponseAlbum{
			Id:    album.ID,
			Title: album.Title,
			Price: album.Price,
			File: &dto.ResponseFile{
				ID:   album.File.ID,
				Name: album.File.Name,
				Path: album.File.Path,
			},
		})
	}

	return responseAlbum

}

func convertToOneAlbumResponse(album model.Album) (responseAlbum dto.ResponseAlbum) {

	responseAlbum = dto.ResponseAlbum{
		Id:    album.ID,
		Title: album.Title,
		Price: album.Price,
		File: &dto.ResponseFile{
			ID:   album.File.ID,
			Name: album.File.Name,
			Path: album.File.Path,
		},
	}

	return responseAlbum

}

package albumService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func convertToAlbumResponseCollection(albums []model.Album) (responseAlbum []dto.AlbumResponse) {

	var file *dto.FileResponse

	for _, album := range albums {

		if album.File != nil {
			file = &dto.FileResponse{
				ID:   album.File.ID,
				Name: album.File.Name,
				Path: album.File.Path,
			}
		}

		responseAlbum = append(responseAlbum, dto.AlbumResponse{
			Id:       album.ID,
			Title:    album.Title,
			Price:    album.Price,
			File:     file,
			ArtistID: album.ArtistID,
		})
	}

	return responseAlbum

}

func convertToOneAlbumResponse(album model.Album) (responseAlbum dto.AlbumResponse) {

	var file *dto.FileResponse

	if album.File != nil {
		file = &dto.FileResponse{
			ID:   album.File.ID,
			Name: album.File.Name,
			Path: album.File.Path,
		}
	}

	responseAlbum = dto.AlbumResponse{
		Id:       album.ID,
		Title:    album.Title,
		Price:    album.Price,
		File:     file,
		ArtistID: album.ArtistID,
	}

	return responseAlbum

}

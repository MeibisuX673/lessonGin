package artistService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func convertToArtistResponseCollection(artists []model.Artist) (responseArtists []dto.ArtistResponse) {

	var responseFiles []dto.FileResponse
	var responseAlbum []dto.AlbumResponse

	for _, artist := range artists {
		if artist.Files != nil {
			for _, file := range artist.Files {
				responseFiles = append(responseFiles, dto.FileResponse{
					ID:   file.ID,
					Name: file.Name,
					Path: file.Path,
				})
			}
		}
		if artist.Albums != nil {
			for _, album := range artist.Albums {
				responseAlbum = append(responseAlbum, dto.AlbumResponse{
					Id:    album.ID,
					Title: album.Title,
					Price: album.Price,
				})
			}
		}
		responseArtists = append(responseArtists, dto.ArtistResponse{
			ID:     artist.ID,
			Name:   artist.Name,
			Age:    artist.Age,
			Files:  responseFiles,
			Albums: responseAlbum,
		})
	}

	return responseArtists

}

func ConvertToOneArtistResponse(artist model.Artist) (responseArtist dto.ArtistResponse) {

	var responseFiles []dto.FileResponse
	var responseAlbum []dto.AlbumResponse

	if artist.Files != nil {
		for _, file := range artist.Files {
			responseFiles = append(responseFiles, dto.FileResponse{
				ID:   file.ID,
				Name: file.Name,
				Path: file.Path,
			})
		}
	}

	if artist.Albums != nil {
		for _, album := range artist.Albums {
			responseAlbum = append(responseAlbum, dto.AlbumResponse{
				Id:       album.ID,
				Title:    album.Title,
				Price:    album.Price,
				ArtistID: artist.ID,
			})
		}
	}

	responseArtist = dto.ArtistResponse{
		ID:     artist.ID,
		Name:   artist.Name,
		Age:    artist.Age,
		Files:  responseFiles,
		Albums: responseAlbum,
	}

	return responseArtist

}

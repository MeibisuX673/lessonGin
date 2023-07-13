package artistService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func convertToArtistResponseCollection(artists []model.Artist) (responseArtists []dto.ResponseArtist) {

	var responseFiles []dto.ResponseFile
	var responseAlbum []dto.ResponseAlbum

	for _, artist := range artists {
		if artist.Files != nil {
			for _, file := range artist.Files {
				responseFiles = append(responseFiles, dto.ResponseFile{
					ID:   file.ID,
					Name: file.Name,
					Path: file.Path,
				})
			}
		}
		if artist.Albums != nil {
			for _, album := range artist.Albums {
				responseAlbum = append(responseAlbum, dto.ResponseAlbum{
					Id:    album.ID,
					Title: album.Title,
					Price: album.Price,
				})
			}
		}
		responseArtists = append(responseArtists, dto.ResponseArtist{
			ID:     artist.ID,
			Name:   artist.Name,
			Age:    artist.Age,
			Files:  responseFiles,
			Albums: responseAlbum,
		})
	}

	return responseArtists

}

func convertToOneArtistResponse(artist model.Artist) (responseArtist dto.ResponseArtist) {

	var responseFiles []dto.ResponseFile
	var responseAlbum []dto.ResponseAlbum

	if artist.Files != nil {
		for _, file := range artist.Files {
			responseFiles = append(responseFiles, dto.ResponseFile{
				ID:   file.ID,
				Name: file.Name,
				Path: file.Path,
			})
		}
	}

	if artist.Albums != nil {
		for _, album := range artist.Albums {
			responseAlbum = append(responseAlbum, dto.ResponseAlbum{
				Id:    album.ID,
				Title: album.Title,
				Price: album.Price,
			})
		}
	}

	responseArtist = dto.ResponseArtist{
		ID:     artist.ID,
		Name:   artist.Name,
		Age:    artist.Age,
		Files:  responseFiles,
		Albums: responseAlbum,
	}

	return responseArtist

}

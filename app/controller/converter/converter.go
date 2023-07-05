package converter

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func ArtistModelToResponse(artist model.Artist) dto.ResponseArtist {

	var responseArtist dto.ResponseArtist = dto.ResponseArtist{
		ID:   int(artist.ID),
		Name: artist.Name,
		Age:  artist.Age,
	}

	return responseArtist

}

func AlbumModelToResponse(album model.Album) dto.ResponseAlbum {

	var responseAlbum dto.ResponseAlbum = dto.ResponseAlbum{
		Id:     int(album.ID),
		Title:  album.Title,
		Artist: ArtistModelToResponse(album.Artist),
		Price:  album.Price,
	}

	return responseAlbum

}

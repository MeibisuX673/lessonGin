package converter

import (
	responseArtist "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func ArtistModelToResponse(artist model.Artist) responseArtist.ResponseArtist {

	var responseArtist responseArtist.ResponseArtist = responseArtist.ResponseArtist{
		ID: artist.ID,
		Name: artist.Name,
		Age:  artist.Age,
	}

	return responseArtist

}

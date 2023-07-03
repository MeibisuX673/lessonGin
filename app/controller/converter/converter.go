package converter

import (
	model2 "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func ArtistModelToResponse(artist model.Artist) model2.ResponseArtist {

	var responseArtist model2.ResponseArtist = model2.ResponseArtist{
		Name: artist.Name,
		Age:  int(artist.Age),
	}

	return responseArtist

}

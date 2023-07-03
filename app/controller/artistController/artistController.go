package artistController

import (
	"github.com/MeibisuX673/lessonGin/app/controller/converter"
	"github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/serivice/artistService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArtistController struct {
}

func (ac *ArtistController) POSTArtist(c *gin.Context) {

	var createArtist = model.CreateArtist{}

	if err := c.BindJSON(&createArtist); err != nil {
		panic(err)
	}

	artist, err := artistService.CreateArtist(&createArtist)
	if err != nil {
		panic(err)
	}

	responseArtist := converter.ArtistModelToResponse(*artist)

	c.JSON(http.StatusCreated, responseArtist)

}

func (ac *ArtistController) GETCollectionArtist(c *gin.Context) {

	var responseArtists []model.ResponseArtist

	artists, err := artistService.GetCollectionArtist()
	if err != nil {
		panic(err)
	}

	for _, value := range artists {

		responseArtists = append(responseArtists, converter.ArtistModelToResponse(*value))
	}

	c.JSON(http.StatusOK, responseArtists)

}

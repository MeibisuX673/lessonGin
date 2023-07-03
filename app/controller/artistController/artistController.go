package artistController

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MeibisuX673/lessonGin/app/controller/converter"
	"github.com/MeibisuX673/lessonGin/app/controller/model"

	"github.com/MeibisuX673/lessonGin/app/serivice/artistService"
	"github.com/gin-gonic/gin"
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

		responseArtists = append(responseArtists, converter.ArtistModelToResponse(value))
	}

	c.JSON(http.StatusOK, responseArtists)

}

func (ac *ArtistController) GETArtistById(c *gin.Context) {

	id := c.Param("id")

	// if !ok {
	// 	c.JSON(http.StatusBadRequest, )
	// }
	artistId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	artist, err := artistService.GetArtistById(artistId)

	fmt.Println(artist)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	var responseArtist model.ResponseArtist
	
	responseArtist = converter.ArtistModelToResponse(*artist)

	c.IndentedJSON(http.StatusOK, responseArtist)

}

func (ac *ArtistController) PUTArtist(c *gin.Context){


	id, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {

		c.AbortWithError(http.StatusBadRequest, err)

	}

	var updateArtist model.UpdateArtist

	c.BindJSON(&updateArtist)

	artist, err := artistService.UpdateArtist(id, updateArtist)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var responseArtist model.ResponseArtist

	responseArtist = converter.ArtistModelToResponse(*artist)

	c.JSON(http.StatusOK, responseArtist)
	


}

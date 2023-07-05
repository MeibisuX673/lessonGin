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
	if artist == nil {
		c.JSON(http.StatusNotFound, model.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		})
		return
	}
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
		return
	}

	for _, value := range artists {

		responseArtists = append(responseArtists, converter.ArtistModelToResponse(value))
	}

	c.JSON(http.StatusOK, responseArtists)

}

func (ac *ArtistController) GETArtistById(c *gin.Context) {

	id := c.Param("id")

	artistId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	artist, err := artistService.GetArtistById(artistId)

	fmt.Println(artist)

	if err != nil {
		c.JSON(http.StatusNotFound, model.Error{
			Status:  http.StatusNotFound,
			Message: "Артист не найден",
		})
		return
	}

	var responseArtist model.ResponseArtist

	responseArtist = converter.ArtistModelToResponse(*artist)

	c.IndentedJSON(http.StatusOK, responseArtist)

}

func (ac *ArtistController) PUTArtist(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.AbortWithError(http.StatusBadRequest, err)
		return

	}

	var updateArtist model.UpdateArtist

	c.BindJSON(&updateArtist)

	artist, err := artistService.UpdateArtist(id, updateArtist)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var responseArtist model.ResponseArtist

	responseArtist = converter.ArtistModelToResponse(*artist)

	c.JSON(http.StatusOK, responseArtist)

}

func (ac *ArtistController) DELETEArtist(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if _, err := artistService.GetArtistById(id); err != nil {

		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	if err := artistService.DeleteArtist(id); err != nil {

		c.AbortWithError(http.StatusInternalServerError, err)
		return

	}

	c.Status(http.StatusNoContent)

}

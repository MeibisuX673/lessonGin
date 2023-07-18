package artistController

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
	"net/http"
	"strconv"

	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/gin-gonic/gin"
)

type ArtistController struct {
}

// POSTArtist  Create Artist
//
//	 @Summary		Create Artist
//		@Description	Create Artist
//		@Tags			artists
//		@Accept			json
//		@Produce		json
//	 @Param 	body body dto.CreateArtist true "body"
//		@Success		201	{object}    dto.ResponseArtist
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/artists [post]
func (ac *ArtistController) POSTArtist(c *gin.Context) {

	var createArtist = dto.CreateArtist{}

	if err := c.BindJSON(&createArtist); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	artist, err := artistService.CreateArtist(&createArtist)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, artist)

}

// GETCollectionArtist  Get Collection Artist
//
//			 @Summary		Get Collection Artist
//				@Description	 Get Collection Artist
//				@Tags			artists
//		     @Param page query string true "page" default(1)
//		     @Param limit query string false "limit" default(5)
//			 @Param filter[id][exact] query string false "filter[id][exact]"
//			 @Param filter[name][partial] query string false "filter[name][partial]"
//			 @Param order[age] query string false "order[age]"
//			 @Param order[created_at] query string false "order[created_at]"
//	         @Param range[age][gt] query string false "range[age][gt]"
//			 @Param range[age][lt] query string false "range[age][lt]"
//				@Accept			json
//				@Produce		json
//				@Success		200	{array}	    dto.ResponseArtist
//				@Failure		500	{object}	dto.Error
//				@Router			/artists [get]
func (ac *ArtistController) GETCollectionArtist(c *gin.Context) {

	queries := queryService.GetQueries(c)

	artists, err := artistService.GetCollectionArtist(*queries)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, artists)

}

// GETArtistById Get Artist
//
//	 @Summary		Get Artist
//		@Description	 Get Artist
//		@Tags			artists
//		@Accept			json
//		@Produce		json
//		@param id path int true "id"
//		@Success		200	{object}	    dto.ResponseArtist
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/artists/{id} [get]
func (ac *ArtistController) GETArtistById(c *gin.Context) {

	id := c.Param("id")

	artistId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	artist, errGetArtistById := artistService.GetArtistById(uint(artistId))
	if errGetArtistById != nil {
		c.JSON(errGetArtistById.GetStatus(), errGetArtistById)
		return
	}

	c.IndentedJSON(http.StatusOK, artist)

}

// PUTArtist Update Artist
//
//		 @Summary		Update Artist
//			@Description	 Update Artist
//			@Tags			artists
//			@Accept			json
//			@Produce		json
//			@param id path int true "id"
//	     @param body body dto.UpdateArtist true "body"
//			@Success		200	{object}	    dto.ResponseArtist
//			@Failure		400	{object}	dto.Error
//			@Failure		404	{object}	dto.Error
//			@Failure		500	{object}	dto.Error
//			@Router			/artists/{id} [put]
func (ac *ArtistController) PUTArtist(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return

	}

	var updateArtist dto.UpdateArtist

	if err := c.BindJSON(&updateArtist); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	artist, errUpdateArtist := artistService.UpdateArtist(id, updateArtist)

	if errUpdateArtist != nil {
		c.JSON(errUpdateArtist.GetStatus(), errUpdateArtist)
		return
	}

	c.JSON(http.StatusOK, artist)

}

// DELETEArtist Delete Artist
//
//	 @Summary		Delete Artist
//		@Description	 Delete Artist
//		@Tags			artists
//		@Accept			json
//		@Produce		json
//		@param id path int true "id"
//		@Success		204
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/artists/{id} [delete]
func (ac *ArtistController) DELETEArtist(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	if err := artistService.DeleteArtist(uint(id)); err != nil {
		c.JSON(err.GetStatus(), err)
		return

	}

	c.Status(http.StatusNoContent)

}

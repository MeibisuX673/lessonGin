package albumController

import (
	"github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/service/albumService"
	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AlbumController struct {
}

// POSTAlbum Create Album
//
//	 @Summary		Create Album
//		@Description	Create Album
//		@Tags			albums
//		@Accept			json
//		@Produce		json
//	 @Param 	body body model.CreateAlbum true "body"
//		@Success		201	{object}	    model.ResponseAlbum
//		@Failure		400	{object}	model.Error
//		@Failure		404	{object}	model.Error
//		@Failure		500	{object}	model.Error
//		@Router			/albums [post]
func (ac *AlbumController) POSTAlbum(c *gin.Context) {

	var createAlbum model.CreateAlbum

	if err := c.BindJSON(&createAlbum); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	artist, err := artistService.GetArtistById(createAlbum.ArtistID)

	if err != nil {
		c.JSON(http.StatusNotFound, model.Error{
			Status:  http.StatusNotFound,
			Message: "Пользователь не найден",
		})
		return
	}

	createAlbum.ArtistID = artist.ID

	album, err := albumService.CreateAlbum(createAlbum)

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, album)

}

// GETCollectionAlbum Get collection Album
//
//	 @Summary		Get collection Album
//		@Description	Get collection Album
//		@Tags			albums
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	    model.ResponseAlbum
//		@Failure		500	{object}	model.Error
//		@Router			/albums [get]
func (ac *AlbumController) GETCollectionAlbum(c *gin.Context) {

	albums, err := albumService.GetCollectionAlbum()

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, albums)

}

// GETAlbumById  Get Album
//
//		 @Summary		Get Album
//			@Description	Get Album
//			@Tags			albums
//			@Accept			json
//	     	@Param id path     int  true "id"
//			@Produce		json
//			@Success		200	{array}	    model.ResponseAlbum
//			@Failure		400	{object}	model.Error
//			@Failure		404	{object}	model.Error
//			@Failure		500	{object}	model.Error
//			@Router			/albums/{id} [get]
func (ac *AlbumController) GETAlbumById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	album, er := albumService.GetAlbumById(id)

	if er != nil {
		c.JSON(er.GetStatus(), er.Error())
		return
	}

	c.JSON(http.StatusOK, album)

}

// PUTAlbum   Update Album
//
//			 @Summary		Update Album
//				@Description	Update Album
//				@Tags			albums
//				@Accept			json
//		     	@Param id path     int  true "id"
//	 @Param 	body body model.CreateAlbum true "body"
//				@Produce		json
//				@Success		200	{array}	    model.ResponseAlbum
//				@Failure		400	{object}	model.Error
//				@Failure		404	{object}	model.Error
//				@Failure		500	{object}	model.Error
//				@Router			/albums/{id} [put]
func (ac *AlbumController) PUTAlbum(c *gin.Context) {

	var updateAlbum model.UpdateAlbum

	if err := c.BindJSON(&updateAlbum); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return

	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	album, errUpdateAlbum := albumService.UpdateAlbum(id, updateAlbum)

	if err != nil {
		c.JSON(errUpdateAlbum.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, album)

}

// DELETEAlbum Delete Album
//
//			 @Summary		Delete Album
//				@Description	Delete Album
//				@Tags			albums
//				@Accept			json
//	 @Param 	id path int true "id"
//				@Produce		json
//				@Success		204
//				@Failure		400	{object}	model.Error
//				@Failure		404	{object}	model.Error
//				@Failure		500	{object}	model.Error
//				@Router			/albums/{id} [delete]
func (ac *AlbumController) DELETEAlbum(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	if err := albumService.DeleteAlbum(id); err != nil {

		c.JSON(http.StatusInternalServerError, model.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	c.Status(http.StatusNoContent)

}

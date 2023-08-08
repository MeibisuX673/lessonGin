package albumController

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/service/albumService"
	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
	"github.com/MeibisuX673/lessonGin/app/service/securityService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AlbumController struct {
	AlbumService  *albumService.AlbumService
	ArtistService *artistService.ArtistService
	QueryService  *queryService.QueryService
	FileService   *fileService.FileService
}

// POSTAlbum Create Album
//
//	    @Summary		Create Album
//	     @security ApiKeyAuth
//		@Description	Create Album
//		@Tags			albums
//		@Accept			json
//		@Produce		json
//	    @Param 	body body model.AlbumCreate true "body"
//		@Success		201	{object}	    model.AlbumResponse
//		@Failure		400	{object}	model.Error
//		@Failure		404	{object}	model.Error
//		@Failure		500	{object}	model.Error
//		@Router			/albums [post]
func (ac *AlbumController) POSTAlbum(c *gin.Context) {

	var createAlbum dto.AlbumCreate

	if err := c.BindJSON(&createAlbum); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	artist, err := ac.ArtistService.GetArtistById(createAlbum.ArtistID)

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	if createAlbum.FileId != nil {
		_, err := ac.FileService.GetFileById(*createAlbum.FileId)
		if err != nil {
			c.JSON(err.GetStatus(), err)
			return
		}
	}

	currentArtist := securityService.GetCurrentUser(c)

	if currentArtist.ID != artist.ID {
		c.JSON(http.StatusForbidden, dto.Error{
			Status:  http.StatusForbidden,
			Message: "Access Denied",
		})
		return
	}

	createAlbum.ArtistID = artist.ID

	album, err := ac.AlbumService.CreateAlbum(createAlbum)

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, album)

}

// GETCollectionAlbum Get collection Album
//
// @Summary		Get collection Album
// @Description	Get collection Album
// @Tags			albums
// @Param page query int true "page number" default(1)
// @Param limit query string false "limit" default(5)
// @Param filter[id][exact] query int false "filter[id][exact]"
// @Param filter[title][partial] query string false "filter[title][partial]"
// @Param filter[artist_id][exact] query int false "filter[artist.id]"
// @Param order[price] query string false "order[price]"
// @Param order[id] query string false "order[id]"
// @Param order[created_at] query string false "order[created_at]"
// @Param range[price][gt] query string false "range[price][gt]"
// @Param range[price][lt] query string false "range[price][lt]"
// @Accept			json
// @Produce		json
// @Success		200	{array}	    model.AlbumResponse
// @Failure		500	{object}	model.Error
// @Router			/albums [get]
func (ac *AlbumController) GETCollectionAlbum(c *gin.Context) {

	queryModel := ac.QueryService.GetQueries(c)

	albums, err := ac.AlbumService.GetCollectionAlbum(*queryModel)

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
//			@Success		200	{array}	    model.AlbumResponse
//			@Failure		400	{object}	model.Error
//			@Failure		404	{object}	model.Error
//			@Failure		500	{object}	model.Error
//			@Router			/albums/{id} [get]
func (ac *AlbumController) GETAlbumById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	album, er := ac.AlbumService.GetAlbumById(uint(id))
	if er != nil {
		c.JSON(er.GetStatus(), er.Error())
		return
	}

	c.JSON(http.StatusOK, album)

}

// PUTAlbum   Update Album
//
//			@Summary		Update Album
//			@security ApiKeyAuth
//			@Description	Update Album
//			@Tags			albums
//			@Accept			json
//		    @Param id path     int  true "id"
//	 		@Param 	body body model.AlbumCreate true "body"
//			@Produce		json
//			@Success		200	{array}	    model.AlbumResponse
//			@Failure		400	{object}	model.Error
//			@Failure		404	{object}	model.Error
//			@Failure		500	{object}	model.Error
//			@Router			/albums/{id} [put]
func (ac *AlbumController) PUTAlbum(c *gin.Context) {

	var updateAlbum dto.UpdateAlbum

	if err := c.BindJSON(&updateAlbum); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	currentArtist := securityService.GetCurrentUser(c)

	artist, errArtist := ac.ArtistService.GetArtistById(*updateAlbum.ArtistID)
	if errArtist != nil {
		c.JSON(errArtist.GetStatus(), errArtist)
		return
	}
	if artist.ID != currentArtist.ID {
		c.JSON(http.StatusForbidden, dto.Error{
			Status:  http.StatusForbidden,
			Message: "Access Denied",
		})
		return
	}

	album, errUpdateAlbum := ac.AlbumService.UpdateAlbum(uint(id), updateAlbum)

	if err != nil {
		c.JSON(errUpdateAlbum.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, album)

}

// DELETEAlbum Delete Album
//
//		@Summary		Delete Album
//		@security ApiKeyAuth
//		@Description	Delete Album
//		@Tags			albums
//		@Accept			json
//	 	@Param 	id path int true "id"
//		@Produce		json
//		@Success		204
//		@Failure		400	{object}	model.Error
//		@Failure		404	{object}	model.Error
//		@Failure		500	{object}	model.Error
//		@Router			/albums/{id} [delete]
func (ac *AlbumController) DELETEAlbum(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	artist := securityService.GetCurrentUser(c)

	if artist.ID != uint(id) {
		c.JSON(http.StatusForbidden, dto.Error{
			Status:  http.StatusForbidden,
			Message: "Access Denied",
		})
		return
	}

	albumsFiles, err := ac.FileService.GetBy(map[string]interface{}{
		"album_id": id,
	})

	if err := ac.FileService.DeleteFileFromDisk(albumsFiles); err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	if err := ac.AlbumService.DeleteAlbum(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return

	}

	c.Status(http.StatusNoContent)

}

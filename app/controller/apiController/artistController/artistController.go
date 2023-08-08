package artistController

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/service/albumService"
	"github.com/MeibisuX673/lessonGin/app/service/emailService"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
	"github.com/MeibisuX673/lessonGin/app/service/securityService"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"github.com/MeibisuX673/lessonGin/app/service/artistService"
	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	ArtistService *artistService.ArtistService
	QueryService  *queryService.QueryService
	FileService   *fileService.FileService
	AlbumService  *albumService.AlbumService
	EmailService  *emailService.EmailService
}

// POSTArtist  Create Artist
//
//	    @Summary		Create Artist
//		@Description	Create Artist
//		@Tags			artists
//		@Accept			json
//		@Produce		json
//	    @Param 	body body dto.ArtistCreate true "body"
//		@Success		201	{object}    dto.ArtistResponse
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/artists [post]
func (ac *ArtistController) POSTArtist(c *gin.Context) {

	var createArtist = dto.ArtistCreate{}

	if err := c.BindJSON(&createArtist); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(createArtist); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if createArtist.FileIds != nil {
		for _, fileId := range createArtist.FileIds {
			_, err := ac.FileService.GetFileById(fileId)
			if err != nil {
				c.JSON(err.GetStatus(), err)
				return
			}
		}
	}

	artist, err := ac.ArtistService.CreateArtist(&createArtist)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	ac.EmailService.SendRegistration(createArtist.Email)

	c.JSON(http.StatusCreated, artist)

}

// GETCollectionArtist  Get Collection Artist
//
//		@Summary		Get Collection Artist
//		@Description	 Get Collection Artist
//		@Tags			artists
//		@Param page query string true "page" default(1)
//		@Param limit query string false "limit" default(5)
//		@Param filter[id][exact] query string false "filter[id][exact]"
//		@Param filter[name][partial] query string false "filter[name][partial]"
//		@Param order[age] query string false "order[age]"
//		@Param order[created_at] query string false "order[created_at]"
//	    @Param range[age][gt] query string false "range[age][gt]"
//		@Param range[age][lt] query string false "range[age][lt]"
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	    dto.ArtistResponse
//		@Failure		500	{object}	dto.Error
//		@Router			/artists [get]
func (ac *ArtistController) GETCollectionArtist(c *gin.Context) {

	queries := ac.QueryService.GetQueries(c)

	artists, err := ac.ArtistService.GetCollectionArtist(*queries)
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
//		@Success		200	{object}	    dto.ArtistResponse
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

	artist, errGetArtistById := ac.ArtistService.GetArtistById(uint(artistId))
	if errGetArtistById != nil {
		c.JSON(errGetArtistById.GetStatus(), errGetArtistById)
		return
	}

	c.JSON(http.StatusOK, artist)

}

// PUTArtist Update Artist
//
//			@Summary		Update Artist
//	     @security ApiKeyAuth
//			@Description	 Update Artist
//			@Tags			artists
//			@Accept			json
//			@Produce		json
//			@param id path int true "id"
//		    @param body body dto.UpdateArtist true "body"
//			@Success		200	{object}	    dto.ArtistResponse
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

	currentArtist := securityService.GetCurrentUser(c)

	if currentArtist.ID != uint(id) {
		c.JSON(http.StatusForbidden, dto.Error{
			Status:  http.StatusForbidden,
			Message: "Access Denied",
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

	artist, errUpdateArtist := ac.ArtistService.UpdateArtist(uint(id), updateArtist)
	if errUpdateArtist != nil {
		c.JSON(errUpdateArtist.GetStatus(), errUpdateArtist)
		return
	}

	c.JSON(http.StatusOK, artist)

}

// DELETEArtist Delete Artist
//
//		    @Summary		Delete Artist
//	     @security ApiKeyAuth
//			@Description	 Delete Artist
//			@Tags			artists
//			@Accept			json
//			@Produce		json
//			@param id path int true "id"
//			@Success		204
//			@Failure		400	{object}	dto.Error
//			@Failure		404	{object}	dto.Error
//			@Failure		500	{object}	dto.Error
//			@Router			/artists/{id} [delete]
func (ac *ArtistController) DELETEArtist(c *gin.Context) {

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

	if err := ac.deleteFiles(artist); err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	if err := ac.ArtistService.DeleteArtist(uint(id)); err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.Status(http.StatusNoContent)

}

func (ac *ArtistController) deleteFiles(artist dto.ArtistResponse) dto.ErrorInterface {

	artistFiles, errGetFiles := ac.FileService.GetBy(map[string]interface{}{
		"artist_id": artist.ID,
	})
	if errGetFiles != nil {
		return errGetFiles
	}

	albums, errGetAlbums := ac.AlbumService.GetBy(map[string]interface{}{
		"artist_id": artist.ID,
	})
	if errGetAlbums != nil {
		return errGetAlbums
	}

	var albumIds []uint
	for _, album := range albums {
		albumIds = append(albumIds, album.Id)
	}

	albumFiles, errGetFiles := ac.FileService.GetBy(map[string]interface{}{
		"album_id": albumIds,
	})
	if errGetFiles != nil {
		return errGetFiles
	}

	if err := ac.FileService.DeleteFileFromDisk(artistFiles); err != nil {
		return err
	}

	if err := ac.FileService.DeleteFileFromDisk(albumFiles); err != nil {
		return err
	}

	return nil

}

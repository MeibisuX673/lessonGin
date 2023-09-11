package musicController

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/service/albumService"
	"github.com/MeibisuX673/lessonGin/app/service/audioService"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
	"github.com/MeibisuX673/lessonGin/app/service/securityService"
	"github.com/MeibisuX673/lessonGin/pkg/slices"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MusicController struct {
	AudioService *audioService.AudioService
	QueryService *queryService.QueryService
	AlbumService *albumService.AlbumService
	FileService  *fileService.FileService
}

// GetMusicById  Get Music
//
//	 @Summary		Get Music
//		@Description	Get Music
//		@Tags			musics
//		@Accept			json
//		@Produce		json
//		@Param id path int true "id"
//		@Success		200	{object}	dto.MusicResponse
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/musics/{id} [get]
func (mc *MusicController) GetMusicById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	result, errMusic := mc.AudioService.GetMusicById(uint(id))
	if err != nil {
		c.JSON(errMusic.GetStatus(), errMusic)
		return
	}

	c.JSON(http.StatusOK, result)

}

// PostMusic  Post Music
//
//	 @Summary		Post Music
//		@Description	Post Music
//		@Security ApiKeyAuth
//		@Tags			musics
//		@Accept			json
//		@Produce		json
//		@Param 	body body dto.MusicCreate true "body"
//		@Success		200	{object}	dto.MusicResponse
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/musics [post]
func (mc *MusicController) PostMusic(c *gin.Context) {

	var createMusic = dto.MusicCreate{}

	if err := c.BindJSON(&createMusic); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(createMusic); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	currentUser := securityService.GetCurrentUser(c)

	if currentUser.ID != createMusic.ArtistID {
		c.JSON(http.StatusForbidden, dto.Error{
			Status:  http.StatusForbidden,
			Message: "Access Denied",
		})
		return
	}

	album, err := mc.AlbumService.GetAlbumById(uint(createMusic.AlbumID))
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	userAlbums := make([]interface{}, len(currentUser.Albums))

	for k, v := range currentUser.Albums {
		userAlbums[k] = v
	}

	if !slices.Contains(userAlbums, *album) {
		c.JSON(http.StatusForbidden, dto.Error{
			Status:  http.StatusForbidden,
			Message: "Access Denied",
		})
		return
	}

	file, err := mc.FileService.GetFileById(createMusic.FileID)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	if !strings.ContainsAny(file.Name, strings.Join(fileService.GetMusicExtension(), "")) {
		c.JSON(http.StatusConflict, dto.Error{
			Status:  http.StatusConflict,
			Message: "File extension not supported",
		})
		return
	}

	response, err := mc.AudioService.CreateMusic(createMusic)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, response)

}

// GetCollection  Get Collection Music
//
//	@Summary		Get Collection Music
//	@Description	 Get Collection Music
//	@Tags			musics
//	@Param page query string true "page" default(1)
//	@Param limit query string false "limit" default(5)
//	@Param filter[id][exact] query string false "filter[id][exact]"
//	@Param filter[name][partial] query string false "filter[name][partial]"
//	@Param order[created_at] query string false "order[created_at]"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	    dto.MusicResponse
//	@Failure		500	{object}	dto.Error
//	@Router			/musics [get]
func (mc *MusicController) GetCollection(c *gin.Context) {

	queryModel := mc.QueryService.GetQueries(c)

	albums, err := mc.AudioService.GetCollectionMusic(*queryModel)

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, albums)

}

// PUTMusic  Put Music
//
//		@Summary		Put Music
//		@Description	 Put Music
//	 @Security ApiKeyAuth
//		@Tags			musics
//				@param id path int true "id"
//			    @param body body dto.MusicUpdate true "body"
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	    dto.MusicResponse
//		@Failure		500	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		403 {object}	dto.Error
//		@Router			/musics/{id} [put]
func (mc *MusicController) PUTMusic(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	music, err := mc.AudioService.GetMusicById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	currentArtist := securityService.GetCurrentUser(c)
	// todo не присваивается id артиста
	log.Fatal(music.ArtistID, "\n", currentArtist)

	if currentArtist.ID != music.ArtistID {
		c.JSON(http.StatusForbidden, dto.Error{
			Status:  http.StatusForbidden,
			Message: "Access Denied",
		})
		return
	}

	var updateMusic dto.MusicUpdate

	if err := c.BindJSON(&updateMusic); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	artist, errUpdateArtist := mc.AudioService.UpdateMusic(uint(id), updateMusic)
	if errUpdateArtist != nil {
		c.JSON(errUpdateArtist.GetStatus(), errUpdateArtist)
		return
	}

	c.JSON(http.StatusOK, artist)

}

// DELETEMusic Delete Music
//
//		    @Summary		Delete Music
//	     @security ApiKeyAuth
//			@Description	 Delete Music
//			@Tags			musics
//			@Accept			json
//			@Produce		json
//			@param id path int true "id"
//			@Success		204
//			@Failure		400	{object}	dto.Error
//			@Failure		404	{object}	dto.Error
//			@Failure		403	{object}	dto.Error
//			@Failure		500	{object}	dto.Error
//			@Router			/musics/{id} [delete]
func (mc *MusicController) DELETEMusic(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	music, err := mc.AudioService.GetMusicById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	currentArtist := securityService.GetCurrentUser(c)

	if currentArtist.ID != music.ArtistID {
		c.JSON(http.StatusForbidden, dto.Error{
			Status:  http.StatusForbidden,
			Message: "Access Denied",
		})
		return
	}

}

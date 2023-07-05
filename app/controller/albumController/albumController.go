package albumController

import (
	"github.com/MeibisuX673/lessonGin/app/controller/converter"
	"github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/serivice/albumService"
	"github.com/MeibisuX673/lessonGin/app/serivice/artistService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AlbumController struct {
}

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

	album, err := albumService.CreateAlbum(createAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	album.Artist = *artist

	response := converter.AlbumModelToResponse(*album)

	c.JSON(http.StatusOK, response)

}

func (ac *AlbumController) GETCollectionAlbum(c *gin.Context) {

	albums, err := albumService.GetCollectionArtist()

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	var response []model.ResponseAlbum

	for _, value := range albums {

		artist, err := artistService.GetArtistById(value.ArtistID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Error{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		value.Artist = *artist

		response = append(response, converter.AlbumModelToResponse(value))

	}

	c.JSON(http.StatusOK, response)

}

func (ac *AlbumController) GETAlbumById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	album, err := albumService.GetAlbumById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, model.Error{
			Status:  http.StatusNotFound,
			Message: "альбом не найден",
		})
		return
	}

	if album == nil {
		c.JSON(http.StatusNotFound, model.Error{
			Status:  http.StatusNotFound,
			Message: "Альбом не найден",
		})
		return
	}

	artist, _ := artistService.GetArtistById(album.ArtistID)
	album.Artist = *artist

	response := converter.AlbumModelToResponse(*album)

	c.JSON(http.StatusOK, response)

}

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

	if album, _ := albumService.GetAlbumById(id); album == nil {
		c.JSON(http.StatusNotFound, model.Error{
			Status:  http.StatusNotFound,
			Message: "Альбом не найден",
		})
		return
	}

	if updateAlbum.ArtistID != nil {
		artist, _ := artistService.GetArtistById(*updateAlbum.ArtistID)
		if artist == nil {
			c.JSON(http.StatusNotFound, model.Error{
				Status:  http.StatusNotFound,
				Message: "Артист не найден",
			})
			return
		}
	}

	album, err := albumService.UpdateAlbum(id, updateAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	artist, _ := artistService.GetArtistById(album.ArtistID)

	album.Artist = *artist

	response := converter.AlbumModelToResponse(*album)

	c.JSON(http.StatusOK, response)

}

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

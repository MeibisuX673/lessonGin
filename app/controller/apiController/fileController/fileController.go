package fileController

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/service/audioService"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/MeibisuX673/lessonGin/app/service/queryService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FileController struct {
	FileService  *fileService.FileService
	QueryService *queryService.QueryService
	AudioService *audioService.AudioService
}

// POSTFile   Create File
//
//	 @Summary		Create File
//		@Description	Create File
//		@Tags			files
//		@Accept			json
//		@Produce		json
//		@Param file	formData file true "file"
//		@Success		201	{object}	    dto.FileResponse
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/files [post]
func (fl *FileController) POSTFile(c *gin.Context) {

	file, err := fl.FileService.UploadFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	createFile := dto.CreateFile{
		Name: file.Name,
		Path: file.Path,
	}

	newFile, errCreate := fl.FileService.CreateFileInDatabase(createFile)
	if errCreate != nil {
		c.JSON(errCreate.GetStatus(), err)
	}

	c.JSON(http.StatusCreated, newFile)

}

// GETFileById    Get File
//
//	 @Summary		Get File
//		@Description	Get File
//		@Tags			files
//		@Accept			json
//		@Produce		json
//		@Param id path int true "id"
//		@Success		200	{object}	dto.FileResponse
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/files/{id} [get]
func (fl *FileController) GETFileById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	file, errGetFileById := fl.FileService.GetFileById(uint(id))

	if errGetFileById != nil {
		c.JSON(errGetFileById.GetStatus(), errGetFileById)
		return
	}

	c.JSON(http.StatusOK, file)

}

// GETFileCollection     Get Collection File
//
//	 @Summary		Get Collection File
//		@Param page query string true "page" default(1)
//		@Param limit query string false "limit" default(5)
//		@Description	Get Collection File
//		@Tags			files
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	dto.FileResponse
//		@Failure		500	{object}	dto.Error
//		@Router			/files [get]
func (fl *FileController) GETFileCollection(c *gin.Context) {

	queries := fl.QueryService.GetQueries(c)

	files, err := fl.FileService.GetFileCollection(*queries)

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, files)

}

// DELETEFile    Delete File
//
//	 @Summary		Delete File
//		@Description	Delete File
//		@Tags			files
//		@Param id path int true "id"
//		@Accept			json
//		@Produce		json
//		@Success		204
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/files/{id} [delete]
func (fl *FileController) DELETEFile(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	if err := fl.FileService.DeleteFile(uint(id)); err != nil {
		c.JSON(err.GetStatus(), err)
	}

	c.Status(http.StatusNoContent)

}

// GetBytesMusicById   Get Bytes
//
//	 @Summary		Get Bytes
//		@Description	Get Bytes
//		@Tags			files
//		@Accept			json
//		@Produce		json
//		@Param id path int true "id"
//		@Success		200	{object}	dto.BytesResponse
//		@Failure		400	{object}	dto.Error
//		@Failure		404	{object}	dto.Error
//		@Failure		500	{object}	dto.Error
//		@Router			/files/music-bytes/{id} [get]
func (fl *FileController) GetBytesMusicById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	file, errFile := fl.FileService.GetFileById(uint(id))
	if err != nil {
		c.JSON(errFile.GetStatus(), errFile)
		return
	}

	bytes := fl.AudioService.GetBytes(file.Name)

	c.JSON(http.StatusOK, dto.BytesResponse{Bytes: bytes})

}

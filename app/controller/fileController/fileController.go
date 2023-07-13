package fileController

import (
	"github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/service/fileService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FileController struct {
}

// POSTFile   Create File
//
//	 @Summary		Create File
//		@Description	Create File
//		@Tags			files
//		@Accept			json
//		@Produce		json
//		@Param file	formData file true "file"
//		@Success		201	{object}	    model.ResponseFile
//		@Failure		400	{object}	model.Error
//		@Failure		404	{object}	model.Error
//		@Failure		500	{object}	model.Error
//		@Router			/files [post]
func (fl *FileController) POSTFile(c *gin.Context) {

	file, err := fileService.UploadFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	var createFile model.CreateFile = model.CreateFile{
		Name: file.Name,
		Path: file.Path,
	}

	newFile, err := fileService.CreateFileInDatabase(createFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
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
//		@Success		200	{object}	model.ResponseFile
//		@Failure		400	{object}	model.Error
//		@Failure		404	{object}	model.Error
//		@Failure		500	{object}	model.Error
//		@Router			/files/{id} [get]
func (fl *FileController) GETFileById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	file, errGetFileById := fileService.GetFileById(uint(id))

	if errGetFileById != nil {
		c.JSON(errGetFileById.GetStatus(), errGetFileById)
		return
	}

	c.JSON(http.StatusOK, file)

}

// GETFileCollection     Get Collection File
//
//	 @Summary		Get Collection File
//		@Description	Get Collection File
//		@Tags			files
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	model.ResponseFile
//		@Failure		500	{object}	model.Error
//		@Router			/files [get]
func (fl *FileController) GETFileCollection(c *gin.Context) {

	files, err := fileService.GetFileCollection()

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
//		@Failure		400	{object}	model.Error
//		@Failure		404	{object}	model.Error
//		@Failure		500	{object}	model.Error
//		@Router			/files/{id} [delete]
func (fl *FileController) DELETEFile(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Status:  http.StatusBadRequest,
			Message: "id должно быть числом",
		})
		return
	}

	if err := fileService.DeleteFile(uint(id)); err != nil {
		c.JSON(err.GetStatus(), err)
	}

	c.Status(http.StatusNoContent)

}

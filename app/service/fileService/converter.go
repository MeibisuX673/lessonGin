package fileService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func convertToFileResponseCollection(files []model.File) (fileResponse []dto.ResponseFile) {

	for _, file := range files {
		fileResponse = append(fileResponse, dto.ResponseFile{
			ID:   file.ID,
			Name: file.Name,
			Path: file.Path,
		})
	}

	return fileResponse

}

func convertToOneFileResponse(file model.File) (fileResponse dto.ResponseFile) {

	fileResponse = dto.ResponseFile{
		ID:   file.ID,
		Name: file.Name,
		Path: file.Path,
	}

	return fileResponse

}

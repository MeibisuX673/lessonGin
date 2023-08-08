package fileService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func convertToFileResponseCollection(files []model.File) (fileResponse []dto.FileResponse) {

	for _, file := range files {
		fileResponse = append(fileResponse, dto.FileResponse{
			ID:   file.ID,
			Name: file.Name,
			Path: file.Path,
		})
	}

	return fileResponse

}

func convertToOneFileResponse(file model.File) (fileResponse dto.FileResponse) {

	fileResponse = dto.FileResponse{
		ID:   file.ID,
		Name: file.Name,
		Path: file.Path,
	}

	return fileResponse

}

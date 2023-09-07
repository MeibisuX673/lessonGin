package audioService

import (
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
)

func convertToMusicResponseCollection(musics []model.Music) (responseMusics []dto.MusicResponse) {

	for _, music := range musics {

		responseMusics = append(responseMusics, dto.MusicResponse{
			ID:       music.ID,
			Name:     music.Name,
			ArtistID: music.ArtistID,
			AlbumID:  music.AlbumID,
			File: dto.FileResponse{
				ID:   music.ID,
				Name: music.File.Name,
				Path: music.File.Path,
			},
		})

	}

	return responseMusics

}

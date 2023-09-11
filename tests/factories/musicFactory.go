package factories

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

func initMusicFactory() {

	Factories[MUSIC_FACTORY] = factory.NewFactory(
		&model.Music{},
	).Attr("Name", func(args factory.Args) (interface{}, error) {
		return randomdata.FirstName(randomdata.Male), nil
	}).Attr("ArtistID", func(args factory.Args) (interface{}, error) {
		return uint(randomdata.Number(0, 1000)), nil
	}).Attr("AlbumID", func(args factory.Args) (interface{}, error) {
		return uint(randomdata.Number(0, 1000)), nil
	})

}

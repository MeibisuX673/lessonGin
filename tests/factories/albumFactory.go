package factories

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

func initAlbumFactory(attributes map[string]interface{}) {

	Factories[ALBUM_FACTORY] = factory.NewFactory(
		&model.Album{},
	).Attr("Title", func(args factory.Args) (interface{}, error) {
		value, ok := attributes["Title"]
		if !ok {
			value = randomdata.FirstName(randomdata.Female)
		}
		return value, nil
	}).Attr("ArtistID", func(args factory.Args) (interface{}, error) {
		value, ok := attributes["ArtistID"]
		if !ok {
			value = uint(randomdata.Number())
		}
		return value, nil
	}).Attr("Price", func(args factory.Args) (interface{}, error) {
		value, ok := attributes["Price"]
		if !ok {
			value = randomdata.Decimal(0, 100)
		}
		return value, nil
	})

}

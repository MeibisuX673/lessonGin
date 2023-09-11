package factories

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

func initArtistFactory() {

	Factories[ARTIST_FACTORY] = factory.NewFactory(
		&model.Artist{},
	).Attr("Name", func(args factory.Args) (interface{}, error) {
		return randomdata.FirstName(randomdata.Male), nil
	}).Attr("Email", func(args factory.Args) (interface{}, error) {
		return randomdata.Email(), nil
	}).Attr("Password", func(args factory.Args) (interface{}, error) {
		return "test", nil
	})

}

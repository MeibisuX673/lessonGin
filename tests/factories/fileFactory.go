package factories

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

func initFileFactory() {

	Factories[FILE_FACTORY] = factory.NewFactory(
		&model.File{},
	).Attr("Path", func(args factory.Args) (interface{}, error) {

		return randomdata.FirstName(randomdata.Male), nil
	}).Attr("Name", func(args factory.Args) (interface{}, error) {

		return randomdata.FirstName(randomdata.Male), nil
	})

}

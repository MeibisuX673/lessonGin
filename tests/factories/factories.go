package factories

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"gorm.io/gorm"
)

type FactoryInterface interface {
	CreateOne()
	CreateMany()
}

type ArtistFactory struct {
}

const (
	ARTIST_FACTORY = "artist"
)

var Factories map[string]*factory.Factory = make(map[string]*factory.Factory)

func InitializationFactory() {

	Factories[ARTIST_FACTORY] = factory.NewFactory(
		&model.Artist{},
	).Attr("Name", func(args factory.Args) (interface{}, error) {
		return randomdata.FirstName(randomdata.Male), nil
	}).Attr("Email", func(args factory.Args) (interface{}, error) {
		return randomdata.Email(), nil
	}).Attr("Password", func(args factory.Args) (interface{}, error) {
		return "test", nil
	}).OnCreate(func(args factory.Args) error {
		db := args.Context().Value("db").(*gorm.DB)
		return db.Create(args.Instance()).Error
	})

}

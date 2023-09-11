package factories

import (
	"github.com/bluele/factory-go/factory"
)

type FactoryInterface interface {
	CreateOne()
	CreateMany()
}

type ArtistFactory struct {
}

const (
	ARTIST_FACTORY = "artist"
	MUSIC_FACTORY  = "music"
	FILE_FACTORY   = "file"
	ALBUM_FACTORY  = "album"
)

var Factories map[string]*factory.Factory = make(map[string]*factory.Factory)

func init() {

	initArtistFactory()
	initAlbumFactory(make(map[string]interface{}))
	initFileFactory()
	initMusicFactory()

}

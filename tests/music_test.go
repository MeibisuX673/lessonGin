package tests

import (
	"context"
	"encoding/json"
	"fmt"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/tests/factories"
	"github.com/bluele/factory-go/factory"
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	API_COLLECTION_URL = "/api/musics"
	API_OBJECT_URL     = "/api/musics/%d"
)

func TestCreateMusic(t *testing.T) {

	t.Cleanup(clean)

	tx := db.BD.Begin()

	ctx := context.WithValue(context.Background(), "db", tx)

	albumFactory := factories.Factories[factories.ALBUM_FACTORY].Attr("ArtistID", func(args factory.Args) (interface{}, error) {
		return uint(1), nil
	}).OnCreate(func(args factory.Args) error {
		db := args.Context().Value("db").(*gorm.DB)
		return db.Create(args.Instance().(*model.Album)).Error
	})

	albums := make([]model.Album, 3)

	for i := 0; i < 3; i++ {
		albumCr, err := albumFactory.CreateWithContext(ctx)
		album := albumCr.(*model.Album)
		if err != nil {
			log.Fatal(err)
		}
		albums[i] = *album
	}

	ctx = context.WithValue(ctx, "albums", albums)

	artistCr, err := factories.Factories[factories.ARTIST_FACTORY].OnCreate(func(args factory.Args) error {
		db := args.Context().Value("db").(*gorm.DB)
		return db.Create(args.Instance().(*model.Artist)).Error
	}).Attr("Albums", func(args factory.Args) (interface{}, error) {
		albums := args.Context().Value("albums").([]model.Album)
		return albums, nil
	}).CreateWithContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	artist := artistCr.(*model.Artist)

	fileCr, err := factories.Factories[factories.FILE_FACTORY].OnCreate(func(args factory.Args) error {
		db := args.Context().Value("db").(*gorm.DB)
		return db.Create(args.Instance()).Error
	}).CreateWithContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	file := fileCr.(*model.File)

	tx.Commit()

	token, err := createJwtToken(*artist)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		http.MethodPost,
		API_COLLECTION_URL,
		strings.NewReader(fmt.Sprintf(
			"{\"albumID\":%d, \"artistID\":%d, \"musicFileID\":%d, \"name\":\"test\"}", artist.Albums[0].ID, artist.ID, file.ID),
		))

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	rout.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

}

func TestGetAll(t *testing.T) {

	t.Cleanup(clean)

	tx := db.BD.Begin()

	ctx := context.WithValue(context.Background(), "db", tx)

	fileCr, err := factories.Factories[factories.FILE_FACTORY].OnCreate(func(args factory.Args) error {
		db := args.Context().Value("db").(*gorm.DB)
		return db.Create(args.Instance().(*model.File)).Error
	}).CreateWithContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	file := fileCr.(*model.File)

	musicFactory := factories.Factories[factories.MUSIC_FACTORY].OnCreate(func(args factory.Args) error {
		db := args.Context().Value("db").(*gorm.DB)
		return db.Create(args.Instance().(*model.Music)).Error
	}).Attr("File", func(args factory.Args) (interface{}, error) {
		return *file, nil
	})

	for i := 0; i < 6; i++ {
		if _, err := musicFactory.CreateWithContext(ctx); err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		http.MethodGet,
		API_COLLECTION_URL,
		strings.NewReader(""),
	)

	rout.ServeHTTP(w, req)

	var musics []dto.MusicResponse

	//log.Fatal(w.Body)
	bytes := w.Body.Bytes()
	json.Unmarshal(bytes, &musics)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, len(musics), 0)

}

func TestGetById(t *testing.T) {

	t.Cleanup(clean)

	tx := db.BD.Begin()

	ctx := context.WithValue(context.Background(), "db", tx)

	fileCr, err := factories.Factories[factories.FILE_FACTORY].OnCreate(func(args factory.Args) error {
		db := args.Context().Value("db").(*gorm.DB)
		return db.Create(args.Instance().(*model.File)).Error
	}).CreateWithContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	file := fileCr.(*model.File)

	musicFactory := factories.Factories[factories.MUSIC_FACTORY].OnCreate(func(args factory.Args) error {
		db := args.Context().Value("db").(*gorm.DB)
		return db.Create(args.Instance().(*model.Music)).Error
	}).Attr("File", func(args factory.Args) (interface{}, error) {
		return *file, nil
	})

	musics := make([]*model.Music, 6)

	for i := 0; i < 6; i++ {
		musicCr, err := musicFactory.CreateWithContext(ctx)
		if err != nil {
			log.Fatal(err)
		}

		musics[i] = musicCr.(*model.Music)
	}

	tx.Commit()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(API_OBJECT_URL, musics[0].ID),
		strings.NewReader(""),
	)

	rout.ServeHTTP(w, req)

	var musicResponse dto.MusicResponse

	bytes := w.Body.Bytes()

	json.Unmarshal(bytes, &musicResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, len(musics), 0)

}

// todo Исправить тест
//func TestUpdate(t *testing.T) {
//
//	t.Cleanup(clean)
//
//	tx := db.BD.Begin()
//
//	ctx := context.WithValue(context.Background(), "db", tx)
//
//	albumCr, err := factories.Factories[factories.ALBUM_FACTORY].Attr("ArtistID", func(args factory.Args) (interface{}, error) {
//		return uint(1), nil
//	}).OnCreate(func(args factory.Args) error {
//		db := args.Context().Value("db").(*gorm.DB)
//		return db.Create(args.Instance().(*model.Album)).Error
//	}).CreateWithContext(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	album := albumCr.(*model.Album)
//
//	ctx = context.WithValue(ctx, "album", *album)
//
//	artistCr, err := factories.Factories[factories.ARTIST_FACTORY].OnCreate(func(args factory.Args) error {
//		db := args.Context().Value("db").(*gorm.DB)
//		return db.Create(args.Instance().(*model.Artist)).Error
//	}).Attr("Albums", func(args factory.Args) (interface{}, error) {
//		album := args.Context().Value("album").(model.Album)
//		return []model.Album{
//			album,
//		}, nil
//	}).CreateWithContext(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	artist := artistCr.(*model.Artist)
//
//	fileCr, err := factories.Factories[factories.FILE_FACTORY].OnCreate(func(args factory.Args) error {
//		db := args.Context().Value("db").(*gorm.DB)
//		return db.Create(args.Instance()).Error
//	}).CreateWithContext(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	file := fileCr.(*model.File)
//
//	musicCr, err := factories.Factories[factories.MUSIC_FACTORY].Attr("ArtistID", func(args factory.Args) (interface{}, error) {
//		return artist.ID, nil
//	}).Attr("AlbumID", func(args factory.Args) (interface{}, error) {
//		return album.ID, nil
//	}).Attr("File", func(factory.Args) (interface{}, error) {
//		return *file, nil
//	}).OnCreate(func(args factory.Args) error {
//		db := args.Context().Value("db").(*gorm.DB)
//		return db.Create(args.Instance().(*model.Music)).Error
//	}).CreateWithContext(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	music := musicCr.(*model.Music)
//
//	tx.Commit()
//
//	token, err := createJwtToken(*artist)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	w := httptest.NewRecorder()
//
//	req, _ := http.NewRequest(
//		http.MethodPut,
//		fmt.Sprintf(API_OBJECT_URL, music.ID),
//		strings.NewReader("\"name\": test"),
//	)
//	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
//	rout.ServeHTTP(w, req)
//
//	var musicResponse dto.MusicResponse
//
//	bytes := w.Body.Bytes()
//	json.Unmarshal(bytes, &musicResponse)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//
//}

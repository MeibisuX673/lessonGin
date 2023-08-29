package tests

import (
	"encoding/json"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/MeibisuX673/lessonGin/app/router"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"github.com/MeibisuX673/lessonGin/tests/factories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var rout *gin.Engine
var db *database.Database

func TestMain(m *testing.M) {

	if err := environment.Env.InitForTest(); err != nil {
		panic(err.Error())
	}

	var errDb error
	db, errDb = database.AppDatabase.Init()
	if errDb != nil {
		panic(errDb.Error())
	}

	factories.InitializationFactory()

	rout = router.AppRouter()

	os.Exit(m.Run())

}

func clean() {

}

func TestCreateArtist(t *testing.T) {

	t.Cleanup(clean)

	expectation := dto.ArtistResponse{
		ID:     1,
		Name:   "Meibisu",
		Age:    120,
		Files:  nil,
		Albums: nil,
	}

	//tx := db.BD.Begin()
	//ctx := context.WithValue(context.Background(), "db", tx)
	//_, err := factories.Factories[factories.ARTIST_FACTORY].CreateWithContext(ctx)
	//if err != nil {
	//	panic(err.Error())
	//}
	//tx.Commit()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/artists", strings.NewReader("{\n  \"age\": 120,\n  \"email\": \"fomesid424@royalka.com\",\n  \"name\": \"Meibisu\",\n  \"password\": \"test\"\n}"))
	//req, _ := http.NewRequest("GET", fmt.Sprintf("/api/artists/%d", artist.ID), nil)
	rout.ServeHTTP(w, req)

	var response dto.ArtistResponse

	bytes := w.Body.Bytes()
	json.Unmarshal(bytes, &response)
	//json.NewDecoder(req.Body).Decode(&response)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, response, expectation)

}

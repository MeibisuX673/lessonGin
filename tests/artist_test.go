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
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var rout *gin.Engine
var bdPath = "./store.db"

func clean() {

	if err := os.Remove(bdPath); err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(bdPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = os.Chmod(bdPath, 0777)
	if err != nil {
		log.Fatal(err)
	}

	initDb()

}

func initDb() {
	var errConnectDb error
	if _, errConnectDb = database.AppDatabase.Init(); errConnectDb != nil {
		panic(errConnectDb.Error())
	}
}

func TestMain(m *testing.M) {

	if err := environment.Env.InitForTest(); err != nil {
		panic(err.Error())
	}

	initDb()

	_ = router.AppRouter()

	factories.InitializationFactory()

	rout = router.AppRouter()

	os.Exit(m.Run())

}

func TestCreateArtist(t *testing.T) {

	expectation := dto.ArtistResponse{
		ID:   1,
		Name: "Meibisu",
		Age:  120,
	}

	//tx := db.BD.Begin()
	//ctx := context.WithValue(context.Background(), "db", tx)
	//_, err := factories.Factories[factories.ARTIST_FACTORY].CreateWithContext(ctx)
	//if err != nil {
	//	panic(err.Error())
	//}
	//tx.Commit()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/artists", strings.NewReader("{\n  \"age\": 120,\n  \"email\": \"test@test.com\",\n  \"name\": \"Meibisu\",\n  \"password\": \"test\"\n}"))
	rout.ServeHTTP(w, req)

	var response dto.ArtistResponse

	bytes := w.Body.Bytes()
	json.Unmarshal(bytes, &response)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, response, expectation)

	t.Cleanup(clean)

}

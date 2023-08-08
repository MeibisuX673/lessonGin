package tests

import (
	"github.com/MeibisuX673/lessonGin/app/router"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func bootKernel() {

	if err := environment.Env.InitForTest(); err != nil {
		panic(err.Error())
	}
	if _, err := database.AppDatabase.Init(); err != nil {
		panic(err.Error())
	}

}

func TestCreateArtist(t *testing.T) {

	bootKernel()

	rout := router.AppRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/artists", strings.NewReader("{\n  \"age\": 120,\n  \"email\": \"fomesid424@royalka.com\",\n  \"name\": \"Meibisu\",\n  \"password\": \"test\"\n}"))
	rout.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

}

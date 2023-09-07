package tests

import (
	"encoding/json"
	dto "github.com/MeibisuX673/lessonGin/app/controller/model"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateArtist(t *testing.T) {

	t.Cleanup(clean)

	expectation := dto.ArtistResponse{
		ID:   1,
		Name: "Meibisu",
		Age:  120,
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/artists", strings.NewReader("{\n  \"age\": 120,\n  \"email\": \"test@test.com\",\n  \"name\": \"Meibisu\",\n  \"password\": \"test\"\n}"))
	rout.ServeHTTP(w, req)

	var response dto.ArtistResponse

	bytes := w.Body.Bytes()
	json.Unmarshal(bytes, &response)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, response, expectation)

}

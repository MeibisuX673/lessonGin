package tests

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/MeibisuX673/lessonGin/app/router"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/MeibisuX673/lessonGin/config/environment"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"testing"
	"time"
)

var rout *gin.Engine
var bdPath = "./store.db"
var db *database.Database

func createJwtToken(artist model.Artist) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["sub"] = artist.ID

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil

}

func clean() {

	if err := os.Remove(bdPath); err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(bdPath)
	if err != nil {
		log.Fatal("clean")
	}
	defer file.Close()

	err = os.Chmod(bdPath, 0777)
	if err != nil {
		log.Fatal("clean")
	}

	initDb()

}

func initDb() {
	var errConnectDb error
	if db, errConnectDb = database.AppDatabase.Init(); errConnectDb != nil {
		panic(errConnectDb.Error())
	}
}

func TestMain(m *testing.M) {

	if err := environment.Env.InitForTest(); err != nil {
		panic(err.Error())
	}

	initDb()

	_ = router.AppRouter()

	rout = router.AppRouter()

	clean()

	os.Exit(m.Run())

}

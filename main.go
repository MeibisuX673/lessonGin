package main

import (
	"github.com/MeibisuX673/lessonGin/app/router"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/MeibisuX673/lessonGin/config/environment"
	_ "github.com/MeibisuX673/lessonGin/docs"
	"github.com/gin-gonic/gin"
	"os"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8081
// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @tokenUrl http://localhost:8081/api/auth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	if _, err := database.AppDatabase.Init(); err != nil {
		panic(err.Error())
	}
	if err := environment.Env.Init(); err != nil {
		panic(err.Error())
	}

	gin.SetMode(os.Getenv(gin.EnvGinMode))
	ge := gin.Default()
	ge.Static("./assets", "./assets")
	ge.MaxMultipartMemory = 8 << 20

	router.AppRouter(ge)

}

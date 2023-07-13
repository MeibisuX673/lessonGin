package main

import (
	"github.com/MeibisuX673/lessonGin/app/router"
	"github.com/MeibisuX673/lessonGin/config/database"
	_ "github.com/MeibisuX673/lessonGin/docs"
	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8080
// @BasePath  /api

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	if _, err := database.AppDatabase.Init(); err != nil {
		panic(err.Error())
	}

	ge := gin.Default()
	ge.Static("./assets", "./assets")
	ge.MaxMultipartMemory = 8 << 20

	router.AppRouter(ge)

}

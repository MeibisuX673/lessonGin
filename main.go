package main

import (
	"github.com/MeibisuX673/lessonGin/app/router"
	"github.com/MeibisuX673/lessonGin/config/database"
	"github.com/gin-gonic/gin"
)

func main() {

	if _, err := database.AppDatabase.Init(); err != nil {
		panic(err.Error())
	}

	ge := gin.Default()

	router.AppRouter(ge)

}

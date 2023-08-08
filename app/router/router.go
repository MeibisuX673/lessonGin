package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func AppRouter() *gin.Engine {

	gin.SetMode(os.Getenv(gin.EnvGinMode))
	ge := gin.Default()
	ge.Static("./assets", "./assets")
	ge.MaxMultipartMemory = 8 << 20

	initApiRouter(ge)

	return ge

}

package router

import (
	"github.com/gin-gonic/gin"
)

func AppRouter(ge *gin.Engine) {

	initApiRouter(ge)
	ge.Run()

}

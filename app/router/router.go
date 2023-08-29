package router

import (
	"context"
	"fmt"
	"github.com/MeibisuX673/lessonGin/config/environment"
	pprof "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"os"
)

func AppRouter() *gin.Engine {

	gin.SetMode(os.Getenv(gin.EnvGinMode))
	ge := gin.Default()
	pprof.Register(ge, "/dev/pprof")
	ge.Static("./assets", "./assets")
	ge.MaxMultipartMemory = 8 << 20

	initApiRouter(ge)

	return ge

}

func RunNgrok(ge *gin.Engine) error {

	ctx := context.Background()
	l, err := ngrok.Listen(
		ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtoken(environment.Env.GetEnv("NGROK_AUTHTOKEN")),
	)
	if err != nil {
		return err
	}

	fmt.Println(l.URL())

	if err := ge.RunListener(l); err != nil {
		return err
	}

	return nil
}

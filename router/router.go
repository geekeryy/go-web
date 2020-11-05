package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web/controllers"
	"go-web/middlewares"
	"net/http"
	_ "net/http/pprof"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	gin.SetMode(viper.GetString("mode"))
	pprof.Register(r)

	demoController:=controllers.NewDemoController()

	r.Use(middlewares.JWTAuth(), middlewares.RequestID())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	user:=r.Group("/user")
	{
		user.GET("/list",demoController.List)
	}

	return r
}

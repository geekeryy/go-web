package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web/middlewares"
	"net/http"
	_ "net/http/pprof"
	"github.com/gin-contrib/pprof"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	gin.SetMode(viper.GetString("mode"))
	pprof.Register(r)

	r.Use(middlewares.JWTAuth(),middlewares.RequestID())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
	return r
}

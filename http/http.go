package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(service ...interface{}) {
	r := gin.Default()

	r.Use()

	r.Any("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"err":  "nil",
		})
	})

}


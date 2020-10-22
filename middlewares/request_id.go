package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

)

const (
	name = "X-Request-Id"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get(name)
		if requestId == "" {
			requestId = uuid.NewV4().String()
			ctx.Writer.Header().Set(name, requestId)
		}
		ctx.Next()
	}
}

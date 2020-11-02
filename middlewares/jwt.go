package middlewares

import (
	"github.com/comeonjy/util/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth .
func JWTAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			return
		}
		tokenEncode := strings.Fields(auth)[0]

		// 校验token
		token, err := jwt.NewJWT().ParseToken(tokenEncode)
		if err != nil {
			return
		}

		ctx.Set("token", token)
		ctx.Next()
	}
}

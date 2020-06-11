package user

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-web/web/models"
)

type UserService struct {

}

func NewUser() *UserService {
	return &UserService{}
}

func (u *UserService) Demo(ctx *gin.Context,req models.BaseReq) {
	name:=ctx.Query("name")
	fmt.Println("UserService Demo name:",name,req)
}
package user

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"go-web/web/models"
)

type UserService struct {

}

func NewUser() *UserService {
	return &UserService{}
}

func (u *UserService) Demo(ctx *gin.Context,req *models.BaseReq) (*models.BaseReq,error) {
	name:=ctx.Query("name")
	fmt.Println("UserService Demo name:",name,req)
	return req,errors.New("no err")
}

func (u *UserService) Demo2(ctx *gin.Context,req []models.BaseReq) ([]models.BaseReq,error) {
	name:=ctx.Query("name")
	fmt.Println("UserService Demo name:",name,req)
	return req,errors.New("no err")
}
package user

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"go-web/web/app/gateway/conf"
	"go-web/web/app/gateway/manager"
	"go-web/web/app/gateway/models"
)

type UserService struct {
	c *conf.Config
	m *manager.Manager
}

func NewUser() *UserService {
	return &UserService{
		c: conf.New(),
		m: manager.New(),
	}
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
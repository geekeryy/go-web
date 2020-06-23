package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"go-web/web/app/gateway/conf"
	"go-web/web/app/gateway/manager"
	"go-web/web/app/gateway/models"
	base_pb "go-web/web/protobuf/base"
	pb_gateway "go-web/web/protobuf/gateway"
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

func (u *UserService) Demo(ctx *gin.Context, req *models.BaseReq) (*models.BaseReq, error) {
	name := ctx.Query("name")
	fmt.Println("UserService Demo name:", name, req)
	return req, errors.New("no err")
}

func (u *UserService) Demo2(ctx *gin.Context, req []models.BaseReq) ([]models.BaseReq, error) {
	name := ctx.Query("name")
	fmt.Println("UserService Demo name:", name, req)
	return req, errors.New("no err")
}

func (u *UserService) Demo3(ctx context.Context, in *base_pb.Base) (*pb_gateway.User, error) {
	return &pb_gateway.User{
		ID: 1,
	}, nil
}

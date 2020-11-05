// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/26 2:08 下午
package controllers

import (
	"github.com/gin-gonic/gin"
	"go-web/services"
)

type DemoController struct {
	demoService *services.DemoService
}

func NewDemoController() *DemoController {
	return &DemoController{
		demoService: services.NewDemoService(),
	}
}

func (c *DemoController) List(ctx *gin.Context) {
	c.demoService.Demo()
}
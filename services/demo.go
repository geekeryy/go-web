// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/26 2:09 下午
package services

import (
	"github.com/sirupsen/logrus"
	"go-web/dao"
)

type DemoService struct {
	demoRepo *dao.DemoRepo
}

func NewDemoService() *DemoService {
	return &DemoService{
		demoRepo: dao.NewDemoRepo(),
	}
}

func (s *DemoService) Demo() {
	logrus.Info("demo")
}
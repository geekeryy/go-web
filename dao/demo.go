// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/26 2:13 下午
package dao

import (
	"github.com/comeonjy/util/mongodb"
	"github.com/comeonjy/util/mysql"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type DemoRepo struct {
	db    *gorm.DB
	mongo *mongo.Collection
}

func NewDemoRepo() *DemoRepo {
	return &DemoRepo{
		db:    mysql.Conn(),
		mongo: mongodb.GetCollection("demo"),
	}
}

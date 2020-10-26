// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/10/26 2:13 下午
package dao

import (
	"github.com/jinzhu/gorm"
	"go-web/util/database/mongodb"
	"go-web/util/database/mysql"
	"go.mongodb.org/mongo-driver/mongo"
)

type DemoRepo struct {
	db *gorm.DB
	mongo *mongo.Collection
}

func NewDemoRepo() *DemoRepo {
	return &DemoRepo{
		db:mysql.GetCollection(),
		mongo: mongodb.GetCollection("demo"),
	}
}

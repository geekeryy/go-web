package igorm

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlConf struct {
	User            string `json:"user"`
	Password        string `json:"password"`
	Host            string `json:"host"`
	Database        string `json:"database"`
	MaxIdleConn     int    `json:"max_idle_conn"`
	MaxOpenConn     int    `json:"max_open_conn"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"` // 连接最大存活时间(单位:秒)
}

func NewMysql(c *MysqlConf) *gorm.DB {

	conn, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Database))
	if err != nil {
		log.Fatal("connMysql err:", err)
	}

	conn.DB().SetMaxIdleConns(c.MaxIdleConn)
	conn.DB().SetMaxOpenConns(c.MaxOpenConn)
	conn.DB().SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second)

	conn.SetLogger(&gormLog{})

	return conn

}

type gormLog struct{}

func (l *gormLog) Print(v ...interface{}) {
	log.Println("GormLog", v)
}

package dao

import (
	"github.com/jinzhu/gorm"

	"go-web/web/app/gateway/conf"
	"go-web/web/library/igorm"
)

type Dao struct {
	c  *conf.Config
	db *gorm.DB
}

func New(c *conf.Config) *Dao {
	return &Dao{
		c:  c,
		db: igorm.NewMysql(c.Mysql),
	}
}

func (d *Dao) Close()  {
	if d.db != nil {
		d.db.Close()
	}
}




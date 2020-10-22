package mysql

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// db instance
var (
	DB   *gorm.DB
	once sync.Once
)

// Init 初始化数据库
func Init() {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			viper.GetString("db.mysql.user"),
			viper.GetString("db.mysql.password"),
			viper.GetString("db.mysql.host"),
			viper.GetInt("db.mysql.port"),
			viper.GetString("db.mysql.dbname"),
		)

		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			logrus.Fatalf("mysql connect failed: %v", err)
		}

		db.DB().SetMaxIdleConns(viper.GetInt("db.mysql.max_idle_conns"))
		db.DB().SetMaxOpenConns(viper.GetInt("db.mysql.max_open_conns"))

		if viper.GetString("mode") == "debug" {
			db.LogMode(true)
		}

		if err = db.DB().Ping(); err != nil {
			logrus.Fatalf("database heartbeat failed: %v", err)
		}

		logrus.Info("mysql connect successfully")
		DB = db
	})
}

// Close method
func Close() error {
	if DB != nil {
		if err := DB.Close(); err != nil {
			return errors.WithStack(err)
		}
	}

	logrus.Info("mysql connect closed")
	return nil
}

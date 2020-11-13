package cmd

import (
	"github.com/comeonjy/util/mysql"
	"github.com/comeonjy/util/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-web/models"
	"go-web/router"
)

func init() {
	// 覆盖配置文件
	httpCmd.PersistentFlags().IntP("port", "p", 0, "http服务端口号")
	viper.BindPFlag("http_port", httpCmd.PersistentFlags().Lookup("port"))
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "go web http server",
	Long:  `go web http服务`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := mysql.Conn().AutoMigrate(models.UserInfo{}).Error; err != nil {
			logrus.Error(err)
		}
		server.Server(router.InitRouter(), viper.GetInt("http_port"))
	},
}

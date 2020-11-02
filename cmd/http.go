package cmd

import (
	"github.com/comeonjy/util/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-web/router"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "go web http server",
	Long:  `go web http服务`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Server(router.InitRouter(),viper.GetInt("http_port"))
	},
}

func init() {
	httpCmd.PersistentFlags().IntP("port", "p", 0, "http服务端口号")
	viper.BindPFlag("http_port", httpCmd.PersistentFlags().Lookup("port"))
}

package cmd

import (
	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/log"
	"github.com/comeonjy/util/mongodb"
	"github.com/comeonjy/util/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config.yaml", "配置文件")

	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(jobCmd)
}

// cmd 执行之前初始化
func initConfig() {
	config.LoadConfig(cfgFile)
}

// 根命令入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}


var rootCmd = &cobra.Command{
	Use: "server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		log.Init(cfg.Log)
		mysql.Init(cfg.Mysql)
		mongodb.Init(cfg.Mongodb)

	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) (err error) {
		errd := mysql.Close()
		errm := mongodb.Disconnect()

		if errd == nil && errm == nil {
			return
		}

		if errd != nil {
			err = errors.Wrap(errd, "close mysql failed")
		}
		if errm != nil {
			err = errors.Wrap(errm, "close mongodb failed")
		}

		return
	},
}
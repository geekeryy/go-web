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

var rootCmd = &cobra.Command{
	Use: "server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Init()
		mysql.Init(config.GetConfig().Mysql)
		mongodb.Init()

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

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config.yaml", "配置文件")

	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(jobCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() {
	config.LoadConfig(cfgFile)
}

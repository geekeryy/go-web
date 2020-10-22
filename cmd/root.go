package cmd

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-web/util/database/mongodb"
	"go-web/util/database/mysql"
	"go-web/util/log"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Init()
		mysql.Init()
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
	rootCmd.AddCommand(migrateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	err := viper.ReadInConfig()
	if err == nil {
		logrus.Info("Using config file:", viper.ConfigFileUsed())
	} else {
		logrus.Fatal(err)
	}
}

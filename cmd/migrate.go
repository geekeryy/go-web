package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-web/util/database/mysql"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "auto migrate mysql table",
	Long:  "自动维护mysql表结构",
	Run: func(cmd *cobra.Command, args []string) {
		createSchema()
	},
}

func createSchema() {

	values := []interface{}{

	}

	if err := mysql.DB.AutoMigrate(values...).Error; err != nil {
		logrus.Fatalf("AutoMigrate failed : %s", err)
	}
}

package cmd

import "github.com/spf13/cobra"

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "go web job server",
	Long: `go web 定时服务`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

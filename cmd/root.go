package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)



var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init()  {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(jobCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}



package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of command line tool",
	Long:  "All software has versions. This is lrbkup's version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lrbkup v0.1")
	},
}

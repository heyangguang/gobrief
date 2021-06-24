package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-brief",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-brief version is 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

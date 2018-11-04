package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "GoAmbari show version",
	Long:  `Show GoAmbari version`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(" Current Version: " + VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

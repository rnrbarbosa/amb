package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "goAmbmbari version",
	Long:  `show amb version`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(" Current Version: " + VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

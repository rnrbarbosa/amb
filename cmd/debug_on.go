package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// onCmd represents the on command
var debugOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Set debug ON",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("on called")
	},
}

func init() {
	debugCmd.AddCommand(debugOnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

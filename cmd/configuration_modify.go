package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// modifyCmd represents the modify command
var configModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify Cluster Configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("modify called")
	},
}

func init() {
	configurationCmd.AddCommand(configModifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

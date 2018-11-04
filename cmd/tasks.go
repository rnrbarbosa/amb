package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Ambari Tasks Operations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tasks called")
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

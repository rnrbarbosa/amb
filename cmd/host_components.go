package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// componentsCmd represents the components command
var hostComponentsCmd = &cobra.Command{
	Use:   "components",
	Short: "List Host Cluster Components",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("components called")
	},
}

func init() {
	hostCmd.AddCommand(hostComponentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

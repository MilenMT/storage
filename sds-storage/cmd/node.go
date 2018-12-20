package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Control node",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node called")
		fmt.Println("Available Commands:")
		for _, c := range cmd.Commands() {
			if c.IsAvailableCommand() {
				fmt.Printf("  %s\t%s\n", c.Name(), c.Short)
			}
		}
	},
}

func init() {
	nodeCmd.AddCommand(nodeListCmd)
	nodeCmd.AddCommand(nodeCheckCmd)
	nodeCmd.AddCommand(nodePingCmd)
	nodeCmd.AddCommand(nodeStoreCmd)
	nodeCmd.AddCommand(nodeInfoCmd)
	nodeCmd.AddCommand(nodeLogCmd)
	nodeCmd.AddCommand(nodeKillCmd)
	nodeCmd.AddCommand(nodeWeightCmd)
	rootCmd.AddCommand(nodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

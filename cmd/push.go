package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdPush)
}

var cmdPush = &cobra.Command{
	Use:   "push <image>",
	Short: "push image to remote repository",
	Long:  "push image to remote repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Implement me!")
	},
}

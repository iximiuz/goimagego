package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdDelete)
}

var cmdDelete = &cobra.Command{
	Use:   "delete <container, image, or layer id>",
	Short: "delete container, image, or layer",
	Long:  "delete container, image, or layer",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := defaultStore().Delete(args[0])
		if err != nil {
			logrus.WithError(err).Fatal("Could not delete")
		}
		fmt.Println("Deleted!")
	},
}

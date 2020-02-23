package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdUnmount)
}

var cmdUnmount = &cobra.Command{
	Use:   "unmount <container-id>",
	Short: "unmount container",
	Long:  "unmount container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mounted, err := defaultStore().Unmount(args[0], true)
		if err != nil {
			logrus.WithError(err).Fatal("Could not unmount container")
		}
		if mounted {
			fmt.Println("Could not unmount container")
		} else {
			fmt.Println("Unmounted!")
		}
	},
}

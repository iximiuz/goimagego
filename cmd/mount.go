package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdMount)
}

var cmdMount = &cobra.Command{
	Use:   "mount <container-id>",
	Short: "mount container",
	Long:  "mount container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mountPoint, err := defaultStore().Mount(args[0], "")
		if err != nil {
			logrus.WithError(err).Fatal("Could not mount container")
		}
		fmt.Println(mountPoint)
	},
}

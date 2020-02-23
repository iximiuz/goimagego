package cmd

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdContainer)
}

var cmdContainer = &cobra.Command{
	Use:   "container <image>",
	Short: "create container",
	Long:  "create container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container, err := defaultStore().CreateContainer(
			"",
			nil,
			args[0],
			"",
			"",
			nil,
		)
		if err != nil {
			logrus.WithError(err).Fatal("Could not create container")
		}
		fmt.Println("Container:")
		spew.Dump(container)
	},
}

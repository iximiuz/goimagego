package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdWipe)
}

var cmdWipe = &cobra.Command{
	Use:   "wipe",
	Short: "wipe the whole storage",
	Long:  "wipe the whole storage",
	Run: func(cmd *cobra.Command, args []string) {
		err := defaultStore().Wipe()
		if err != nil {
			logrus.WithError(err).Fatal("Could not wipe the storage")
		}
		fmt.Println("Wiped!")
	},
}

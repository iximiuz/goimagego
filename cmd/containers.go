package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdContainers)
}

var cmdContainers = &cobra.Command{
	Use:   "containers",
	Short: "list local containers",
	Long:  "list local containers",
	Run: func(cmd *cobra.Command, args []string) {
		containers, err := defaultStore().Containers()
		if err != nil {
			logrus.WithError(err).Fatal("Could not list local containers")
		}
		spew.Println("Containers:")
		spew.Dump(containers)
	},
}

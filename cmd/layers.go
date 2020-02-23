package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdLayers)
}

var cmdLayers = &cobra.Command{
	Use:   "layers",
	Short: "list local layers",
	Long:  "list local layers",
	Run: func(cmd *cobra.Command, args []string) {
		layers, err := defaultStore().Layers()
		if err != nil {
			logrus.WithError(err).Fatal("Could not list local layers")
		}
		spew.Println("Layers:")
		spew.Dump(layers)
	},
}

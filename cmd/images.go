package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdImages)
}

var cmdImages = &cobra.Command{
	Use:   "images",
	Short: "list local images",
	Long:  "list local images",
	Run: func(cmd *cobra.Command, args []string) {
		images, err := defaultStore().Images()
		if err != nil {
			logrus.WithError(err).Fatal("Could not list local images")
		}
		spew.Println("Images:")
		spew.Dump(images)
	},
}

package cmd

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdStore)
}

var cmdStore = &cobra.Command{
	Use:   "store",
	Short: "show store info",
	Long:  "show store info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("root = %v\n", defaultStore().GraphRoot())
		fmt.Printf("run-root = %v\n", defaultStore().RunRoot())
		fmt.Printf("driver = %v\n", defaultStore().GraphDriverName())
		fmt.Printf("driver options = %v\n", defaultStore().GraphOptions())

		status, err := defaultStore().Status()
		if err != nil {
			logrus.WithError(err).Fatal("Could not get store status")
		}
		fmt.Println("status =")
		spew.Dump(status)
	},
}

package cmd

import (
	"fmt"
	"os"

	"github.com/containers/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var _storeOptions storage.StoreOptions
var _defaultStore storage.Store

func init() {
	configureLogLevel()
	initDefaultStoreOptions()
}

func initDefaultStoreOptions() {
	options, err := storage.DefaultStoreOptions(false, 0)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create default image store options")
	}
	_storeOptions = options

	cmdRoot.PersistentFlags().StringVarP(
		&_storeOptions.RunRoot,
		"run-root", "R",
		_storeOptions.RunRoot,
		"image store run root directory",
	)

	cmdRoot.PersistentFlags().StringVarP(
		&_storeOptions.GraphRoot,
		"root", "r",
		_storeOptions.GraphRoot,
		"image store root directory",
	)
	cmdRoot.PersistentFlags().StringVarP(
		&_storeOptions.GraphDriverName,
		"driver", "d",
		"overlay",
		"image store driver (overlay, vfs, etc)",
	)
}

func defaultStore() storage.Store {
	if _defaultStore == nil {
		store, err := storage.GetStore(_storeOptions)
		if err != nil {
			logrus.WithError(err).Fatal("Could not create image store")
		}
		_defaultStore = store
	}
	return _defaultStore
}

func configureLogLevel() {
	logLevelStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevelStr = "info"
	}
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)
}

var cmdRoot = &cobra.Command{
	Use:   "goimagego",
	Short: "goimagego - work with container images in Go",
	Long:  `goimagego - work with container images in Go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Missed or unknown command.\n\n")
		cmd.Help()
	},
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

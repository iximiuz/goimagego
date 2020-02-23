package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/image/copy"
	"github.com/containers/image/signature"
	"github.com/containers/image/storage"
	"github.com/containers/image/transports/alltransports"
	"github.com/containers/image/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdPull)
}

var cmdPull = &cobra.Command{
	Use:   "pull <image>",
	Short: "pull image from remote repository",
	Long:  "pull image from remote repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]
		fmt.Printf("Pulling image %v\n", imageName)

		srcRef, err := alltransports.ParseImageName(imageName)
		if err != nil {
			logrus.WithError(err).Fatal("Could not parse image name")
		}

		systemCtx := &types.SystemContext{}
		policy, err := signature.DefaultPolicy(systemCtx)
		if err != nil {
			logrus.WithError(err).Fatal("Could not create policy")
		}
		policyCtx, err := signature.NewPolicyContext(policy)
		if err != nil {
			logrus.WithError(err).Fatal("Could not create policy context")
		}

		dstRefStr := srcRef.DockerReference().String()
		dstRef, err := storage.Transport.ParseStoreReference(defaultStore(), dstRefStr)
		if err != nil {
			logrus.WithError(err).Fatal("Could not parse local image reference")
		}

		copyOptions := &copy.Options{
			ReportWriter: os.Stdout,
		}
		manifest, err := copy.Image(
			context.Background(),
			policyCtx,
			dstRef,
			srcRef,
			copyOptions,
		)
		if err != nil {
			logrus.WithError(err).Fatal("Could not pull image")
		}
		fmt.Printf("Image pulled - %v\n", string(manifest))
	},
}

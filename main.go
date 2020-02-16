package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/image/copy"
	"github.com/containers/image/signature"
	istorage "github.com/containers/image/storage"
	"github.com/containers/image/transports/alltransports"
	"github.com/containers/image/types"
	"github.com/containers/storage"
	_ "github.com/containers/storage/drivers"
	_ "github.com/containers/storage/drivers/overlay"
	_ "github.com/containers/storage/drivers/register"
	"github.com/containers/storage/pkg/reexec"
	"github.com/sirupsen/logrus"
)

func main() {
	if reexec.Init() {
		panic("HERE 0 - storage.reexec.Init() failed")
	}

	logrus.SetLevel(logrus.DebugLevel)

	imageName := os.Args[1]
	fmt.Printf("Pulling image %v\n", imageName)

	srcRef, err := alltransports.ParseImageName(imageName)
	if err != nil {
		panic(fmt.Errorf("HERE 1 - %w", err))
	}
	fmt.Printf("src image ref: %v\n", srcRef)

	systemCtx := &types.SystemContext{}

	policy, err := signature.DefaultPolicy(systemCtx)
	if err != nil {
		panic(fmt.Errorf("HERE 2 - %w", err))
	}
	policyCtx, err := signature.NewPolicyContext(policy)
	if err != nil {
		panic(fmt.Errorf("HERE 3 - %w", err))
	}

	storeOptions := storage.StoreOptions{
		RunRoot:         "/home/vagrant/images/tmp2/run",
		GraphRoot:       "/home/vagrant/images/tmp2/storage",
		GraphDriverName: "overlay", // "vfs",
		// GraphDriverOptions: []string{"mountopt=nodev"},
		GraphDriverOptions: []string{"overlay.mount_program=/usr/bin/mount"},
	}

	store, err := storage.GetStore(storeOptions)
	if err != nil {
		panic(fmt.Errorf("HERE 4 - %w", err))
	}

	dstRef, err := istorage.Transport.ParseStoreReference(store, srcRef.DockerReference().String())
	if err != nil {
		panic(fmt.Errorf("HERE 5 - %w", err))
	}
	fmt.Printf("dst image ref: %+v\n", dstRef)

	ctx := context.Background()
	copyOptions := &copy.Options{
		ReportWriter: os.Stdout,
	}
	fmt.Println(copy.Image(ctx, policyCtx, dstRef, srcRef, copyOptions))
}

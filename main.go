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
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func main() {
	if reexec.Init() {
		return
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
		RunRoot:         "/home/vagrant/images_tmp/run",
		GraphRoot:       "/home/vagrant/images_tmp/storage",
		GraphDriverName: "overlay", // "vfs",
		// GraphDriverOptions: []string{"overlay.mount_program=/usr/bin/mount"},
	}

	store, err := storage.GetStore(storeOptions)
	if err != nil {
		panic(fmt.Errorf("HERE 4 - %w", err))
	}
	storeStatus, err := store.Status()
	if err != nil {
		panic(fmt.Errorf("HERE 5 - %w", err))
	}
	fmt.Printf("store status:\n%v\n", storeStatus)

	layers, err := store.Layers()
	if err != nil {
		panic(fmt.Errorf("HERE 8 - %w", err))
	}
	fmt.Println("layers:")
	spew.Dump(layers)

	images, err := store.Images()
	if err != nil {
		panic(fmt.Errorf("HERE 9 - %w", err))
	}
	spew.Println("images:")
	spew.Dump(images)

	containers, err := store.Containers()
	if err != nil {
		panic(fmt.Errorf("HERE 10 - %w", err))
	}
	spew.Println("containers:")
	spew.Dump(containers)

	dstRef, err := istorage.Transport.ParseStoreReference(store, srcRef.DockerReference().String())
	if err != nil {
		panic(fmt.Errorf("HERE 6 - %w", err))
	}
	fmt.Printf("dst image ref: %+v\n", dstRef)

	ctx := context.Background()
	copyOptions := &copy.Options{
		ReportWriter: os.Stdout,
	}
	manifest, err := copy.Image(ctx, policyCtx, dstRef, srcRef, copyOptions)
	if err != nil {
		panic(fmt.Errorf("HERE 7 - %w", err))
	}
	fmt.Printf("Image pulled - %v\n", string(manifest))

	cont, err := store.CreateContainer("", nil, images[0].ID, "", "", nil)
	if err != nil {
		panic(fmt.Errorf("HERE 11 - %w", err))
	}
	fmt.Println("container created")
	spew.Dump(cont)

	contDir, err := store.ContainerDirectory(cont.ID)
	if err != nil {
		panic(fmt.Errorf("HERE 12 - %w", err))
	}
	fmt.Printf("container dir = %v\n", contDir)

	contRunDir, err := store.ContainerRunDirectory(cont.ID)
	if err != nil {
		panic(fmt.Errorf("HERE 13 - %w", err))
	}
	fmt.Printf("container run dir = %v\n", contRunDir)

	mountPoint, err := store.Mount(cont.ID, "home_vagrant_images")
	if err != nil {
		panic(fmt.Errorf("HERE 14 - %w", err))
	}
	fmt.Printf("container mounted to %v\n", mountPoint)
}

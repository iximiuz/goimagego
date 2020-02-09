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
)

func main() {
	imageName := os.Args[1]
	fmt.Printf("Pulling image %v\n", imageName)

	srcRef, err := alltransports.ParseImageName(imageName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Image ref: %v\n", srcRef)

	ctx := context.Background()
	systemCtx := &types.SystemContext{}

	policy, err := signature.DefaultPolicy(systemCtx)
	if err != nil {
		panic(err)
	}
	policyCtx, err := signature.NewPolicyContext(policy)
	if err != nil {
		panic(err)
	}

	storeOptions, err := storage.DefaultStoreOptions(true, -1)
	if err != nil {
		panic(err)
	}

	store, err := storage.GetStore(storeOptions)
	if err != nil {
		panic(err)
	}

	dstRef, err := istorage.Transport.ParseStoreReference(store, srcRef.DockerReference().String())
	if err != nil {
		panic(err)
	}

	copyOptions := &copy.Options{}
	fmt.Println(copy.Image(ctx, policyCtx, dstRef, srcRef, copyOptions))
}

package main

import (
	"github.com/containers/storage/pkg/reexec"

	"github.com/iximiuz/goimagego/cmd"
)

func main() {
	if reexec.Init() {
		return
	}

	cmd.Execute()
}

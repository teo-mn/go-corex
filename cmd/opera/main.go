package main

import (
	"fmt"
	"os"

	"github.com/corex-mn/go-corex/cmd/opera/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

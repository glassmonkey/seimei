package main

import (
	"fmt"
	"os"

	"github.com/glassmonkey/seimei"
)

var (
	Version  = "unknown"
	Revision = "unknown"
)

func main() {
	if err := seimei.Run(os.Args, os.Stdin, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "raised error: %s\n", err)
		os.Exit(1)
	}
}

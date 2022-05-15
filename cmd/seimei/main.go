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
	if err := seimei.Run(Version, Revision); err != nil {
		fmt.Fprintf(os.Stderr, "raised error: %s\n", err)
		os.Exit(1)
	}
}

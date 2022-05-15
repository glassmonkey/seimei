package main

import (
	"fmt"
	"os"

	"github.com/glassmonkey/seimei/v2"
)

var (
	Version  = "unknown"
	Revision = "unknown"
)

func main() {
	if err := seimei.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "raised error: %s\n", err)
		os.Exit(1)
	}
}

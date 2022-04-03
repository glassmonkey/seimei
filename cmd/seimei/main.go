package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/glassmonkey/seimei"
)

var (
	Version  = "unknown"
	Revision = "unknown"
)

func main() {
	name := flag.String("name", "", "separate full name(ex. 田中太郎)")
	parse := flag.String("parse", " ", "separate characters")
	flag.Usage = func() {
		_, err := fmt.Fprintf(flag.CommandLine.Output(), "Usage of seimei(%s-%s):\n", Version, Revision)
		if err != nil {
			panic(err)
		}

		flag.PrintDefaults()
	}
	flag.Parse()

	if err := seimei.Run(os.Stdout, *name, *parse); err != nil {
		fmt.Printf("raised error: %s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"github.com/glassmonkey/seimei"
	"os"
)

func main() {
	name := flag.String("name", "", "separate full name(ex. 田中太郎)")
	parse := flag.String("parse", " ", "separate characters")

	flag.Parse()
	if err := seimei.Run(*name, *parse); err != nil {
		fmt.Printf("raised error: %s\n", err)
		os.Exit(1)
	}
}

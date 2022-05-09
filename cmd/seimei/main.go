package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/glassmonkey/seimei"
)

var (
	Version  = "unknown"
	Revision = "unknown"
)

func main() {
	name := flag.String("name", "", "separate full name(ex. 田中太郎)")
	parse := flag.String("parse", " ", "separate characters")
	file := flag.String("file", "", "path to text file of separated by break line undivided name list (ex. /tmp/undivided_name.txt)")
	flag.Usage = func() {
		_, err := fmt.Fprintf(flag.CommandLine.Output(), "Usage of seimei(%s-%s):\n", Version, Revision)
		if err != nil {
			panic(err)
		}

		flag.PrintDefaults()
	}
	flag.Parse()

	if *file != "" {
		data, err := os.ReadFile(*file)

		if err != nil {
			fmt.Printf("raised error: %s\n", err)
			os.Exit(1)
		}

		trimmed := strings.TrimSpace(string(data))
		undividedNames := strings.Split(trimmed, "\n")

		for index, name := range undividedNames {
			if err := seimei.Run(os.Stdout, name, *parse); err != nil {
				fmt.Printf("[Line:%d] raised error : %s\n", index+1, err)
				os.Exit(1)
			}
		}

		return
	}

	if err := seimei.Run(os.Stdout, *name, *parse); err != nil {
		fmt.Printf("raised error: %s\n", err)
		os.Exit(1)
	}
}

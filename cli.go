package seimei

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

var (
	ErrRequiredSubCommand = errors.New("subcommand must be required")
	ErrInvalidSubCommand  = errors.New("subcommand is not defined")
)

type ParseMode string

const (
	NameParse ParseMode = "name"
	FileParse ParseMode = "file"
)

func SetFlagForName(params []string) (Name, ParseString, error) {
	nameCmd := flag.NewFlagSet("name", flag.ContinueOnError)
	name := nameCmd.String("name", "", "separate full name(ex. 田中太郎)")
	parse := nameCmd.String("parse", " ", "separate characters")
	err := nameCmd.Parse(params)
	if err != nil {
		return "", "", fmt.Errorf("name command parse error: %w", err)
	}

	return Name(*name), ParseString(*parse), nil
}

func SetFlagForFile(params []string) (Path, ParseString, error) {
	fileCmd := flag.NewFlagSet("file", flag.ContinueOnError)
	path := fileCmd.String("file", "", "separate full name(ex. /tmp/hoge.csv)")
	parse := fileCmd.String("parse", " ", "separate characters")
	err := fileCmd.Parse(params)
	if err != nil {
		return "", "", fmt.Errorf("file command parse error: %w", err)
	}

	return Path(*path), ParseString(*parse), nil
}

func Run(args []string, stdout, stderr io.Writer) error {
	if len(args) < 2 {
		return ErrRequiredSubCommand
	}
	mode := ParseMode(args[1])

	switch mode {
	case NameParse:
		n, p, err := SetFlagForName(args)
		if err != nil {
			return fmt.Errorf("sub command: %s: %w", os.Args[1], err)
		}
		return ParseName(stdout, stderr, n, p)
	case FileParse:
		f, p, err := SetFlagForFile(args)
		if err != nil {
			return fmt.Errorf("sub command: %s: %w", os.Args[1], err)
		}
		return ParseFile(stdout, stderr, f, p)

	default:
		return fmt.Errorf("sub command: %s: %w", os.Args[1], ErrInvalidSubCommand)
	}
}

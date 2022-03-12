package seimei

import (
	_ "embed"
	"fmt"

	"github.com/glassmonkey/seimei/parser"
)

//go:embed namedivider-python/assets/kanji.csv
var Assets string

//nolint:unparam
func initNameParser(parseString string) (parser.NameParser, error) {
	return parser.NewNameParser(parser.Separator(parseString)), nil
}

func Run(fullname string, parseString string) error {
	p, err := initNameParser(parseString)
	if err != nil {
		return err
	}

	name, err := p.Parse(parser.FullName(fullname))
	if err != nil {
		return fmt.Errorf("happen error: %w", err)
	}

	fmt.Printf("%s\n", name.String())

	return nil
}

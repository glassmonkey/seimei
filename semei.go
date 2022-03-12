package seimei

import (
	"fmt"

	"github.com/glassmonkey/seimei/parser"
)

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
		return err
	}

	fmt.Printf("%s\n", name.String())

	return nil
}

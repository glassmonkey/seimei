package seimei

import "fmt"

//nolint:unparam
func initNameParser(parseString string) (NameParser, error) {
	return NewNameParser(parseString), nil
}

func Run(fullname string, parseString string) error {
	p, err := initNameParser(parseString)
	if err != nil {
		return err
	}

	name, err := p.Parse(fullname)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", name.String())

	return nil
}

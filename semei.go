package seimei

import "fmt"

func initNameParser(parseString string) (NameParser, error) {
	return NewNameParser(parseString), nil
}

func Run(name string, parseString string) error {
	p, err := initNameParser(parseString)
	if err != nil {
		return err
	}
	n, err := p.Parse(name)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", n.String())
	return nil
}
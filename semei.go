package seimei

import "fmt"

func initNameParser() (NameParser, error) {
	return NewNameParser(), nil
}

func Run(name string) error {
	p, err := initNameParser()
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
package seimei

import "fmt"

func initNameParser() (NameParser, error) {
	return NameParser{}, nil
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
	fmt.Printf("%s %s\n", n.FirstName, n.LastName)
	return nil
}
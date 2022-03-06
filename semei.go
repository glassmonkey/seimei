package seimei

import "fmt"

type NameParser struct {
}
func(NameParser) Parse(fullname string) (DividedName, error) {
	return DividedName{}, nil
}

type DividedName struct {
	FirstName string
	LastName string
}


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
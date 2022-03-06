package seimei

type NameParser struct {
}
func(NameParser) Parse(fullname string) (DividedName, error) {
	return DividedName{}, nil
}

type DividedName struct {
	FirstName string
	LastName string
}
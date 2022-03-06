package seimei

import (
	"regexp"
)

type NameParser struct {
	Separator string
	Re *regexp.Regexp
}

func NewNameParser() NameParser {
	re := regexp.MustCompile(`\p{Han}+`)
	return NameParser{
		Separator: " ",
		Re: re,
	}
}


func(n NameParser) Parse(fullname string) (DividedName, error) {
	// Dummy Data. Todo: make from parser.
	return DividedName{
		FirstName: "太郎",
		LastName: "田中",
		Separator: n.Separator,
		Score: 0,
		Algorithm: "",
	}, nil
}

type DividedName struct {
	FirstName string
	LastName string
	Separator string
	Score float64
	Algorithm string
}

func (n DividedName) String() string {
	return n.LastName + n.Separator + n.FirstName
}
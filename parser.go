package seimei

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

var (
	ErrTextLength = errors.New("name length needs at least 2 chars")
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

	if err := n.validate(fullname); err != nil {
		return DividedName{}, fmt.Errorf("parse error: %w", err)
	}

	// Dummy Data. Todo: make from parser.
	return DividedName{
		FirstName: "太郎",
		LastName: "田中",
		Separator: n.Separator,
		Score: 0,
		Algorithm: "",
	}, nil
}

func (n NameParser) validate(fullname string) error {
	v := utf8.RuneCountInString(fullname)

	if v < 2{
		return ErrTextLength
	}
	return nil
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
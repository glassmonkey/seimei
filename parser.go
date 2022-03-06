package seimei

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Algorithm string

const (
	Rule = Algorithm("rule")
	Dummy = Algorithm("dummy")
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
	v, err := n.parseByRule(fullname)
	if err != nil  {
		return DividedName{}, fmt.Errorf("parse error: %w", err)
	}
	if !v.IsZero() {
		return v, nil
	}
	// Dummy Data. Todo: make from parser.
	return DividedName{
		FirstName: "太郎",
		LastName: "田中",
		Separator: n.Separator,
		Score: 0,
		Algorithm: Dummy,
	}, nil
}

func (n NameParser) validate(fullname string)  error {
	v := utf8.RuneCountInString(fullname)

	if v < 2{
		return ErrTextLength
	}
	return nil
}

func (n NameParser) parseByRule(fullname string) (DividedName, error) {
	v := utf8.RuneCountInString(fullname)

	if v == 2{
		 return DividedName{
			FirstName: string([]rune(fullname)[1:2]),
			LastName: string([]rune(fullname)[0:1]),
			Separator: n.Separator,
			Score: 0,
			Algorithm: Rule,
		}, nil
	}
	 return DividedName{}, nil
}

type DividedName struct {
	FirstName string
	LastName string
	Separator string
	Score float64
	Algorithm Algorithm
}

func (n DividedName) String() string {
	return n.LastName + n.Separator + n.FirstName
}

func (n DividedName) IsZero() bool {
	return n == DividedName{}
}
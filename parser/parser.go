package parser

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Algorithm string

const (
	Dummy         = Algorithm("dummy")
	minNameLength = 2
)

var ErrTextLength = errors.New("name length needs at least 2 chars")

type Parser interface {
	Parse(fullname FullName, separator Separator) (DividedName, error)
}
type FullName string

type Separator string

type NameParser struct {
	Parsers   []Parser
	Separator Separator
}

func NewNameParser(separatorString Separator) NameParser {
	s := make([]Parser, 0)
	s = append(s, NewRuleBaseParser())

	return NameParser{
		Parsers:   s,
		Separator: separatorString,
	}
}

func (n NameParser) Parse(fullname FullName) (DividedName, error) {
	if err := n.validate(fullname); err != nil {
		return DividedName{}, fmt.Errorf("parse error: %w", err)
	}

	for _, p := range n.Parsers {
		v, err := p.Parse(fullname, n.Separator)
		if err != nil {
			return DividedName{}, fmt.Errorf("parse error: %w", err)
		}

		if !v.IsZero() {
			return v, nil
		}
	}

	return DividedName{
		FirstName: "太郎",
		LastName:  "田中",
		Separator: n.Separator,
		Score:     0,
		Algorithm: Dummy,
	}, nil
}

func (n NameParser) validate(fullname FullName) error {
	v := utf8.RuneCountInString(string(fullname))

	if v < minNameLength {
		return ErrTextLength
	}

	return nil
}

type DividedName struct {
	FirstName string
	LastName  string
	Separator Separator
	Score     float64
	Algorithm Algorithm
}

func (n DividedName) String() string {
	return n.LastName + string(n.Separator) + n.FirstName
}

//nolint:exhaustivestruct
func (n DividedName) IsZero() bool {
	return n == DividedName{}
}

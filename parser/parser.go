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

type FistName string

type LastName string

func (f FullName) Length() int {
	return utf8.RuneCountInString(string(f))
}

func (f FullName) Split(position int) (LastName, FistName, error) {
	len := f.Length()
	if position < 0 {
		return "", "", errors.New(fmt.Sprintf("position(=%d) must be positive", position))
	}
	if len < position {
		return "", "", errors.New(fmt.Sprintf("position(=%d) is over text length(=%d)", position, len))
	}
	return LastName([]rune(f)[:position]), FistName([]rune(f)[position:]), nil
}

func (f FullName) Slice() []rune {
	return []rune(f)
}

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
	if fullname.Length() < minNameLength {
		return ErrTextLength
	}

	return nil
}

type DividedName struct {
	FirstName FistName
	LastName  LastName
	Separator Separator
	Score     float64
	Algorithm Algorithm
}

func (n DividedName) String() string {
	return string(n.LastName) + string(n.Separator) + string(n.FirstName)
}

//nolint:exhaustivestruct
func (n DividedName) IsZero() bool {
	return n == DividedName{}
}

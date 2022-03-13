package parser

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/glassmonkey/seimei/feature"
)

type Algorithm string

const (
	Dummy         = Algorithm("dummy")
	minNameLength = 2
)

var (
	ErrNameLength    = errors.New("name length needs at least 2 chars")
	ErrSplitPosition = errors.New("split position is invalid")
)

type Parser interface {
	Parse(fullname FullName, separator Separator) (DividedName, error)
}
type FullName string

type FirstName string

func (n FirstName) IsLastName() bool {
	return false
}

func (n FirstName) Length() int {
	return utf8.RuneCountInString(string(n))
}

func (n FirstName) Slice() []rune {
	return []rune(n)
}

type LastName string

func (n LastName) IsLastName() bool {
	return true
}

func (n LastName) Slice() []rune {
	return []rune(n)
}

func (n LastName) Length() int {
	return utf8.RuneCountInString(string(n))
}

func JoinName(lastName LastName, firstName FirstName) FullName {
	return FullName(string(lastName) + string(firstName))
}

func (f FullName) Length() int {
	return utf8.RuneCountInString(string(f))
}

func (f FullName) Split(position int) (LastName, FirstName, error) {
	length := f.Length()

	if position < 0 {
		return "", "", fmt.Errorf("%w: position(=%d) must be positive", ErrSplitPosition, position)
	}

	if length < position {
		return "", "", fmt.Errorf("%w: position(=%d) is over text length(=%d)", ErrSplitPosition, position, length)
	}

	return LastName([]rune(f)[:position]), FirstName([]rune(f)[position:]), nil
}

func (f FullName) Slice() []rune {
	return []rune(f)
}

type Separator string

type NameParser struct {
	Parsers   []Parser
	Separator Separator
}

func NewNameParser(separatorString Separator, m feature.KanjiFeatureManager) NameParser {
	s := make([]Parser, 0)
	s = append(s, NewRuleBaseParser())
	s = append(s, NewStatisticsParser(m))

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
		return ErrNameLength
	}

	return nil
}

type DividedName struct {
	FirstName FirstName
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

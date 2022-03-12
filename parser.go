package seimei

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Algorithm string

const (
	Rule                   = Algorithm("rule")
	Dummy                  = Algorithm("dummy")
	minNameLength          = 2
	separateConditionCount = 2
)

var ErrTextLength = errors.New("name length needs at least 2 chars")

type NameParser struct {
	Separator string
	Re        *regexp.Regexp
}

func NewNameParser(parserString string) NameParser {
	re := regexp.MustCompile(`\p{Han}+`)

	return NameParser{
		Separator: parserString,
		Re:        re,
	}
}

func (n NameParser) Parse(fullname string) (DividedName, error) {
	if err := n.validate(fullname); err != nil {
		return DividedName{}, fmt.Errorf("parse error: %w", err)
	}

	vByRule, err := n.parseByRule(fullname)
	if err != nil {
		return DividedName{}, fmt.Errorf("parse error: %w", err)
	}

	if !vByRule.IsZero() {
		return vByRule, nil
	}
	// Dummy Data. Todo: make from parser.
	return DividedName{
		FirstName: "太郎",
		LastName:  "田中",
		Separator: n.Separator,
		Score:     0,
		Algorithm: Dummy,
	}, nil
}

func (n NameParser) validate(fullname string) error {
	v := utf8.RuneCountInString(fullname)

	if v < minNameLength {
		return ErrTextLength
	}

	return nil
}

//nolint:unparam
//referer: https://github.com/rskmoi/namedivider-python/blob/master/namedivider/name_divider.py#L238
func (n NameParser) parseByRule(fullname string) (DividedName, error) {
	length := utf8.RuneCountInString(fullname)

	if length == minNameLength {
		return DividedName{
			FirstName: string([]rune(fullname)[1:2]),
			LastName:  string([]rune(fullname)[0:1]),
			Separator: n.Separator,
			Score:     1,
			Algorithm: Rule,
		}, nil
	}

	isKanjiList := make([]bool, length)

	for i, c := range []rune(fullname) {
		isKanji := n.Re.MatchString(string(c))
		isKanjiList[i] = isKanji

		if i >= separateConditionCount {
			if isKanjiList[0] != isKanji && isKanjiList[i-1] == isKanji {
				return DividedName{
					FirstName: string([]rune(fullname)[i-1:]),
					LastName:  string([]rune(fullname)[:i-1]),
					Separator: n.Separator,
					Score:     1,
					Algorithm: Rule,
				}, nil
			}
		}
	}
	//nolint:exhaustivestruct
	return DividedName{}, nil
}

type DividedName struct {
	FirstName string
	LastName  string
	Separator string
	Score     float64
	Algorithm Algorithm
}

func (n DividedName) String() string {
	return n.LastName + n.Separator + n.FirstName
}

//nolint:exhaustivestruct
func (n DividedName) IsZero() bool {
	return n == DividedName{}
}

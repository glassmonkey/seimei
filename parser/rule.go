package parser

import (
	"fmt"
	"regexp"
)

const (
	Rule                   = Algorithm("rule")
	separateConditionCount = 2
)

func NewRuleBaseParser() RuleBaseParser {
	re := regexp.MustCompile(`\p{Han}+`)

	return RuleBaseParser{
		re: re,
	}
}

type RuleBaseParser struct {
	re *regexp.Regexp
}

// Parse referer: https://github.com/rskmoi/namedivider-python/blob/master/namedivider/name_divider.py#L238
func (p RuleBaseParser) Parse(fullname FullName, separator Separator) (DividedName, error) {
	length := fullname.Length()

	if length == minNameLength {
		l, f, err := fullname.Split(1)
		if err != nil {
			return DividedName{}, fmt.Errorf("rule parser error: %w", err)
		}

		return DividedName{
			FirstName: f,
			LastName:  l,
			Separator: separator,
			Score:     1,
			Algorithm: Rule,
		}, nil
	}

	isKanjiList := make([]bool, length)

	for i, c := range fullname.Slice() {
		isKanji := p.re.MatchString(string(c))
		isKanjiList[i] = isKanji

		if i >= separateConditionCount {
			if isKanjiList[0] != isKanji && isKanjiList[i-1] == isKanji {
				l, f, err := fullname.Split(i - 1)
				if err != nil {
					return DividedName{}, fmt.Errorf("rule parser error: %w", err)
				}
				return DividedName{
					FirstName: f,
					LastName:  l,
					Separator: separator,
					Score:     1,
					Algorithm: Rule,
				}, nil
			}
		}
	}
	//nolint:exhaustivestruct
	return DividedName{}, nil
}

package seimei

import (
	"regexp"
	"unicode/utf8"
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
	length := utf8.RuneCountInString(string(fullname))

	if length == minNameLength {
		return DividedName{
			FirstName: string([]rune(fullname)[1:2]),
			LastName:  string([]rune(fullname)[0:1]),
			Separator: separator,
			Score:     1,
			Algorithm: Rule,
		}, nil
	}

	isKanjiList := make([]bool, length)

	for i, c := range []rune(fullname) {
		isKanji := p.re.MatchString(string(c))
		isKanjiList[i] = isKanji

		if i >= separateConditionCount {
			if isKanjiList[0] != isKanji && isKanjiList[i-1] == isKanji {
				return DividedName{
					FirstName: string([]rune(fullname)[i-1:]),
					LastName:  string([]rune(fullname)[:i-1]),
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

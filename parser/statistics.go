package parser

import "fmt"

const (
	Statistics = Algorithm("statistics")
)

func NewStatisticsParser() StatisticsParser {
	return StatisticsParser{
		Calculator: Calculator{},
	}
}

type KanjiFeatureCalculator interface {
	Score(lastName LastName, firstName FirstName) float64
}

type StatisticsParser struct {
	Calculator KanjiFeatureCalculator
}

func (s StatisticsParser) Parse(fullname FullName, separator Separator) (DividedName, error) {
	ms := 0.0
	mi := 0

	for i := range fullname.Slice() {
		l, f, err := fullname.Split(i)
		if err != nil {
			return DividedName{}, fmt.Errorf("parse error: %w", err)
		}

		cs := s.Calculator.Score(l, f)

		if cs > ms {
			ms = cs
			mi = i
		}
	}

	l, f, err := fullname.Split(mi)
	if err != nil {
		return DividedName{}, fmt.Errorf("parse error: %w", err)
	}

	return DividedName{
		FirstName: f,
		LastName:  l,
		Separator: separator,
		Score:     ms,
		Algorithm: Statistics,
	}, nil
}

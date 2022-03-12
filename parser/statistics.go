package parser

import "fmt"

const (
	Statistics = Algorithm("statistics")
)

func NewStatisticsParser() StatisticsParser {
	return StatisticsParser{}
}

type StatisticsParser struct{}

func (s StatisticsParser) Parse(fullname FullName, separator Separator) (DividedName, error) {
	ms := 0.0
	mi := 0

	for i := range fullname.Slice() {
		l, f, err := fullname.Split(i)
		if err != nil {
			return DividedName{}, fmt.Errorf("parse error: %w", err)
		}

		cs := s.score(l, f)

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

// stub implement.
func (s StatisticsParser) score(lastName LastName, firstName FirstName) float64 {
	v := float64(len(lastName) - len(firstName))
	if v == 0 {
		return 1
	}

	return 1 / (v * v)
}

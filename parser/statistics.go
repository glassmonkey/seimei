package parser

import (
	"fmt"

	"github.com/glassmonkey/seimei/feature"
)

const (
	Statistics = Algorithm("statistics")
)

func NewStatisticsParser(m feature.KanjiFeatureManager) StatisticsParser {
	return StatisticsParser{
		OrderCalculator: feature.KanjiOrderFeatureCalculator{
			Manager: m,
		},
		LengthCalculator: feature.KanjiLengthFeatureCalculator{
			Manager: m,
		},
	}
}

type StatisticsParser struct {
	OrderCalculator  feature.KanjiOrderFeatureCalculator
	LengthCalculator feature.KanjiLengthFeatureCalculator
}

func (s StatisticsParser) Parse(fullname FullName, separator Separator) (DividedName, error) {
	ms := 0.0
	mi := 1
	features := feature.Features{}

	for i := range fullname.Slice() {
		l, f, err := fullname.Split(i)
		if err != nil {
			return DividedName{}, fmt.Errorf("parse error: %w", err)
		}

		cs, err := s.score(l, f)
		if err != nil {
			return DividedName{}, fmt.Errorf("parse error: %w", err)
		}

		features = append(features, cs)

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
		Score:     features.SoftMax()[mi],
		Algorithm: Statistics,
	}, nil
}

const orderOnlyScoreLength = 4

// Score referer: https://github.com/rskmoi/namedivider-python/blob/master/namedivider/name_divider.py#L206
func (s StatisticsParser) score(lastName LastName, firstName FirstName) (float64, error) {
	fullname := JoinName(lastName, firstName)

	ols, err := s.OrderCalculator.Score(lastName, fullname.Length())
	if err != nil {
		return 0, fmt.Errorf("failed Order Score: %w", err)
	}

	ofs, err := s.OrderCalculator.Score(firstName, fullname.Length())
	if err != nil {
		return 0, fmt.Errorf("failed Order Score: %w", err)
	}

	os := (ols + ofs) / (float64(fullname.Length()) - minNameLength)
	// https://github.com/rskmoi/namedivider-python/blob/d87a488d4696bc26d2f6444ed399d83a6a1911a7/namedivider/name_divider.py#L219
	if fullname.Length() == orderOnlyScoreLength {
		return os, nil
	}

	lls, err := s.LengthCalculator.Score(lastName, fullname.Length())
	if err != nil {
		return 0, fmt.Errorf("failed Length Score: %w", err)
	}

	lfs, err := s.LengthCalculator.Score(firstName, fullname.Length())
	if err != nil {
		return 0, fmt.Errorf("failed Length Score: %w", err)
	}

	ls := (lls + lfs) / float64(fullname.Length())

	return (os + ls) / 2, nil
}

package feature

import (
	"errors"
	"fmt"
)

const (
	CharacterFeatureSize = 1
	OrderFeatureSize     = 6
	LengthFeatureSize    = 8
)

var ErrOrderFeatureInvalidSize = errors.New("order feature's length must be 6")

var ErrLengthFeatureInvalidSize = errors.New("length feature's length must be 8")

type Character string

type KanjiFeatureManager struct {
	KanjiFeatureMap map[Character]KanjiFeature
}

func (m KanjiFeatureManager) Get(c Character) KanjiFeature {
	v, ok := m.KanjiFeatureMap[c]
	if !ok {
		return DefaultKanjiFeature()
	}

	return v
}

func DefaultKanjiFeature() KanjiFeature {
	return KanjiFeature{
		Character: "Default",
		Order:     defaultOrderFeature(),
		Length:    defaultLengthFeature(),
	}
}

type OrderFeature []float64

func NewOrderFeature(d []float64) (OrderFeature, error) {
	if len(d) != OrderFeatureSize {
		return OrderFeature{}, ErrOrderFeatureInvalidSize
	}

	return d, nil
}

func defaultOrderFeature() OrderFeature {
	return []float64{0, 0, 0, 0, 0, 0}
}

type LengthFeature []float64

func NewLengthFeature(d []float64) (LengthFeature, error) {
	if len(d) != LengthFeatureSize {
		return LengthFeature{}, ErrLengthFeatureInvalidSize
	}

	return d, nil
}

func defaultLengthFeature() LengthFeature {
	return []float64{0, 0, 0, 0, 0, 0, 0, 0}
}

type KanjiFeature struct {
	Character Character
	Order     OrderFeature
	Length    LengthFeature
}

func NewKanjiFeature(c Character, o, l []float64) (KanjiFeature, error) {
	of, err := NewOrderFeature(o)
	if err != nil {
		return KanjiFeature{}, fmt.Errorf("failed create kanji feature: %w", err)
	}

	lf, err := NewLengthFeature(l)
	if err != nil {
		return KanjiFeature{}, fmt.Errorf("failed create kanji feature: %w", err)
	}

	return KanjiFeature{
		Character: c,
		Order:     of,
		Length:    lf,
	}, nil
}

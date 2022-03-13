package feature

import (
	"errors"
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
		Order:     defaultFeature(OrderFeatureSize),
		Length:    defaultFeature(LengthFeatureSize),
	}
}

type Features []float64

func defaultFeature(size int) Features {
	return make(Features, size)
}

type KanjiFeature struct {
	Character Character
	Order     Features
	Length    Features
}

func NewKanjiFeature(c Character, o, l []float64) (KanjiFeature, error) {
	if len(o) != OrderFeatureSize {
		return KanjiFeature{}, ErrOrderFeatureInvalidSize
	}

	if len(l) != LengthFeatureSize {
		return KanjiFeature{}, ErrLengthFeatureInvalidSize
	}

	return KanjiFeature{
		Character: c,
		Order:     o,
		Length:    l,
	}, nil
}

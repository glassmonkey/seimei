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

var ErrInvalidFeatureSize = errors.New("feature-to-feature calculations must be the same size")

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

func (f Features) Multiple(mask Features) (Features, error) {
	if len(f) != len(mask) {
		return Features{}, ErrInvalidFeatureSize
	}

	r := make(Features, len(f))
	for i, v := range f {
		r[i] = v * mask[i]
	}

	return r, nil
}

func (f Features) Sum() float64 {
	t := 0.0
	for _, v := range f {
		t += v
	}

	return t
}

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

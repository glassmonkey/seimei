package feature

import (
	"errors"
)

const (
	CharacterFeatureSize = 1
	OrderFeatureSize     = 6
	LengthFeatureSize    = 8
)

var (
	ErrOrderFeatureInvalidSize  = errors.New("order feature's length must be 6")
	ErrLengthFeatureInvalidSize = errors.New("length feature's length must be 8")
	ErrInvalidFeatureSize       = errors.New("feature-to-feature calculations must be the same size")
	ErrOutRangeOrderMask        = errors.New("character position is out of range when creating mask")
	ErrInvalidOrderMask         = errors.New("first character and last character must not be created order mask")
	ErrOutRangeFeatureIndex     = errors.New("character position is out of range when selecting features")
)

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

func (m KanjiFeatureManager) Mask(fullNameLength, charPosition int) (Features, error) {
	if charPosition == 0 || charPosition == fullNameLength-1 {
		return Features{}, ErrInvalidOrderMask
	}

	if charPosition < 0 || charPosition >= fullNameLength {
		return Features{}, ErrOutRangeOrderMask
	}
	//nolint:gomnd
	if fullNameLength == 3 {
		return Features{0, 0, 1, 1, 0, 0}, nil
	}

	if charPosition == 1 {
		return Features{0, 1, 1, 1, 0, 0}, nil
	}

	if charPosition == fullNameLength-2 {
		return Features{0, 0, 1, 1, 1, 0}, nil
	}

	return Features{0, 1, 1, 1, 1, 0}, nil
}

func (m KanjiFeatureManager) SelectFeaturePosition(pieceOfName PartOfNameCharacters, positionInPieceOfName int) (OrderFeatureIndexPosition, error) {
	if positionInPieceOfName < 0 || positionInPieceOfName >= pieceOfName.Length() {
		return 0, ErrOutRangeFeatureIndex
	}

	if positionInPieceOfName == 0 {
		if pieceOfName.IsLastName() {
			return OrderFirstFeatureIndex, nil
		}

		return OrderFirstFeatureIndex.MoveFirstNameIndex(), nil
	}

	if positionInPieceOfName != pieceOfName.Length()-1 {
		if pieceOfName.IsLastName() {
			return OrderMiddleFeatureIndex, nil
		}

		return OrderMiddleFeatureIndex.MoveFirstNameIndex(), nil
	}

	if pieceOfName.IsLastName() {
		return OrderEndFeatureIndex, nil
	}

	return OrderEndFeatureIndex.MoveFirstNameIndex(), nil
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

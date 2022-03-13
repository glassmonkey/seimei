package feature

import (
	"errors"
	"fmt"
)

const (
	CharacterFeatureSize    = 1
	OrderFeatureSize        = 6
	LengthFeatureSize       = 8
	OrderFirstFeatureIndex  = OrderFeatureIndexPosition(0)
	OrderMiddleFeatureIndex = OrderFeatureIndexPosition(1)
	OrderEndFeatureIndex    = OrderFeatureIndexPosition(2)
)

var (
	ErrOrderFeatureInvalidSize  = errors.New("order feature's length must be 6")
	ErrLengthFeatureInvalidSize = errors.New("length feature's length must be 8")
	ErrInvalidFeatureSize       = errors.New("feature-to-feature calculations must be the same size")
	ErrOutRangeOrderMask        = errors.New("character position is out of range when creating mask")
	ErrInvalidOrderMask         = errors.New("first character and last character must not be created order mask")
	ErrOutRangeFeatureIndex     = errors.New("character position is out of range when selecting features")
)

type OrderFeatureIndexPosition int

type LengthFeatureIndexPosition int

func (i OrderFeatureIndexPosition) MoveFirstNameIndex() OrderFeatureIndexPosition {
	return i + OrderFeatureSize/2
}

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

func (m KanjiFeatureManager) OrderMask(fullNameLength, charPosition int) (Features, error) {
	if charPosition == 0 || charPosition == fullNameLength-1 {
		return Features{}, ErrInvalidOrderMask
	}

	if charPosition < 0 || charPosition >= fullNameLength {
		return Features{}, ErrOutRangeOrderMask
	}
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

func (m KanjiFeatureManager) LengthMask(fullNameLength, charPosition int) (Features, error) {
	if charPosition < 0 || charPosition >= fullNameLength {
		return Features{}, ErrOutRangeOrderMask
	}
	minLastName := charPosition + 1
	maxLastName := fullNameLength - 1
	lf := m.maskLengthFuturesForPart(minLastName, maxLastName)

	minFirstName := fullNameLength - charPosition
	maxFirstName := fullNameLength - 1
	ff := m.maskLengthFuturesForPart(minFirstName, maxFirstName)
	for _, ffv := range ff {
		lf = append(lf, ffv)
	}
	return lf, nil
}

func (m KanjiFeatureManager) maskLengthFuturesForPart(min, max int) []float64 {
	minv := min
	maxv := max
	if maxv > LengthFeatureSize/2 {
		maxv = LengthFeatureSize / 2
	}
	f := []float64{0, 0, 0, 0}
	if minv <= maxv {
		for i := minv - 1; i < maxv; i++ {
			f[i] = 1
		}
	}
	return f
}

func (m KanjiFeatureManager) SelectOrderFeaturePosition(pieceOfName PartOfNameCharacters, positionInPieceOfName int) (OrderFeatureIndexPosition, error) {
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

func (k KanjiFeature) GetOrderValue(p OrderFeatureIndexPosition, mask Features) (float64, error) {
	if p < 0 || p >= OrderFeatureSize {
		return 0.0, ErrOutRangeFeatureIndex
	}

	os, err := k.Order.Multiple(mask)
	if err != nil {
		return 0.0, fmt.Errorf("failed order value: %w", err)
	}

	total := os.Sum()
	if total == 0 {
		return 0, nil
	}

	return os[p] / total, nil
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

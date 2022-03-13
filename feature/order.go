package feature

import (
	"errors"
	"fmt"
)

var (
	ErrOutRangeOrderMask    = errors.New("character position is out of range when creating mask")
	ErrInvalidOrderMask     = errors.New("first character and last character must not be created order mask")
	ErrOutRangeFeatureIndex = errors.New("character position is out of range when selecting features")
)

type OrderFeatureIndexPosition int

func (i OrderFeatureIndexPosition) MoveFirstNameIndex() OrderFeatureIndexPosition {
	return i + OrderFeatureSize/2
}

const (
	OrderFirstFeatureIndex  = OrderFeatureIndexPosition(0)
	OrderMiddleFeatureIndex = OrderFeatureIndexPosition(1)
	OrderEndFeatureIndex    = OrderFeatureIndexPosition(2)
)

type KanjiFeatureOrderCalculator struct {
	Manager KanjiFeatureManager
}

func (fc KanjiFeatureOrderCalculator) Mask(fullNameLength, charPosition int) (Features, error) {
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

func (fc KanjiFeatureOrderCalculator) SelectFeaturePosition(pieceOfName PartOfNameCharacters, positionInPieceOfName int) (OrderFeatureIndexPosition, error) {
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

func (fc KanjiFeatureOrderCalculator) Score(pieceOfName PartOfNameCharacters, fullNameLength int) (float64, error) {
	score := 0.0
	offset := 0
	if !pieceOfName.IsLastName() {
		offset = fullNameLength - pieceOfName.Length()
	}

	for i, c := range pieceOfName.Slice() {
		ci := i + offset
		if ci == 0 || ci == fullNameLength-1 {
			continue
		}

		mask, err := fc.Mask(fullNameLength, ci)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		index, err := fc.SelectFeaturePosition(pieceOfName, i)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		os, err := fc.Manager.Get(Character(c)).Order.Multiple(mask)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		total := os.Sum()
		if total == 0 {
			continue
		}
		v := os[index] / total
		score += v
	}

	return score, nil
}

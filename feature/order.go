package feature

import (
	"errors"
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

type KanjiFeatureOrderCalculator struct{}

func (fc KanjiFeatureOrderCalculator) Mask(fullNameLength, charPosition int) (OrderFeature, error) {
	if charPosition == 0 || charPosition == fullNameLength-1 {
		return OrderFeature{}, ErrInvalidOrderMask
	}

	if charPosition < 0 || charPosition >= fullNameLength {
		return OrderFeature{}, ErrOutRangeOrderMask
	}
	//nolint:gomnd
	if fullNameLength == 3 {
		return NewOrderFeature([]float64{0, 0, 1, 1, 0, 0})
	}

	if charPosition == 1 {
		return NewOrderFeature([]float64{0, 1, 1, 1, 0, 0})
	}

	if charPosition == fullNameLength-2 {
		return NewOrderFeature([]float64{0, 0, 1, 1, 1, 0})
	}

	return NewOrderFeature([]float64{0, 1, 1, 1, 1, 0})
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

// Score patch work implementation.
func (fc KanjiFeatureOrderCalculator) Score(pieceOfName PartOfNameCharacters, fullNameLength, startPosition int) (float64, error) {
	// isLastName := startPosition == 0
	score := 0.0
	/*
		for i, c := range pieceOfName.Slice() {
			ci := i + startPosition
			if ci == 0 {
				continue
			}
			if ci == fullNameLength-1 {
				continue
			}
			mask, err := fc.Mask(fullNameLength, ci)
			if err != nil {
				return 0.0, fmt.Errorf("failed order score: %w", err)
			}
			// idx := c.Status(pieceOfName, i, isLastName)
			// currentScores := self.kanji_dict.get(c, self.default_kanji).order_counts * mask
			//scores = fucn(c, mask)
			//total = score.sum()
			// if total == 0 continue
			// scores.get(idx) / total
			score += 0
		}*/
	return score, nil
}

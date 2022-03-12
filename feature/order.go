package feature

import (
	"errors"
)

var (
	ErrOutRangeOrderMask    = errors.New("character position is out of range when creating mask")
	ErrInvalidOrderMask     = errors.New("first character and last character must not be created order mask")
	ErrOutRangeFeatureIndex = errors.New("character position is out of range when selecting features")
)

type IndexPosition int

func (i IndexPosition) MoveFirstNameIndex() IndexPosition {
	return i + featureSize/2
}

const (
	featureSize        = 6
	FirstFeatureIndex  = IndexPosition(0)
	MiddleFeatureIndex = IndexPosition(1)
	EndFeatureIndex    = IndexPosition(2)
)

type KanjiFeatureOrderCalculator struct{}

func (fc KanjiFeatureOrderCalculator) Mask(fullNameLength, charPosition int) ([]float64, error) {
	if charPosition == 0 || charPosition == fullNameLength-1 {
		return []float64{}, ErrInvalidOrderMask
	}

	if charPosition < 0 || charPosition >= fullNameLength {
		return []float64{}, ErrOutRangeOrderMask
	}
	//nolint:gomnd
	if fullNameLength == 3 {
		return []float64{0, 0, 1, 1, 0, 0}, nil
	}

	if charPosition == 1 {
		return []float64{0, 1, 1, 1, 0, 0}, nil
	}

	if charPosition == fullNameLength-2 {
		return []float64{0, 0, 1, 1, 1, 0}, nil
	}

	return []float64{0, 1, 1, 1, 1, 0}, nil
}

func (fc KanjiFeatureOrderCalculator) SelectFeaturePosition(pieceOfName PartOfNameCharacters, positionInPieceOfName int) (IndexPosition, error) {
	if positionInPieceOfName < 0 || positionInPieceOfName >= pieceOfName.Length() {
		return 0, ErrOutRangeFeatureIndex
	}

	if positionInPieceOfName == 0 {
		if pieceOfName.IsLastName() {
			return FirstFeatureIndex, nil
		}

		return FirstFeatureIndex.MoveFirstNameIndex(), nil
	}

	if positionInPieceOfName != pieceOfName.Length()-1 {
		if pieceOfName.IsLastName() {
			return MiddleFeatureIndex, nil
		}

		return MiddleFeatureIndex.MoveFirstNameIndex(), nil
	}

	if pieceOfName.IsLastName() {
		return EndFeatureIndex, nil
	}

	return EndFeatureIndex.MoveFirstNameIndex(), nil
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

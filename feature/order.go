package feature

import (
	"errors"
)

var (
	ErrOutRangeOrderMaskError = errors.New("character position is out of range")
	ErrInvalidOrderMaskError  = errors.New("first character and last character must not be created order mask")
)

type KanjiFeatureOrderCalculator struct{}

func (fc KanjiFeatureOrderCalculator) Mask(fullNameLength, charPosition int) ([]float64, error) {
	if charPosition == 0 || charPosition == fullNameLength-1 {
		return []float64{}, ErrInvalidOrderMaskError
	}
	if charPosition < 0 || charPosition >= fullNameLength {
		return []float64{}, ErrOutRangeOrderMaskError
	}

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

func (fc KanjiFeatureOrderCalculator) Status(pieceOfName MultiCharacters, positionInPieceOfName int, isLastName bool) int {
	if positionInPieceOfName == 0 {
		if isLastName {
			return 0
		}
		return 3
	}
	if positionInPieceOfName == pieceOfName.Length()-1 {
		if isLastName {
			return 2
		}
		return 5
	}
	if isLastName {
		return 1
	}
	return 4
}

// Score patch work implementation.
func (fc KanjiFeatureOrderCalculator) Score(pieceOfName MultiCharacters, fullNameLength, startPosition int) (float64, error) {
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

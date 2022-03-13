package feature

import "fmt"

type KanjiFeatureOrderCalculator struct {
	Manager KanjiFeatureManager
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

		mask, err := fc.Manager.OrderMask(fullNameLength, ci)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		index, err := fc.Manager.SelectFeatureOrderPosition(pieceOfName, i)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		v, err := fc.Manager.Get(Character(c)).GetOrderValue(index, mask)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}
		score += v
	}

	return score, nil
}

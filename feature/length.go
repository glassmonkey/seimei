package feature

import "fmt"

type KanjiLengthFeatureCalculator struct {
	Manager KanjiFeatureManager
}

func (fc KanjiLengthFeatureCalculator) Score(pieceOfName PartOfNameCharacters, fullNameLength int) (float64, error) {
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

		mask, err := fc.Manager.LengthMask(fullNameLength, ci)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		index, err := fc.Manager.SelectLengthFeaturePosition(pieceOfName)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		v, err := fc.Manager.Get(Character(c)).GetLengthValue(index, mask)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		score += v
	}

	return score, nil
}

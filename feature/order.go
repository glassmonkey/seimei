package feature

import (
	"fmt"
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

		mask, err := fc.Manager.Mask(fullNameLength, ci)
		if err != nil {
			return 0.0, fmt.Errorf("failed order score: %w", err)
		}

		index, err := fc.Manager.SelectFeaturePosition(pieceOfName, i)
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

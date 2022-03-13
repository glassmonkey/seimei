package feature

const (
	CharacterFeatureSize = 1
	OrderFeatureSize     = 6
	LengthFeatureSize    = 8
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

func DefaultKanjiFeature() KanjiFeature {
	return KanjiFeature{
		Character: "Default",
		Order:     []float64{0, 0, 0, 0, 0, 0},
		Length:    []float64{0, 0, 0, 0, 0, 0, 0, 0},
	}
}

type KanjiFeature struct {
	Character Character
	Order     []float64
	Length    []float64
}

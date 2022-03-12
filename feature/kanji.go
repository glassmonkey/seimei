package feature

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

type Character string

type KanjiFeatureManager struct {
	KanjiFeatureMap map[Character]KanjiFeature
}

func NewKanjiFeatureManager(assets string) KanjiFeatureManager {
	r := csv.NewReader(strings.NewReader(assets))
	m := make(map[Character]KanjiFeature)

	for i := 0; ; i++ {
		record, err := r.Read()
		//nolint:errorlint
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		if i == 0 {
			continue
		}

		c := Character(record[0])

		var os, ls []float64

		for _, o := range record[1 : 1+OrderFeatureSize] {
			//nolint:gomnd
			s, err := strconv.ParseFloat(o, 64)
			if err != nil {
				panic(err)
			}

			os = append(os, s)
		}

		for _, l := range record[7:15] {
			//nolint:gomnd
			s, err := strconv.ParseFloat(l, 64)
			if err != nil {
				panic(err)
			}

			ls = append(ls, s)
		}

		m[c] = KanjiFeature{
			Character: c,
			Order:     os,
			Length:    ls,
		}
	}

	return KanjiFeatureManager{
		KanjiFeatureMap: m,
	}
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

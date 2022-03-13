package seimei

import (
	// Using embed.
	_ "embed"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/glassmonkey/seimei/feature"
	"github.com/glassmonkey/seimei/parser"
)

//go:embed namedivider-python/assets/kanji.csv
var assets string

func InitNameParser(parseString string, manager feature.KanjiFeatureManager) parser.NameParser {
	return parser.NewNameParser(parser.Separator(parseString), manager)
}

func InitKanjiFeatureManager() feature.KanjiFeatureManager {
	r := csv.NewReader(strings.NewReader(assets))
	m := make(map[feature.Character]feature.KanjiFeature)

	for i := 0; ; i++ {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}

		if i == 0 {
			continue
		}

		c := feature.Character(record[0])

		var os, ls []float64

		for _, o := range record[feature.CharacterFeatureSize : feature.CharacterFeatureSize+feature.OrderFeatureSize] {
			s, err := strconv.ParseFloat(o, 64)
			if err != nil {
				panic(err)
			}

			os = append(os, s)
		}

		for _, l := range record[feature.CharacterFeatureSize+feature.OrderFeatureSize : feature.CharacterFeatureSize+feature.OrderFeatureSize+feature.LengthFeatureSize] {
			s, err := strconv.ParseFloat(l, 64)
			if err != nil {
				panic(err)
			}

			ls = append(ls, s)
		}

		kf, err := feature.NewKanjiFeature(c, os, ls)
		if err != nil {
			panic(err)
		}

		m[c] = kf
	}

	return feature.KanjiFeatureManager{
		KanjiFeatureMap: m,
	}
}

func Run(fullname string, parseString string) error {
	m := InitKanjiFeatureManager()
	p := InitNameParser(parseString, m)

	name, err := p.Parse(parser.FullName(fullname))
	if err != nil {
		return fmt.Errorf("happen error: %w", err)
	}

	fmt.Printf("%s\n", name.String())

	return nil
}

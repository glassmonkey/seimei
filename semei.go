package seimei

import (
	// Using embed.
	_ "embed"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/glassmonkey/seimei/feature"
	"github.com/glassmonkey/seimei/parser"
)

type (
	Name        string
	ParseString string
	Path        string
)

//go:embed namedivider-python/assets/kanji.csv
var assets string

func InitNameParser(parseString ParseString, manager feature.KanjiFeatureManager) parser.NameParser {
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

func InitReader(path Path) (*csv.Reader, error) {
	f, err := os.Open(string(path))
	if err != nil {
		return nil, fmt.Errorf("fatal error file load: %w", err)
	}

	return csv.NewReader(f), nil
}

func ParseName(out, stderr io.Writer, fullname Name, parseString ParseString) error {
	m := InitKanjiFeatureManager()
	p := InitNameParser(parseString, m)

	name, err := p.Parse(parser.FullName(fullname))
	if err != nil {
		_, err := fmt.Fprintf(stderr, "%s\n", err.Error())
		if err != nil {
			return fmt.Errorf("happen error write stderr: %w", err)
		}
		return nil
	}

	_, err = fmt.Fprintf(out, "%s\n", name.String())
	if err != nil {
		return fmt.Errorf("happen error write stdout: %w", err)
	}

	return nil
}

func ParseFile(out, stderr io.Writer, path Path, parseString ParseString) error {
	m := InitKanjiFeatureManager()
	p := InitNameParser(parseString, m)

	r, err := InitReader(path)
	if err != nil {
		return fmt.Errorf("happen error load file: %w", err)
	}

	for c := 1; ; c++ {
		record, err := r.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			fmt.Fprintf(stderr, "load line error on line %d: %v\n", c, err)
			continue
		}

		if len(record) != 1 {
			fmt.Fprintf(stderr, "format error on line %d: %v\n", c, record)
			continue
		}

		name, err := p.Parse(parser.FullName(record[0]))
		if err != nil {
			fmt.Fprintf(stderr, "parse error on line %d: %v\n", c, err)
			continue
		}

		fmt.Fprintf(out, "%s\n", name.String())
	}

	return nil
}

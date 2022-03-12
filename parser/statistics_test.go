package parser_test

import (
	"testing"

	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

func TestStatisticsParser_Parse(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name  string
		input parser.FullName
		want  parser.DividedName
		skip  bool
	}

	separator := parser.Separator("/")
	tests := []testdata{
		{
			name:  "3文字",
			input: "菅義偉",
			want: parser.DividedName{
				LastName:  "菅",
				FirstName: "義偉",
				Separator: separator,
				Score:     0.1111111111111111, // patch work score, todo fix.
				Algorithm: parser.Statistics,
			},
			skip: false,
		},
		{
			name:  "4文字",
			input: "阿部晋三",
			want: parser.DividedName{
				LastName:  "阿部",
				FirstName: "晋三",
				Separator: separator,
				Score:     1, // patch work score, todo fix.
				Algorithm: parser.Statistics,
			},
			skip: false,
		},
		{
			name:  "5文字",
			input: "中曽根康弘",
			want: parser.DividedName{
				LastName:  "中曽根",
				FirstName: "康弘",
				Separator: separator,
				Score:     0.1111111111111111, // patch work score, todo fix.
				Algorithm: parser.Statistics,
			},
			skip: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.skip {
				t.Skip()
			}
			sut := parser.StatisticsParser{
				Calculator: DummyKanjiFeatureCalculator{},
			}
			got, err := sut.Parse(tt.input, separator)
			if err != nil {
				t.Errorf("error is not nil, err=%v", err)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("divided name mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

type DummyKanjiFeatureCalculator struct{}

func (s DummyKanjiFeatureCalculator) Score(lastName parser.LastName, firstName parser.FirstName) float64 {
	v := float64(len(lastName) - len(firstName))
	if v == 0 {
		return 1
	}

	return 1 / (v * v)
}

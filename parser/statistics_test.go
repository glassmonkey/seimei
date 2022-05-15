package parser_test

import (
	"testing"

	"github.com/glassmonkey/seimei/v2"
	"github.com/glassmonkey/seimei/v2/parser"
	"github.com/google/go-cmp/cmp"
)

func TestStatisticsParser_Parse(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name  string
		input parser.FullName
		want  parser.DividedName
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
				Score:     0.48027055739279506,
				Algorithm: parser.Statistics,
			},
		},
		{
			name:  "4文字",
			input: "阿部晋三",
			want: parser.DividedName{
				LastName:  "阿部",
				FirstName: "晋三",
				Separator: separator,
				Score:     0.47397644480584417,
				Algorithm: parser.Statistics,
			},
		},
		{
			name:  "5文字",
			input: "中曽根康弘",
			want: parser.DividedName{
				LastName:  "中曽根",
				FirstName: "康弘",
				Separator: separator,
				Score:     0.3127240879300895,
				Algorithm: parser.Statistics,
			},
		},
		{
			name:  "すべてひらがな",
			input: "やまだはなこ",
			want: parser.DividedName{
				LastName:  "や",
				FirstName: "まだはなこ",
				Separator: separator,
				Score:     0.16666666666666666,
				Algorithm: parser.Statistics,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := parser.NewStatisticsParser(seimei.InitKanjiFeatureManager())
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

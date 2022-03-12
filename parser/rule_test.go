package parser_test

import (
	"testing"

	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

func TestRuleBaseParser_Parse(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name  string
		input parser.FullName
		want  parser.DividedName
	}

	separator := parser.Separator("/")
	tests := []testdata{
		{
			name:  "カタカナの名前の場合",
			input: "中山マサ",
			want: parser.DividedName{
				LastName:  "中山",
				FirstName: "マサ",
				Separator: separator,
				Score:     1,
				Algorithm: parser.Rule,
			},
		},
		{
			name:  "名字にカタカナを含む場合で名前がカタカナ",
			input: "関ヶ原マサ",
			want: parser.DividedName{
				LastName:  "関ヶ原",
				FirstName: "マサ",
				Separator: separator,
				Score:     1,
				Algorithm: parser.Rule,
			},
		},
		{
			name:  "名字にカタカナを含む名前が漢字の場合",
			input: "関ヶ原太郎",
			//nolint:exhaustivestruct
			want: parser.DividedName{},
		},
		{
			name:  "フルネームが漢字の場合",
			input: "中山太郎",
			//nolint:exhaustivestruct
			want: parser.DividedName{},
		},
		{
			name:  "名前がひらがなの場合",
			input: "平塚らいてう",
			want: parser.DividedName{
				LastName:  "平塚",
				FirstName: "らいてう",
				Separator: separator,
				Score:     1,
				Algorithm: parser.Rule,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := parser.NewRuleBaseParser()
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

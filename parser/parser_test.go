package parser_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

func TestNameParser_Parse(t *testing.T) {
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
			name:  "ルールベースの場合",
			input: "中山マサ",
			want: parser.DividedName{
				LastName:  "中山",
				FirstName: "マサ",
				Separator: separator,
				Score:     1,
				Algorithm: parser.Rule,
			},
			skip: false,
		},
		{
			name:  "フルネームが漢字の場合",
			input: "田中太郎",
			want: parser.DividedName{
				LastName:  "田中",
				FirstName: "太郎",
				Separator: "/",
				Score:     0,
				Algorithm: parser.Dummy,
			},
			skip: false,
		},
		{
			name:  "フルネームが漢字の場合",
			input: "竈門炭治郎",
			want: parser.DividedName{
				LastName:  "竈門",
				FirstName: "炭治郎",
				Separator: "/",
				Score:     0,
				Algorithm: parser.Dummy,
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
			sut := parser.NewNameParser(separator)
			got, err := sut.Parse(tt.input)
			if err != nil {
				t.Errorf("error is not nil, err=%v", err)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("divided name mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestNameParser_Parse_Validate(t *testing.T) {
	t.Parallel()

	sut := parser.NewNameParser("/")
	_, gotErr := sut.Parse("あ")
	wantErr := parser.ErrTextLength

	if !errors.Is(gotErr, wantErr) {
		t.Errorf("error is not expected, got error=(%v), want error=(%v)", gotErr, wantErr)
	}
}

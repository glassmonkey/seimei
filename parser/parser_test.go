package parser_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/seimei/v2"
	"github.com/glassmonkey/seimei/v2/feature"
	"github.com/glassmonkey/seimei/v2/parser"
	"github.com/google/go-cmp/cmp"
)

func TestNameParser_Parse(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name  string
		input parser.FullName
		want  parser.DividedName
	}

	separator := parser.Separator("/")
	tests := []testdata{
		{
			name:  "原敬",
			input: "原敬",
			want: parser.DividedName{
				LastName:  "原",
				FirstName: "敬",
				Separator: separator,
				Score:     1,
				Algorithm: parser.Rule,
			},
		},
		{
			name:  "菅義偉",
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
			name:  "阿部晋三",
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
			name:  "中曽根康弘",
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
			name:  "中山マサ",
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
			name:  "やまだはなこ",
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
			sut := parser.NewNameParser(separator, seimei.InitKanjiFeatureManager())
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
	//nolint:exhaustivestruct
	sut := parser.NewNameParser("/", feature.KanjiFeatureManager{})
	_, gotErr := sut.Parse("あ")
	wantErr := parser.ErrNameLength

	if !errors.Is(gotErr, wantErr) {
		t.Errorf("error is not expected, got error=(%v), want error=(%v)", gotErr, wantErr)
	}
}

func TestFullName_Length(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name  string
		input parser.FullName
		want  int
	}

	tests := []testdata{
		{
			name:  "漢字",
			input: "中山",
			want:  2,
		},
		{
			name:  "アルファベット混合",
			input: "DJ田中",
			want:  4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := tt.input
			got := sut.Length()

			if got != tt.want {
				t.Errorf("length is not expected, got=(%d), want=(%d)", got, tt.want)
			}
		})
	}
}

func TestFullName_Sprint(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name          string
		input         parser.FullName
		inputPosition int
		wantLastName  parser.LastName
		wantFirstName parser.FirstName
		wantErr       error
	}

	tests := []testdata{
		{
			name:          "0文字目",
			input:         "寿限無寿限無",
			inputPosition: 0,
			wantLastName:  "",
			wantFirstName: "寿限無寿限無",
			wantErr:       nil,
		},
		{
			name:          "4文字目",
			input:         "寿限無寿限無",
			inputPosition: 4,
			wantLastName:  "寿限無寿",
			wantFirstName: "限無",
			wantErr:       nil,
		},
		{
			name:          "6文字目",
			input:         "寿限無寿限無",
			inputPosition: 6,
			wantLastName:  "寿限無寿限無",
			wantFirstName: "",
			wantErr:       nil,
		},
		{
			name:          "7文字目は制限を超えるのでエラーになる",
			input:         "寿限無寿限無",
			inputPosition: 7,
			wantLastName:  "",
			wantFirstName: "",
			wantErr:       parser.ErrSplitPosition,
		},
		{
			name:          "-1文字目指定はエラーになる",
			input:         "寿限無寿限無",
			inputPosition: -1,
			wantLastName:  "",
			wantFirstName: "",
			wantErr:       parser.ErrSplitPosition,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := tt.input
			l, f, err := sut.Split(tt.inputPosition)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}
			if l != tt.wantLastName {
				t.Errorf("LastName is not expected, got=(%s), want=(%s)", l, tt.wantLastName)
			}
			if f != tt.wantFirstName {
				t.Errorf("LastName is not expected, got=(%s), want=(%s)", f, tt.wantFirstName)
			}
			got := parser.JoinName(l, f)
			if got != tt.input {
				t.Errorf("fullname is not expected, got=(%s), want=(%s)", got, tt.input)
			}
			if f.Length()+l.Length() != tt.input.Length() {
				t.Errorf("fullname's length is not expected, got(fist_name=(%s), last_name=(%s)), want(%s)", f, l, tt.input)
			}
		})
	}
}

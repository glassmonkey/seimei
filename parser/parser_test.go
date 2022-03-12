package parser_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

//nolint:godox
// TODO: Refactor parameterized test.
func TestNameParser_Parse(t *testing.T) {
	t.Parallel()

	sut := parser.NewNameParser("/")
	got, err := sut.Parse("田中太郎")
	want := parser.DividedName{
		LastName:  "田中",
		FirstName: "太郎",
		Separator: "/",
		Score:     0,
		Algorithm: parser.Dummy,
	}

	if err != nil {
		t.Errorf("error is not nil, err=%v", err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("divided name mismatch (-got +want):\n%s", diff)
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

func TestNameParser_Parse_SingleFirstNameAndSingleLastName(t *testing.T) {
	t.Parallel()

	sut := parser.NewNameParser("/")
	got, err := sut.Parse("乙一")
	want := parser.DividedName{
		LastName:  "乙",
		FirstName: "一",
		Separator: "/",
		Score:     1,
		Algorithm: parser.Rule,
	}

	if err != nil {
		t.Errorf("error is not nil, err=%v", err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("divided name mismatch (-got +want):\n%s", diff)
	}
}

func TestNameParser_Parse_NameHasNotKanjiName(t *testing.T) {
	t.Parallel()

	sut := parser.NewNameParser("/")
	got, err := sut.Parse("中山マサ")
	want := parser.DividedName{
		LastName:  "中山",
		FirstName: "マサ",
		Separator: "/",
		Score:     1,
		Algorithm: parser.Rule,
	}

	if err != nil {
		t.Errorf("error is not nil, err=%v", err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("divided name mismatch (-got +want):\n%s", diff)
	}
}

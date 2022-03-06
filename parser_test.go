package seimei_test

import (
	"errors"
	"github.com/glassmonkey/seimei"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNameParser_Parse(t *testing.T) {
	sut := seimei.NewNameParser("/")
	got, err := sut.Parse("田中太郎")
	want := seimei.DividedName{
		LastName: "田中",
		FirstName: "太郎",
		Separator: "/",
		Score: 0,
		Algorithm: seimei.Dummy,
	}
	if err != nil {
		t.Errorf("error is not nil, err=%v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("divided name mismatch (-got +want):\n%s", diff)
	}
}

func TestNameParser_Parse_Validate(t *testing.T) {
	sut := seimei.NewNameParser("/")
	_, err := sut.Parse("あ")
	wantErr := seimei.ErrTextLength
	if !errors.Is(err, wantErr) {
		t.Errorf("error is not expected, got error=(%v), want error=(%v)", err, wantErr)
	}
}

func TestNameParser_Parse_SingleFirstNameAndSingleLastName(t *testing.T) {
	sut := seimei.NewNameParser("/")
	got, err := sut.Parse("乙一")
	want := seimei.DividedName{
		LastName: "乙",
		FirstName: "一",
		Separator: "/",
		Score: 0,
		Algorithm: seimei.Rule,
	}
	if err != nil {
		t.Errorf("error is not nil, err=%v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("divided name mismatch (-got +want):\n%s", diff)
	}
}
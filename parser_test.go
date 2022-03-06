package seimei_test

import (
	"github.com/glassmonkey/seimei"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNameParser_Parse(t *testing.T) {
	sut := seimei.NameParser{}
	got, err := sut.Parse("田中太郎")
	want := seimei.DividedName{LastName: "田中", FirstName: "太郎"}
	if err != nil {
		t.Errorf("error is not nil, err=%v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("divided name mismatch (-got +want):\n%s", diff)
	}
}
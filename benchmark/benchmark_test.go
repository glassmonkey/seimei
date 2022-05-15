package benchmark_test

import (
	"bytes"
	_ "embed"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/glassmonkey/seimei/v2"
)

//go:embed sample.csv
// generate from https://testdata.userlocal.jp
var testdata string

func TestRunCompareOrigin(t *testing.T) {
	tests := strings.Split(testdata, "\n")
	total := len(tests)

	t.Parallel()
	// nolint: paralleltest
	for i, tt := range tests {
		tt := tt
		i := i

		t.Run(tt, func(t *testing.T) {
			t.Parallel()
			out := &bytes.Buffer{}
			input := strings.ReplaceAll(tt, " ", "")
			origin, err := exec.Command("nmdiv", "name", input).Output()
			if err != nil {
				t.Fatalf("happen error: %v", err)
			}
			err = seimei.Run(out, input, " ")
			if err != nil {
				t.Fatalf("happen error: %v", err)
			}
			if out.String() != string(origin) {
				t.Errorf("failed to test diff from origin. got: %s, origin: %s", out, origin)
			}
			t.Logf("ok: %s, %d/%d", tt, i, total)
		})
	}
}

func TestRunCompareAnswer(t *testing.T) {
	t.Parallel()

	tests := strings.Split(testdata, "\n")
	total := len(tests)
	// nolint: paralleltest
	for i, tt := range tests {
		tt := tt

		if len(tt) == 0 {
			continue
		}

		i := i

		t.Run(tt, func(t *testing.T) {
			t.Parallel()
			orig := tt
			out := &bytes.Buffer{}
			input := strings.ReplaceAll(orig, " ", "")
			want := fmt.Sprintf("%s\n", orig)
			err := seimei.Run(out, input, " ")
			if err != nil {
				t.Fatalf("happen error: %v", err)
			}
			if out.String() != want {
				t.Errorf("failed to test diff correct answer. got: %s, want: %s", out, want)
			}
			t.Logf("ok: %s, %d/%d", tt, i, total)
		})
	}
}

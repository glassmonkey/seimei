package seimei_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/glassmonkey/seimei"
)

func TestSetFlagForName(t *testing.T) {
	type testdata struct {
		name            string
		input           []string
		wantName        seimei.Name
		wantParseString seimei.ParseString
		wantErrMsg      string
	}

	tests := []testdata{
		{
			name:            "基本",
			input:           []string{"-name", "田中太郎"},
			wantName:        "田中太郎",
			wantParseString: " ",
		},
		{
			name:            "パース文字指定",
			input:           []string{"-name", "田中太郎", "-parse", "@"},
			wantName:        "田中太郎",
			wantParseString: "@",
		},
		{
			name:            "--2個でも良い",
			input:           []string{"--name", "田中太郎"},
			wantName:        "田中太郎",
			wantParseString: " ",
		},
		{
			name:       "指定がない",
			input:      []string{"-name"},
			wantErrMsg: "name command parse error: flag needs an argument: -name",
		},
		{
			name:       "未定義のパラメータ利用",
			input:      []string{"--name", "田中太郎", "-x"},
			wantErrMsg: "name command parse error: flag provided but not defined: -x",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotName, gotParseString, gotErr := seimei.SetFlagForName(tt.input)
			if tt.wantErrMsg == "" && gotErr != nil {
				t.Fatalf("happen error: %v", gotErr)
			}
			if tt.wantErrMsg != "" {
				if gotErr == nil {
					t.Fatal("happen no error")
				}
				if diff := cmp.Diff(tt.wantErrMsg, gotErr.Error()); diff != "" {
					t.Fatalf("failed to test on error. diff: %s", diff)
				}
			}

			if gotName != tt.wantName {
				t.Errorf("failed to test. got: %s, want: %s", gotName, tt.wantName)
			}
			if gotParseString != tt.wantParseString {
				t.Errorf("failed to test. got: %s, want: %s", gotParseString, tt.wantParseString)
			}
		})
	}
}

func TestSetFlagForFile(t *testing.T) {

	type testdata struct {
		name            string
		input           []string
		wantPath        seimei.Path
		wantParseString seimei.ParseString
		wantErrMsg      string
	}

	tests := []testdata{
		{
			name:            "基本",
			input:           []string{"-file", "/tmb/hoge.csv"},
			wantPath:        "/tmb/hoge.csv",
			wantParseString: " ",
		},
		{
			name:            "パース文字指定",
			input:           []string{"-file", "/tmb/hoge.csv", "-parse", "@"},
			wantPath:        "/tmb/hoge.csv",
			wantParseString: "@",
		},
		{
			name:            "--2個でも良い",
			input:           []string{"--file", "/tmb/hoge.csv"},
			wantPath:        "/tmb/hoge.csv",
			wantParseString: " ",
		},
		{
			name:       "指定がない",
			input:      []string{"-file"},
			wantErrMsg: "file command parse error: flag needs an argument: -file",
		},
		{
			name:       "未定義のパラメータ利用",
			input:      []string{"--file", "/tmb/hoge.csv", "-x"},
			wantErrMsg: "file command parse error: flag provided but not defined: -x",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotName, gotParseString, gotErr := seimei.SetFlagForFile(tt.input)
			if tt.wantErrMsg == "" && gotErr != nil {
				t.Fatalf("happen error: %v", gotErr)
			}
			if tt.wantErrMsg != "" {
				if gotErr == nil {
					t.Fatal("happen no error")
				}
				if diff := cmp.Diff(tt.wantErrMsg, gotErr.Error()); diff != "" {
					t.Fatalf("failed to test on error. diff: %s", diff)
				}
			}

			if gotName != tt.wantPath {
				t.Errorf("failed to test. got: %s, want: %s", gotName, tt.wantPath)
			}
			if gotParseString != tt.wantParseString {
				t.Errorf("failed to test. got: %s, want: %s", gotParseString, tt.wantParseString)
			}
		})
	}
}

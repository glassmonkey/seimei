package seimei_test

import (
	"bytes"
	"testing"

	"github.com/glassmonkey/seimei"
	"github.com/google/go-cmp/cmp"
)

func TestSetFlagForName(t *testing.T) {
	type testdata struct {
		name       string
		input      []string
		wantOut    string
		wantErrOut string
		wantErrMsg string
	}

	tests := []testdata{
		{
			name:    "基本",
			input:   []string{"--name", "田中太郎"},
			wantOut: "田中 太郎\n",
		},
		{
			name:    "パース文字指定",
			input:   []string{"--name", "田中太郎", "--parse", "@"},
			wantOut: "田中@太郎\n",
		},
		{
			name:    "短縮指定でも良い",
			input:   []string{"-n", "田中太郎"},
			wantOut: "田中 太郎\n",
		},
		{
			name:       "指定がない",
			input:      []string{"--name"},
			wantErrMsg: "flag needs an argument: --name",
		},
		{
			name:       "未定義のパラメータ利用",
			input:      []string{"--name", "田中太郎", "-x"},
			wantErrMsg: "unknown shorthand flag: 'x' in -x",
		},
		{
			name:       "未定義のパラメータ利用",
			input:      []string{"--name", "田中太郎", "--any"},
			wantErrMsg: "unknown flag: --any",
		},
		{
			name:       "空",
			input:      []string{},
			wantErrMsg: "required flag(s) \"name\" not set",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			sut := seimei.BuildNameCmd()
			sut.SetOut(stdout)
			sut.SetErr(stderr)
			sut.SetArgs(tt.input)

			_, gotErr := sut.ExecuteC()
			if tt.wantErrMsg == "" && gotErr != nil {
				t.Fatalf("happen error: %v", gotErr)
			}
			if tt.wantErrMsg != "" {
				if gotErr == nil {
					t.Fatal("happen no error")
				}
				if diff := cmp.Diff(gotErr.Error(), tt.wantErrMsg); diff != "" {
					t.Fatalf("failed to test on error. diff: %s", diff)
				}
				return
			}
			if diff := cmp.Diff(stdout.String(), tt.wantOut); diff != "" {
				t.Errorf("failed to test on error. diff: %s", diff)
			}
			if diff := cmp.Diff(stderr.String(), tt.wantErrOut); diff != "" {
				t.Errorf("failed to test on error. diff: %s", diff)
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
		{
			name:       "空",
			input:      []string{},
			wantErrMsg: "-path must be required (ex. /tmp/hoge.csv)",
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

func TestRun(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name       string
		input      []string
		want       string
		wantErrMsg string
	}

	tests := []testdata{
		{
			name:  "名前指定",
			input: []string{"seimei", "name", "-name", "田中太郎"},
			want:  "田中 太郎\n",
		},
		{
			name:  "名前指定実行のヘルプ",
			input: []string{"seimei", "name", "-h"},
			want:  ``,
		},
		{
			name:  "ファイル経由の実行",
			input: []string{"seimei", "file", "-file", "testdata/success.csv"},
			want: `田中 太郎
乙 一
竈門 炭治郎
中曽根 康弘
`,
		},
		{
			name:  "ファイル経由の実行のヘルプ",
			input: []string{"seimei", "file", "-h"},
			want:  ``,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}

			if err := seimei.Run(tt.input, stdout, stderr); err != nil {
				t.Fatalf("happen error: %v", err)
			}

			if stdout.String() != tt.want {
				t.Errorf("failed to test. got: %s, want: %s", stdout, tt.want)
			}
			if stderr.String() != tt.wantErrMsg {
				t.Errorf("failed to test. got: %s, want: %s", stderr, tt.wantErrMsg)
			}
		})
	}
}

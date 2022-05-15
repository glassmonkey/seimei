package seimei_test

import (
	"bytes"
	"testing"

	"github.com/glassmonkey/seimei"
	"github.com/google/go-cmp/cmp"
)

func TestBuildNameCmd(t *testing.T) {
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
			input:   []string{"-n", "田中太郎", "-p", "/"},
			wantOut: "田中/太郎\n",
		},
		{
			name:       "指定がない",
			input:      []string{"--name"},
			wantErrMsg: "flag needs an argument: --name",
		},
		{
			name:       "未定義の短縮パラメータ利用",
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

func TestBuildFileCmd(t *testing.T) {

	type testdata struct {
		name       string
		input      []string
		wantOut    string
		wantErrOut string
		wantErrMsg string
	}

	tests := []testdata{
		{
			name:  "基本",
			input: []string{"--file", "./testdata/success.csv"},
			wantOut: `田中 太郎
乙 一
竈門 炭治郎
中曽根 康弘
`,
		},
		{
			name:  "パース文字指定",
			input: []string{"--file", "./testdata/success.csv", "--parse", "@"},
			wantOut: `田中@太郎
乙@一
竈門@炭治郎
中曽根@康弘
`,
		},
		{
			name:  "短縮指定でも良い",
			input: []string{"-f", "./testdata/success.csv"},
			wantOut: `田中 太郎
乙 一
竈門 炭治郎
中曽根 康弘
`,
		},
		{
			name:  "実行時エラーが混ざる場合",
			input: []string{"-f", "./testdata/part_of_error.csv"},
			wantOut: `田中 太郎
竈門 炭治郎
中曽根 康弘
`,
			wantErrOut: `parse error on line 2: parse error: name length needs at least 2 chars
`,
		},
		{
			name:       "指定がない",
			input:      []string{"--file"},
			wantErrMsg: "flag needs an argument: --file",
		},
		{
			name:       "未定義の短縮パラメータ利用",
			input:      []string{"--file", "/tmb/hoge.csv", "-x"},
			wantErrMsg: "unknown shorthand flag: 'x' in -x",
		},
		{
			name:       "未定義のパラメータ利用",
			input:      []string{"--file", "/tmb/hoge.csv", "--any"},
			wantErrMsg: "unknown flag: --any",
		},
		{
			name:       "空",
			input:      []string{},
			wantErrMsg: "required flag(s) \"file\" not set",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			sut := seimei.BuildFileCmd()
			sut.SetOut(stdout)
			sut.SetErr(stderr)
			sut.SetArgs(tt.input)

			gotErr := sut.Execute()
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

func TestRun(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name       string
		input      []string
		wantOut    string
		wantErrOut string
	}

	tests := []testdata{
		{
			name:    "名前指定",
			input:   []string{"name", "--name", "田中太郎"},
			wantOut: "田中 太郎\n",
		},
		{
			name:  "名前指定実行のヘルプ",
			input: []string{"name", "-h"},
			wantOut: `It parse single full name.
Provide the full name to be parsed with the required flag (--name).

Usage:
  seimei name [flags]

Examples:
seimei name --name 田中太郎

Flags:
  -n, --name string    田中太郎
  -p, --parse string     (default " ")
  -h, --help           help for name
`,
		},
		{
			name:  "ファイル経由の実行",
			input: []string{"file", "--file", "testdata/success.csv"},
			wantOut: `田中 太郎
乙 一
竈門 炭治郎
中曽根 康弘
`,
		},
		{
			name:  "ファイル経由の実行のヘルプ",
			input: []string{"file", "-h"},
			wantOut: `It bulk parse full name lit in the file.
Provide the file path with full name list to the required flag (--file).

Usage:
  seimei file [flags]

Examples:
seimei file --file /path/to/dir/foo.csv

Flags:
  -f, --file string    /path/to/dir/foo.csv
  -p, --parse string     (default " ")
  -h, --help           help for file
`,
		},
		{
			name:  "サブコマンドなしはヘルプが表示される",
			input: []string{},
			wantOut: `Usage:
  seimei [flags]
  seimei [command]

Available Commands:
  name        It parse single full name.
  file        It bulk parse full name lit in the file.
  help        Help about any command

Flags:
  -h, --help   help for seimei

Use "seimei [command] --help" for more information about a command.
`,
		},
		{
			name:  "ヘルプが表示される",
			input: []string{"-h"},
			wantOut: `Usage:
  seimei [flags]
  seimei [command]

Available Commands:
  name        It parse single full name.
  file        It bulk parse full name lit in the file.
  help        Help about any command

Flags:
  -h, --help   help for seimei

Use "seimei [command] --help" for more information about a command.
`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			sut := seimei.BuildMainCmd()
			sut.SetOut(stdout)
			sut.SetErr(stderr)
			sut.SetArgs(tt.input)

			if err := sut.Execute(); err != nil {
				t.Fatalf("happen error: %v", err)
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

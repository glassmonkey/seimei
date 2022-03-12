package seimei_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/glassmonkey/seimei"
)

//nolint:tparallel
func TestRun(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name        string
		inputName   string
		inputParser string
		want        string
		skip        bool
	}

	tests := []testdata{
		{
			name:        "サンプル",
			inputName:   "田中太郎",
			inputParser: " ",
			want:        "田中 太郎",
			skip:        false,
		},
		{
			name:        "分割文字列が反映される",
			inputName:   "田中太郎",
			inputParser: "/",
			want:        "田中/太郎",
			skip:        false,
		},
		{
			name:        "ルールベースで動作する",
			inputName:   "乙一",
			inputParser: " ",
			want:        "乙 一",
			skip:        false,
		},
		{
			name:        "統計量ベースで動作する(仮実装)",
			inputName:   "竈門炭治郎",
			inputParser: " ",
			want:        "竈門 炭治郎",
			skip:        false,
		},
		{
			name:        "統計量ベースで動作しない(仮実装)",
			inputName:   "中曽根康弘",
			inputParser: " ",

			skip: true,
		},
	}
	//nolint:paralleltest
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}
			got := extractStdout(t, func() error {
				if err := seimei.Run(tt.inputName, tt.inputParser); err != nil {
					return fmt.Errorf("happen error: %w", err)
				}

				return nil
			})

			if got != tt.want {
				t.Errorf("failed to test. got: %s, want: %s", got, tt.want)
			}
		})
	}
}

// extractStdout
// refer: https://zenn.dev/glassonion1/articles/8ac939208bd455
func extractStdout(t *testing.T, runner func() error) string {
	t.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}()
	// パイプの作成(r: Reader, w: Writer)
	reader, writer, _ := os.Pipe()
	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = writer
	// テスト対象の関数を実行する

	if err := runner(); err != nil {
		t.Fatalf("failed to run: %v", err)
	}
	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	writer.Close()
	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(reader); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	// 文字列を取得する
	return strings.TrimRight(buf.String(), "\n")
}

package seimei_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/glassmonkey/seimei"
	"github.com/glassmonkey/seimei/feature"
	"github.com/google/go-cmp/cmp"
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
			want:        "中曽根 康弘",
			skip:        true,
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

func TestInitKanjiFeatureManager(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name        string
		inputKanji  feature.Character
		wantFeature feature.KanjiFeature
	}

	tests := []testdata{
		{
			name:        "デフォルト",
			inputKanji:  "無",
			wantFeature: feature.DefaultKanjiFeature(),
		},
		{
			name:       "csvの最初",
			inputKanji: "々",
			wantFeature: feature.KanjiFeature{
				Character: "々",
				Order: []float64{
					0, 275, 9, 0, 14, 25,
				},
				Length: []float64{
					0, 7, 276, 1, 0, 23, 16, 0,
				},
			},
		},
		{
			name:       "csvの最後",
			inputKanji: "葵",
			wantFeature: feature.KanjiFeature{
				Character: "葵",
				Order: []float64{
					1, 0, 0, 0, 0, 9,
				},
				Length: []float64{
					0, 1, 0, 0, 6, 3, 0, 0,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := seimei.InitKanjiFeatureManager()
			got := sut.Get(tt.inputKanji)

			if diff := cmp.Diff(got, tt.wantFeature); diff != "" {
				t.Errorf("feature value mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

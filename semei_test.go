package seimei_test

import (
	"bytes"
	"testing"

	"github.com/glassmonkey/seimei"
	"github.com/glassmonkey/seimei/feature"
	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name        string
		inputName   seimei.Name
		inputParser seimei.ParseString
		want        string
	}

	tests := []testdata{
		{
			name:        "サンプル",
			inputName:   "田中太郎",
			inputParser: " ",
			want:        "田中 太郎\n",
		},
		{
			name:        "分割文字列が反映される",
			inputName:   "田中太郎",
			inputParser: "/",
			want:        "田中/太郎\n",
		},
		{
			name:        "ルールベースで動作する",
			inputName:   "乙一",
			inputParser: " ",
			want:        "乙 一\n",
		},
		{
			name:        "統計量ベースで動作する",
			inputName:   "竈門炭治郎",
			inputParser: " ",
			want:        "竈門 炭治郎\n",
		},
		{
			name:        "統計量ベースで分割できる",
			inputName:   "中曽根康弘",
			inputParser: " ",
			want:        "中曽根 康弘\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := &bytes.Buffer{}

			if err := seimei.ParseName(out, tt.inputName, tt.inputParser); err != nil {
				t.Fatalf("happen error: %v", err)
			}

			if out.String() != tt.want {
				t.Errorf("failed to test. got: %s, want: %s", out, tt.want)
			}
		})
	}
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

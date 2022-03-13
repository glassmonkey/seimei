package feature_test

import (
	"testing"

	"github.com/glassmonkey/seimei/feature"
	"github.com/google/go-cmp/cmp"
)

func TestInitKanjiFeatureManager(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name        string
		inputKanji  feature.Character
		wantFeature feature.KanjiFeature
	}

	testMap := map[feature.Character]feature.KanjiFeature{
		"有": {
			Character: "有",
			Order: []float64{
				1, 0, 0, 0, 0, 9,
			},
			Length: []float64{
				0, 1, 0, 0, 6, 3, 0, 0,
			},
		},
	}

	tests := []testdata{
		{
			name:        "マッピングデータにいない場合はデフォルト",
			inputKanji:  "無",
			wantFeature: feature.DefaultKanjiFeature(),
		},
		{
			name:       "マッピングデータにいない場合はそれが取得される",
			inputKanji: "有",
			wantFeature: feature.KanjiFeature{
				Character: "有",
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
			sut := feature.KanjiFeatureManager{
				KanjiFeatureMap: testMap,
			}
			got := sut.Get(tt.inputKanji)

			if diff := cmp.Diff(got, tt.wantFeature); diff != "" {
				t.Errorf("feature value mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

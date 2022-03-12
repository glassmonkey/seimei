package feature_test

import (
	"testing"

	"github.com/glassmonkey/seimei"

	"github.com/glassmonkey/seimei/feature"
	"github.com/google/go-cmp/cmp"
)

func TestKanjiFeatureManager_Get(t *testing.T) {
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
			sut := feature.NewKanjiFeatureManager(seimei.Assets)
			got := sut.Get(tt.inputKanji)

			if diff := cmp.Diff(got, tt.wantFeature); diff != "" {
				t.Errorf("feature value mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

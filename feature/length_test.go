package feature_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/seimei/feature"
	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

func TestKanjiLengthFeatureCalculator_Score(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name                string
		inputName           feature.PartOfNameCharacters
		inputFullNameLength int
		wantSrore           float64
		wantErr             error
	}

	tests := []testdata{
		{
			name:                "名字",
			inputName:           parser.FirstName("冬馬"),
			inputFullNameLength: 5,
			wantSrore:           1.0 / 4,
			wantErr:             nil,
		},
		{
			name:                "名前",
			inputName:           parser.LastName("天ケ瀬"),
			inputFullNameLength: 5,
			wantSrore:           1.0 / 2,
			wantErr:             nil,
		},
		{
			name:                "フルネームと同じサイズ指定の場合はスコアは0",
			inputName:           parser.LastName("天ケ瀬"),
			inputFullNameLength: 3,
			wantSrore:           0,
			wantErr:             nil,
		},
		{
			name:                "指定文字列がフルネームより大きい場合はマスクデータの作成でエラーになる",
			inputName:           parser.LastName("天ケ瀬"),
			inputFullNameLength: 2,
			wantSrore:           0,
			wantErr:             feature.ErrOutRangeOrderMask,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := feature.KanjiLengthFeatureCalculator{
				Manager: stubKanjiManagerForLengthFeature(),
			}
			got, err := sut.Score(tt.inputName, tt.inputFullNameLength)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			if diff := cmp.Diff(got, tt.wantSrore); diff != "" {
				t.Errorf("score value mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func stubKanjiManagerForLengthFeature() feature.KanjiFeatureManager {
	o := feature.Features{1, 1, 1, 1, 1, 1}
	l := feature.Features{1, 1, 1, 1, 1, 1, 1, 1}

	return feature.KanjiFeatureManager{
		KanjiFeatureMap: map[feature.Character]feature.KanjiFeature{
			"冬": {Character: "冬", Order: o, Length: l},
			"馬": {Character: "馬", Order: o, Length: l},
			"天": {Character: "天", Order: o, Length: l},
			"ケ": {Character: "ケ", Order: o, Length: l},
			"瀬": {Character: "瀬", Order: o, Length: l},
		},
	}
}

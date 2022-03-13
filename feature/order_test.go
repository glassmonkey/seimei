package feature_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/seimei"

	"github.com/glassmonkey/seimei/feature"
	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

func TestKanjiOrderFeatureCalculator_ScoreWithStub(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name                string
		inputName           feature.PartOfNameCharacters
		inputFullNameLength int
		wantScore           float64
		wantErr             error
	}

	tests := []testdata{
		{
			name:                "名字",
			inputName:           parser.FirstName("冬馬"),
			inputFullNameLength: 5,
			wantScore:           1.0 / 3, // 1/4
			wantErr:             nil,
		},
		{
			name:                "名前",
			inputName:           parser.LastName("天ケ瀬"),
			inputFullNameLength: 5,
			wantScore:           0.5833333333333333, // 1/4 + 1/3
			wantErr:             nil,
		},
		{
			name:                "フルネームと同じサイズ指定の場合はスコアは0",
			inputName:           parser.LastName("天ケ瀬"),
			inputFullNameLength: 3,
			wantScore:           0,
			wantErr:             nil,
		},
		{
			name:                "指定文字列がフルネームより大きい場合はマスクデータの作成でエラーになる",
			inputName:           parser.LastName("天ケ瀬"),
			inputFullNameLength: 2,
			wantScore:           0,
			wantErr:             feature.ErrOutRangeOrderMask,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := feature.KanjiOrderFeatureCalculator{
				Manager: stubKanjiManagerForOrderFeature(),
			}
			got, err := sut.Score(tt.inputName, tt.inputFullNameLength)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			if diff := cmp.Diff(got, tt.wantScore); diff != "" {
				t.Errorf("score value mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestKanjiOrderFeatureCalculator_ScoreWithCSVData(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name                string
		inputName           feature.PartOfNameCharacters
		inputFullNameLength int
		wantScore           float64
		wantErr             error
	}

	tests := []testdata{
		{
			name:                "新海誠(名前)",
			inputName:           parser.FirstName("誠"),
			inputFullNameLength: 3,
			wantScore:           0,
			wantErr:             nil,
		},
		{
			name:                "新海誠(名字)",
			inputName:           parser.LastName("新海"),
			inputFullNameLength: 3,
			wantScore:           0.8305084745762712,
			wantErr:             nil,
		},
		{
			name:                "清武弘嗣(名前)",
			inputName:           parser.FirstName("弘嗣"),
			inputFullNameLength: 4,
			wantScore:           0.9919571045576407,
			wantErr:             nil,
		},
		{
			name:                "清武弘嗣(名字)",
			inputName:           parser.LastName("清武"),
			inputFullNameLength: 4,
			wantScore:           0.2222222222222222,
			wantErr:             nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := feature.KanjiOrderFeatureCalculator{
				Manager: seimei.InitKanjiFeatureManager(),
			}
			got, err := sut.Score(tt.inputName, tt.inputFullNameLength)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			if diff := cmp.Diff(got, tt.wantScore); diff != "" {
				t.Errorf("score value mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func stubKanjiManagerForOrderFeature() feature.KanjiFeatureManager {
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

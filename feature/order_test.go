package feature_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/seimei/feature"
	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

func TestKanjiFeatureOrderCalculator_Mask(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name          string
		inputLength   int
		inputPosition int
		wantMask      feature.Features
		wantErr       error
	}

	tests := []testdata{
		{
			name:          "-1文字目指定",
			inputLength:   5,
			inputPosition: -1,
			wantMask:      feature.Features{},
			wantErr:       feature.ErrOutRangeOrderMask,
		},
		{
			name:          "(1/5)文字目指定",
			inputLength:   5,
			inputPosition: 0,
			wantMask:      feature.Features{},
			wantErr:       feature.ErrInvalidOrderMask,
		},
		{
			name:          "(2/5)文字目指定",
			inputLength:   5,
			inputPosition: 1,
			wantMask: feature.Features{
				0, 1, 1, 1, 0, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(3/5)文字目指定",
			inputLength:   5,
			inputPosition: 2,
			wantMask: feature.Features{
				0, 1, 1, 1, 1, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(4/5)文字目指定",
			inputLength:   5,
			inputPosition: 3,
			wantMask: feature.Features{
				0, 0, 1, 1, 1, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(5/5)文字目指定",
			inputLength:   5,
			inputPosition: 4,
			wantMask:      feature.Features{},
			wantErr:       feature.ErrInvalidOrderMask,
		},
		{
			name:          "5文字目指定",
			inputLength:   5,
			inputPosition: 5,
			wantMask: feature.Features{
				0, 0, 1, 1, 1, 0,
			},
			wantErr: feature.ErrOutRangeOrderMask,
		},
		{
			name:          "(1/3)文字目指定",
			inputLength:   3,
			inputPosition: 0,
			wantMask:      feature.Features{},
			wantErr:       feature.ErrInvalidOrderMask,
		},
		{
			name:          "(2/3)文字目指定",
			inputLength:   3,
			inputPosition: 1,
			wantMask: feature.Features{
				0, 0, 1, 1, 0, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(3/3)文字目指定",
			inputLength:   3,
			inputPosition: 2,
			wantMask:      feature.Features{},
			wantErr:       feature.ErrInvalidOrderMask,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sut := feature.KanjiFeatureOrderCalculator{
				Manager: stubKanjiManagerForOrderFeature(),
			}
			got, err := sut.Mask(tt.inputLength, tt.inputPosition)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			if diff := cmp.Diff(got, tt.wantMask); diff != "" {
				t.Errorf("mask value mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestKanjiFeatureOrderCalculator_SelectFeaturePosition(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name          string
		inputName     feature.PartOfNameCharacters
		inputPosition int
		wantPosition  feature.OrderFeatureIndexPosition
		wantErr       error
	}

	tests := []testdata{
		{
			name:          "名前1文字目",
			inputName:     parser.FirstName("あきら"),
			inputPosition: 0,
			wantPosition:  feature.OrderFirstFeatureIndex.MoveFirstNameIndex(),
			wantErr:       nil,
		},
		{
			name:          "名前2文字目",
			inputName:     parser.FirstName("あきら"),
			inputPosition: 1,
			wantPosition:  feature.OrderMiddleFeatureIndex.MoveFirstNameIndex(),
			wantErr:       nil,
		},
		{
			name:          "名前3文字目",
			inputName:     parser.FirstName("あきら"),
			inputPosition: 2,
			wantPosition:  feature.OrderEndFeatureIndex.MoveFirstNameIndex(),
			wantErr:       nil,
		},
		{
			name:          "名字1文字目",
			inputName:     parser.LastName("中山田"),
			inputPosition: 0,
			wantPosition:  feature.OrderFirstFeatureIndex,
			wantErr:       nil,
		},
		{
			name:          "名字2文字目",
			inputName:     parser.LastName("中山田"),
			inputPosition: 1,
			wantPosition:  feature.OrderMiddleFeatureIndex,
			wantErr:       nil,
		},
		{
			name:          "名前3文字目",
			inputName:     parser.LastName("中山田"),
			inputPosition: 2,
			wantPosition:  feature.OrderEndFeatureIndex,
			wantErr:       nil,
		},
		{
			name:          "負の数を指定",
			inputName:     parser.LastName("中山田"),
			inputPosition: -1,
			wantPosition:  feature.OrderEndFeatureIndex,
			wantErr:       feature.ErrOutRangeFeatureIndex,
		},
		{
			name:          "最大を超える",
			inputName:     parser.LastName("中山田"),
			inputPosition: 3,
			wantPosition:  feature.OrderEndFeatureIndex,
			wantErr:       feature.ErrOutRangeFeatureIndex,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := feature.KanjiFeatureOrderCalculator{stubKanjiManagerForOrderFeature()}
			got, err := sut.SelectFeaturePosition(tt.inputName, tt.inputPosition)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			if diff := cmp.Diff(got, tt.wantPosition); diff != "" {
				t.Errorf("mask value mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestKanjiFeatureOrderCalculator_Score(t *testing.T) {
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
			wantSrore:           0.3333333333333333, // 1/4
			wantErr:             nil,
		},
		{
			name:                "名前",
			inputName:           parser.LastName("天ケ瀬"),
			inputFullNameLength: 5,
			wantSrore:           0.5833333333333333, // 1/4 + 1/3
			wantErr:             nil,
		},
		{
			name:                "名前",
			inputName:           parser.LastName("天ケ瀬"),
			inputFullNameLength: 5,
			wantSrore:           0.5833333333333333, // 1/4 + 1/3
			wantErr:             nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := feature.KanjiFeatureOrderCalculator{
				Manager: stubKanjiManagerForOrderFeature(),
			}
			got, err := sut.Score(tt.inputName, tt.inputFullNameLength)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
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

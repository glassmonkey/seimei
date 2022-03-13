package feature_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/seimei/feature"
	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

func TestKanjiOrderFeatureCalculator_OrderMask(t *testing.T) {
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

			sut := feature.KanjiFeatureManager{
				KanjiFeatureMap: map[feature.Character]feature.KanjiFeature{},
			}
			got, err := sut.OrderMask(tt.inputLength, tt.inputPosition)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
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

func TestKanjiOrderFeatureCalculator_LengthMask(t *testing.T) {
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
			wantMask: feature.Features{
				1, 1, 1, 1, 0, 0, 0, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(2/5)文字目指定",
			inputLength:   5,
			inputPosition: 1,
			wantMask: feature.Features{
				0, 1, 1, 1, 0, 0, 0, 1,
			},
			wantErr: nil,
		},
		{
			name:          "(3/5)文字目指定",
			inputLength:   5,
			inputPosition: 2,
			wantMask: feature.Features{
				0, 0, 1, 1, 0, 0, 1, 1,
			},
			wantErr: nil,
		},
		{
			name:          "(4/5)文字目指定",
			inputLength:   5,
			inputPosition: 3,
			wantMask: feature.Features{
				0, 0, 0, 1, 0, 1, 1, 1,
			},
			wantErr: nil,
		},
		{
			name:          "(5/5)文字目指定",
			inputLength:   5,
			inputPosition: 4,
			wantMask: feature.Features{
				0, 0, 0, 0, 1, 1, 1, 1,
			},
			wantErr: nil,
		},
		{
			name:          "5文字目指定",
			inputLength:   5,
			inputPosition: 5,
			wantMask:      feature.Features{},
			wantErr:       feature.ErrOutRangeOrderMask,
		},
		{
			name:          "(1/3)文字目指定",
			inputLength:   3,
			inputPosition: 0,
			wantMask: feature.Features{
				1, 1, 0, 0, 0, 0, 0, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(2/3)文字目指定",
			inputLength:   3,
			inputPosition: 1,
			wantMask: feature.Features{
				0, 1, 0, 0, 0, 1, 0, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(3/3)文字目指定",
			inputLength:   3,
			inputPosition: 2,
			wantMask: feature.Features{
				0, 0, 0, 0, 1, 1, 0, 0,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sut := feature.KanjiFeatureManager{
				KanjiFeatureMap: map[feature.Character]feature.KanjiFeature{},
			}
			got, err := sut.LengthMask(tt.inputLength, tt.inputPosition)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
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

func TestKanjiOrderFeatureCalculator_SelectOrderFeaturePosition(t *testing.T) {
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
			sut := feature.KanjiFeatureManager{
				KanjiFeatureMap: map[feature.Character]feature.KanjiFeature{},
			}
			got, err := sut.SelectOrderFeaturePosition(tt.inputName, tt.inputPosition)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
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

func TestKanjiOrderFeatureCalculator_SelectLengthFeaturePosition(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name         string
		inputName    feature.PartOfNameCharacters
		wantPosition feature.LengthFeatureIndexPosition
		wantErr      error
	}

	tests := []testdata{
		{
			name:         "名前指定",
			inputName:    parser.FirstName("あきら"),
			wantPosition: feature.LengthFeatureIndexPosition(6),
			wantErr:      nil,
		},
		{
			name:         "名字指定",
			inputName:    parser.LastName("中山田"),
			wantPosition: feature.LengthFeatureIndexPosition(2),
			wantErr:      nil,
		},
		{
			name:         "指定文字数が大きい場合の名字",
			inputName:    parser.LastName("寿限無寿限"),
			wantPosition: feature.LengthFeatureIndexPosition(3),
			wantErr:      nil,
		},
		{
			name:         "指定文字数が大きい場合の指名",
			inputName:    parser.FirstName("寿限無寿限"),
			wantPosition: feature.LengthFeatureIndexPosition(7),
			wantErr:      nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := feature.KanjiFeatureManager{
				KanjiFeatureMap: map[feature.Character]feature.KanjiFeature{},
			}
			got, err := sut.SelectLengthFeaturePosition(tt.inputName)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error is not expected, got error=(%v), want error=(%v)", err, tt.wantErr)
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

func TestKanjiFeature_GetOrderValue(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name          string
		inputFeature  feature.Features
		inputPosition feature.OrderFeatureIndexPosition
		inputMask     feature.Features
		wantScore     float64
		wantErr       error
	}

	tests := []testdata{
		{
			name:          "mask is all 1",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32, 64, 128},
			inputMask:     feature.Features{1, 1, 1, 1, 1, 1, 1, 1},
			inputPosition: 0,
			wantScore:     1.0 / (1 + 2 + 4 + 8 + 16 + 32 + 64 + 128),
			wantErr:       nil,
		},
		{
			name:          "mask is all 0",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32, 64, 128},
			inputMask:     feature.Features{0, 0, 0, 0, 0, 0, 0, 0},
			inputPosition: 0,
			wantScore:     0,
			wantErr:       nil,
		},
		{
			name:          "mask is half 1",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32, 64, 128},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0, 1, 0},
			inputPosition: 0,
			wantScore:     1.0 / (1 + 4 + 16 + 64),
			wantErr:       nil,
		},
		{
			name:          "mask is half 1 and target index 1",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32, 64, 128},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0, 1, 0},
			inputPosition: 1,
			wantScore:     0,
			wantErr:       nil,
		},
		{
			name:          "mask is half 1 and target index 2",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32, 64, 128},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0, 1, 0},
			inputPosition: 2,
			wantScore:     4.0 / (1 + 4 + 16 + 64),
			wantErr:       nil,
		},
		{
			name:          "target index -1",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32, 64, 128},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0, 1, 0},
			inputPosition: -1,
			wantScore:     0,
			wantErr:       feature.ErrOutRangeFeatureIndex,
		},
		{
			name:          "target index 6",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32, 64, 128},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0, 1, 0},
			inputPosition: 8,
			wantScore:     0,
			wantErr:       feature.ErrOutRangeFeatureIndex,
		},
		{
			name:          "input mask is un match size",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32, 64, 128},
			inputMask:     feature.Features{1, 0, 1, 0, 1},
			inputPosition: 0,
			wantScore:     0,
			wantErr:       feature.ErrInvalidFeatureSize,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			//nolint:exhaustivestruct
			sut := featureFixtures(inputForFixtures{
				orders: tt.inputFeature,
			})
			got, err := sut.GetOrderValue(tt.inputPosition, tt.inputMask)
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

func TestKanjiFeature_GetLengthValue(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name          string
		inputFeature  feature.Features
		inputPosition feature.LengthFeatureIndexPosition
		inputMask     feature.Features
		wantScore     float64
		wantErr       error
	}

	tests := []testdata{
		{
			name:          "mask is all 1",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32},
			inputMask:     feature.Features{1, 1, 1, 1, 1, 1},
			inputPosition: 0,
			wantScore:     1.0 / 63,
			wantErr:       nil,
		},
		{
			name:          "mask is all 0",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32},
			inputMask:     feature.Features{0, 0, 0, 0, 0, 0},
			inputPosition: 0,
			wantScore:     0,
			wantErr:       nil,
		},
		{
			name:          "mask is half 1",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0},
			inputPosition: 0,
			wantScore:     1.0 / (1 + 4 + 16),
			wantErr:       nil,
		},
		{
			name:          "mask is half 1 and target index 1",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0},
			inputPosition: 1,
			wantScore:     0,
			wantErr:       nil,
		},
		{
			name:          "mask is half 1 and target index 2",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0},
			inputPosition: 2,
			wantScore:     4.0 / (1 + 4 + 16),
			wantErr:       nil,
		},
		{
			name:          "target index -1",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0},
			inputPosition: -1,
			wantScore:     0,
			wantErr:       feature.ErrOutRangeFeatureIndex,
		},
		{
			name:          "target index 6",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32},
			inputMask:     feature.Features{1, 0, 1, 0, 1, 0},
			inputPosition: 6,
			wantScore:     0,
			wantErr:       feature.ErrOutRangeFeatureIndex,
		},
		{
			name:          "input mask is un match size",
			inputFeature:  feature.Features{1, 2, 4, 8, 16, 32},
			inputMask:     feature.Features{1, 0, 1, 0, 1},
			inputPosition: 0,
			wantScore:     0,
			wantErr:       feature.ErrInvalidFeatureSize,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			//nolint:exhaustivestruct
			sut := featureFixtures(inputForFixtures{
				length: tt.inputFeature,
			})
			got, err := sut.GetLengthValue(tt.inputPosition, tt.inputMask)
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

type inputForFixtures struct {
	orders []float64
	length []float64
}

func featureFixtures(input inputForFixtures) feature.KanjiFeature {
	o := input.orders
	l := input.length

	if len(o) == 0 {
		o = []float64{1, 1, 1, 1, 1, 1}
	}

	if len(l) == 0 {
		l = []float64{1, 1, 1, 1, 1, 1, 1, 1}
	}
	return feature.KanjiFeature{
		Character: "dummy",
		Order:     o,
		Length:    l,
	}
}

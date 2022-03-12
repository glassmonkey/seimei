package feature_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/seimei/feature"
	"github.com/google/go-cmp/cmp"
)

func TestKanjiFeatureOrderCalculator_Mask(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name          string
		inputLength   int
		inputPosition int
		wantMask      []float64
		wantErr       error
	}

	tests := []testdata{
		{
			name:          "-1文字目指定",
			inputLength:   5,
			inputPosition: -1,
			wantMask:      []float64{},
			wantErr:       feature.ErrOutRangeOrderMaskError,
		},
		{
			name:          "(1/5)文字目指定",
			inputLength:   5,
			inputPosition: 0,
			wantMask:      []float64{},
			wantErr:       feature.ErrInvalidOrderMaskError,
		},
		{
			name:          "(2/5)文字目指定",
			inputLength:   5,
			inputPosition: 1,
			wantMask: []float64{
				0, 1, 1, 1, 0, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(3/5)文字目指定",
			inputLength:   5,
			inputPosition: 2,
			wantMask: []float64{
				0, 1, 1, 1, 1, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(4/5)文字目指定",
			inputLength:   5,
			inputPosition: 3,
			wantMask: []float64{
				0, 0, 1, 1, 1, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(5/5)文字目指定",
			inputLength:   5,
			inputPosition: 4,
			wantMask:      []float64{},
			wantErr:       feature.ErrInvalidOrderMaskError,
		},
		{
			name:          "5文字目指定",
			inputLength:   5,
			inputPosition: 5,
			wantMask: []float64{
				0, 0, 1, 1, 1, 0,
			},
			wantErr: feature.ErrOutRangeOrderMaskError,
		},
		{
			name:          "(1/3)文字目指定",
			inputLength:   3,
			inputPosition: 0,
			wantMask:      []float64{},
			wantErr:       feature.ErrInvalidOrderMaskError,
		},
		{
			name:          "(2/3)文字目指定",
			inputLength:   3,
			inputPosition: 1,
			wantMask: []float64{
				0, 0, 1, 1, 0, 0,
			},
			wantErr: nil,
		},
		{
			name:          "(3/3)文字目指定",
			inputLength:   3,
			inputPosition: 2,
			wantMask:      []float64{},
			wantErr:       feature.ErrInvalidOrderMaskError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut := feature.KanjiFeatureOrderCalculator{}
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

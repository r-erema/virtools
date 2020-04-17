package modules

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEightToOneSelectorProcess(t *testing.T) {
	tests := []struct {
		dataInputs, selectInputs []uint8
		wantResult               uint8
	}{
		{
			[]uint8{1, 0, 0, 0, 0, 0, 0, 0},
			[]uint8{0, 0, 0},
			1,
		},
		{
			[]uint8{0, 1, 0, 0, 0, 0, 0, 0},
			[]uint8{1, 0, 0},
			1,
		},
		{
			[]uint8{0, 0, 1, 0, 0, 0, 0, 0},
			[]uint8{0, 1, 0},
			1,
		},
		{
			[]uint8{0, 0, 0, 1, 0, 0, 0, 0},
			[]uint8{1, 1, 0},
			1,
		},
		{
			[]uint8{0, 0, 0, 0, 1, 0, 0, 0},
			[]uint8{0, 0, 1},
			1,
		},
		{
			[]uint8{0, 0, 0, 0, 0, 1, 0, 0},
			[]uint8{1, 0, 1},
			1,
		},
		{
			[]uint8{0, 0, 0, 0, 0, 0, 1, 0},
			[]uint8{0, 1, 1},
			1,
		},
		{
			[]uint8{0, 0, 0, 0, 0, 0, 0, 1},
			[]uint8{1, 1, 1},
			1,
		},
	}
	s := NewSelector8x1()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.wantResult, s.Process(tt.dataInputs, tt.selectInputs))
		})
	}
}

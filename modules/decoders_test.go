package modules

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeToEightDecoder_Process(t *testing.T) {
	tests := []struct {
		dataInput                uint8
		selectInputs, wantResult []uint8
	}{
		{
			1,
			[]uint8{0, 0, 0},
			[]uint8{1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			1,
			[]uint8{1, 0, 0},
			[]uint8{0, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			1,
			[]uint8{0, 1, 0},
			[]uint8{0, 0, 1, 0, 0, 0, 0, 0},
		},
		{
			1,
			[]uint8{1, 1, 0},
			[]uint8{0, 0, 0, 1, 0, 0, 0, 0},
		},
		{
			1,
			[]uint8{0, 0, 1},
			[]uint8{0, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			1,
			[]uint8{1, 0, 1},
			[]uint8{0, 0, 0, 0, 0, 1, 0, 0},
		},
		{
			1,
			[]uint8{0, 1, 1},
			[]uint8{0, 0, 0, 0, 0, 0, 1, 0},
		},
		{
			1,
			[]uint8{1, 1, 1},
			[]uint8{0, 0, 0, 0, 0, 0, 0, 1},
		},
	}
	d := NewThreeToEightDecoder()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.wantResult, d.Process(tt.dataInput, tt.selectInputs))
		})
	}
}

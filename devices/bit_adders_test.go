package devices

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEightBitAdderAdd(t *testing.T) {
	tests := []struct{ input1, input2, wantResult []uint8 }{
		{
			[]uint8{0, 0, 0, 0, 1, 1, 1, 1},
			[]uint8{1, 1, 1, 1, 0, 0, 0, 0},
			[]uint8{0, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			[]uint8{0, 0, 0, 0, 0, 0, 0, 0},
			[]uint8{0, 0, 0, 0, 0, 0, 0, 0},
			[]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[]uint8{1, 1, 1, 1, 1, 1, 1, 1},
			[]uint8{1, 1, 1, 1, 1, 1, 1, 1},
			[]uint8{1, 1, 1, 1, 1, 1, 1, 1, 0},
		},
		{
			[]uint8{1, 0, 1, 1, 0, 0, 1, 1},
			[]uint8{1, 1, 0, 0, 1, 1, 1, 0},
			[]uint8{1, 1, 0, 0, 0, 0, 0, 0, 1},
		},
	}
	a := NewEightBitAdder()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.wantResult, a.Add(0, tt.input1, tt.input2))
		})
	}
}

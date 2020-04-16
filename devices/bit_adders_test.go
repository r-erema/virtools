package devices

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEightBitAdderAdd(t *testing.T) {
	tests := []struct {
		isSubtraction              uint8
		input1, input2, wantResult []uint8
	}{
		{0,
			[]uint8{0, 0, 0, 0, 1, 1, 1, 1},
			[]uint8{1, 1, 1, 1, 0, 0, 0, 0},
			[]uint8{0, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			0,
			[]uint8{0, 0, 0, 0, 0, 0, 0, 0},
			[]uint8{0, 0, 0, 0, 0, 0, 0, 0},
			[]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			0,
			[]uint8{1, 1, 1, 1, 1, 1, 1, 1},
			[]uint8{1, 1, 1, 1, 1, 1, 1, 1},
			[]uint8{1, 1, 1, 1, 1, 1, 1, 1, 0},
		},
		{
			0,
			[]uint8{1, 0, 1, 1, 0, 0, 1, 1},
			[]uint8{1, 1, 0, 0, 1, 1, 1, 0},
			[]uint8{1, 1, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			1,
			[]uint8{1, 1, 1, 1, 1, 1, 0, 1},
			[]uint8{1, 0, 1, 1, 0, 0, 0, 0},
			[]uint8{0, 0, 1, 0, 0, 1, 1, 0, 1},
		},
		{
			1,
			[]uint8{1, 0, 0, 1, 1, 0, 1, 0},
			[]uint8{0, 1, 1, 0, 0, 0, 1, 1},
			[]uint8{0, 0, 0, 1, 1, 0, 1, 1, 1},
		},
	}

	a := NewEightBitAdder()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.wantResult, a.Add(tt.isSubtraction, tt.input1, tt.input2))
		})
	}
}

func TestEightBitLatchedAdderAdd(t *testing.T) {
	a := NewEightBitLatchedAdder()

	input1 := []uint8{1, 0, 1, 1, 0, 0, 1, 1}
	input2 := []uint8{1, 1, 0, 0, 1, 1, 1, 0}
	result := a.Add(0, 0, input1)
	assert.Equal(t, input1, result)
	result = a.Add(0, 0, input2)
	assert.Equal(t, []uint8{1, 0, 0, 0, 0, 0, 0, 1}, result)

	result = a.Add(0, 1, input2)
	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, result)
}

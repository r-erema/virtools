package modules

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecoder1x2Process(t *testing.T) {
	tests := []struct {
		dataInput                uint8
		selectInputs, wantResult []uint8
	}{
		{
			0,
			[]uint8{0},
			[]uint8{0, 0},
		},
		{
			1,
			[]uint8{0},
			[]uint8{1, 0},
		},
		{
			0,
			[]uint8{1},
			[]uint8{0, 0},
		},
		{
			1,
			[]uint8{1},
			[]uint8{0, 1},
		},
	}
	d := NewDecoder1x2()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.wantResult, d.Process(tt.dataInput, tt.selectInputs))
		})
	}
}

func TestDecoder3x8Process(t *testing.T) {
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
	d := NewDecoder3x8()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.wantResult, d.Process(tt.dataInput, tt.selectInputs))
		})
	}
}

func TestDecoder16x65536Process(t *testing.T) {
	d := NewDecoder16x65536()
	result := d.Process(1, []uint8{1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1})
	assert.Equal(t, 65536, len(result))
}

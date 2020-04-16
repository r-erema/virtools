package modules

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEightBitOnesComplementProcess(t *testing.T) {
	tests := []struct {
		isSubtraction     uint8
		input, wantResult []uint8
	}{
		{
			0,
			[]uint8{1, 0, 1, 0, 0, 1, 1, 0},
			[]uint8{1, 0, 1, 0, 0, 1, 1, 0},
		},
		{
			1,
			[]uint8{1, 1, 1, 0, 0, 0, 1, 0},
			[]uint8{0, 0, 0, 1, 1, 1, 0, 1},
		},
	}
	oc := NewEightBitOnesComplement()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.wantResult, oc.Process(tt.isSubtraction, tt.input))
		})
	}
}

package modules

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHalfAdderAdd(t *testing.T) {

	tests := []struct{ input1, input2, wantSum, wantCarry uint8 }{
		{1, 1, 0, 1},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			ha := &halfAdder{}
			sum, carry := ha.Add(tt.input1, tt.input2)
			assert.Equal(t, tt.wantSum, sum)
			assert.Equal(t, tt.wantCarry, carry)
		})
	}
}

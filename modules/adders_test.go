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
    ha := &halfAdder{}
    for i, tt := range tests {
        t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
            sum, carry := ha.Add(tt.input1, tt.input2)
            assert.Equal(t, tt.wantSum, sum)
            assert.Equal(t, tt.wantCarry, carry)
        })
    }
}

func TestFullAdderAdd(t *testing.T) {
    tests := []struct{ carryInput, input1, input2, wantSum, wantCarry uint8 }{
        {0, 0, 0, 0, 0},
        {0, 1, 0, 1, 0},
        {1, 0, 0, 1, 0},
        {1, 1, 0, 0, 1},
        {0, 0, 1, 1, 0},
        {0, 1, 1, 0, 1},
        {1, 0, 1, 0, 1},
        {1, 1, 1, 1, 1},
    }
    fa := &fullAdder{}
    for i, tt := range tests {
        t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
            sum, carry := fa.Add(tt.carryInput, tt.input1, tt.input2)
            assert.Equal(t, tt.wantSum, sum)
            assert.Equal(t, tt.wantCarry, carry)
        })
    }
}

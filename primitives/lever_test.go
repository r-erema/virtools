package primitives

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestLeverOpened(t *testing.T) {

    inputs := []struct {
        inSignal, magnetSignal, wantOut uint8
    }{
        {1, 1, 1},
        {1, 0, 0},
        {0, 1, 0},
        {0, 0, 0},
    }

    inSignals := make(chan LeverInSignals)
    lever := NewLever(inSignals, LeverTypeOpened)

    for _, input := range inputs {
        inSignals <- LeverInSignals{input.inSignal, input.magnetSignal}
        assert.Equal(t, input.wantOut, <-lever.Out())
    }
}
func TestLeverClosed(t *testing.T) {

    inputs := []struct {
        inSignal, magnetSignal, wantOut uint8
    }{
        {1, 0, 1},
        {1, 1, 0},
        {0, 1, 0},
        {0, 0, 0},
    }

    inSignals := make(chan LeverInSignals)
    lever := NewLever(inSignals, LeverTypeClosed)

    for _, input := range inputs {
        inSignals <- LeverInSignals{input.inSignal, input.magnetSignal}
        assert.Equal(t, input.wantOut, <-lever.Out())
    }
}

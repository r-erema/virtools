package primitives

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelayOpenedLever(t *testing.T) {
	inputs := []struct {
		inSignal, leverSignal, wantOut uint8
	}{
		{1, 1, 1},
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, 0},
	}

	relayInSignals := make(chan RelayInSignals)
	relay := NewRelay(relayInSignals, LeverTypeOpened)

	for _, input := range inputs {
		relayInSignals <- RelayInSignals{input.inSignal, input.leverSignal}
		assert.Equal(t, input.wantOut, <-relay.Out())
	}
}

func TestRelayClosedLever(t *testing.T) {
	inputs := []struct {
		inSignal, leverSignal, wantOut uint8
	}{
		{1, 1, 0},
		{1, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}

	relayInSignals := make(chan RelayInSignals)
	relay := NewRelay(relayInSignals, LeverTypeClosed)

	for _, input := range inputs {
		relayInSignals <- RelayInSignals{input.inSignal, input.leverSignal}
		assert.Equal(t, input.wantOut, <-relay.Out())
	}
}

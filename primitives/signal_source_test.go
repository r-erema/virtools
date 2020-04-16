package primitives

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewSignalSource(t *testing.T) {

	battery := NewSignalSource()

	var receivedSignals []uint8
	go func() {
		for signal := range battery.Out() {
			receivedSignals = append(receivedSignals, signal)
		}
	}()

	<-time.After(time.Millisecond)

	assert.Greater(t, len(receivedSignals), 0)
}

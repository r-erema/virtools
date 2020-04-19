package primitives

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOscillator(t *testing.T) {
	o := NewOscillator(time.Millisecond)
	iterationsCount := 20
	var receivedSignals []uint8
	for i := 0; i < iterationsCount; i++ {
		receivedSignals = append(receivedSignals, <-o.Out())
	}
	assert.Equal(t, []uint8{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, receivedSignals)
}

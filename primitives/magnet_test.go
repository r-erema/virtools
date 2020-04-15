package primitives

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestMagnet(t *testing.T) {
	battery := NewSignalSource()
	magnet := NewMagnet(battery.Out())

	var receivedSignals []uint8
	mu := sync.Mutex{}
	go func() {
		for signal := range magnet.Out() {
			mu.Lock()
			receivedSignals = append(receivedSignals, signal)
			mu.Unlock()
		}
	}()

	<-time.After(time.Millisecond)

	mu.Lock()
	assert.Greater(t, len(receivedSignals), 0)
	mu.Unlock()
}

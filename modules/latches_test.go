package modules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEightBitLatchProcess(t *testing.T) {
	l := NewEightBitLatch()
	data := []uint8{1, 1, 1, 0, 0, 0, 0, 1}

	result := l.Process(0, 0, data)
	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, result)
	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, l.GetCurrentQuits())

	result = l.Process(1, 0, data)
	assert.Equal(t, data, result)
	assert.Equal(t, data, l.GetCurrentQuits())

	result = l.Process(1, 1, data)
	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, result)
	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, l.GetCurrentQuits())
}

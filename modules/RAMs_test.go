package modules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRAM8x1(t *testing.T) {
	ram := NewRAM8x1()
	result := ram.Process([]uint8{0, 0, 1}, 1, 0)
	assert.Equal(t, uint8(0), result)
	ram.Process([]uint8{0, 0, 1}, 1, 1)
	result = ram.Process([]uint8{0, 0, 1}, 1, 0)
	assert.Equal(t, uint8(1), result)
	result = ram.Process([]uint8{0, 0, 1}, 0, 0)
	assert.Equal(t, uint8(1), result)
	result = ram.Process([]uint8{0, 1, 1}, 0, 0)
	assert.Equal(t, uint8(0), result)
	ram.Process([]uint8{0, 1, 1}, 1, 1)
	ram.Process([]uint8{1, 0, 1}, 1, 1)
	ram.Process([]uint8{0, 1, 0}, 1, 1)
	assert.Equal(t, uint8(1), ram.Process([]uint8{0, 1, 1}, 0, 0))
	assert.Equal(t, uint8(1), ram.Process([]uint8{1, 0, 1}, 0, 0))
	assert.Equal(t, uint8(1), ram.Process([]uint8{0, 1, 0}, 0, 0))
}

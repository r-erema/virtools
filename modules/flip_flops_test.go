package modules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsFlipFlopProcess(t *testing.T) {

	trigger := NewRsFlipFlop()

	quit, quitOpposite := trigger.Process(1, 0)
	assert.Equal(t, uint8(0), quit)
	assert.Equal(t, uint8(1), quitOpposite)

	quit, quitOpposite = trigger.Process(0, 0)
	assert.Equal(t, uint8(0), quit)
	assert.Equal(t, uint8(1), quitOpposite)

	quit, quitOpposite = trigger.Process(0, 1)
	assert.Equal(t, uint8(1), quit)
	assert.Equal(t, uint8(0), quitOpposite)

	quit, quitOpposite = trigger.Process(1, 1)
	assert.Equal(t, uint8(0), quit)
	assert.Equal(t, uint8(0), quitOpposite)

	quit, quitOpposite = trigger.Process(1, 0)
	assert.Equal(t, uint8(0), quit)
	assert.Equal(t, uint8(1), quitOpposite)

	quit, quitOpposite = trigger.Process(0, 0)
	assert.Equal(t, uint8(0), quit)
	assert.Equal(t, uint8(1), quitOpposite)

}

func TestDTypeFlipFlopProcess(t *testing.T) {

	trigger := NewDTypeFlipFlop()

	quit, quitOpposite := trigger.Process(0, 1, 0)
	assert.Equal(t, uint8(0), quit)
	assert.Equal(t, uint8(1), quitOpposite)

	quit, quitOpposite = trigger.Process(0, 1, 1)
	assert.Equal(t, uint8(1), quit)
	assert.Equal(t, uint8(0), quitOpposite)

	quit, quitOpposite = trigger.Process(0, 0, 1)
	assert.Equal(t, uint8(1), quit)
	assert.Equal(t, uint8(0), quitOpposite)

	quit, quitOpposite = trigger.Process(0, 0, 0)
	assert.Equal(t, uint8(1), quit)
	assert.Equal(t, uint8(0), quitOpposite)

	quit, quitOpposite = trigger.Process(1, 1, 1)
	assert.Equal(t, uint8(0), quit)
	assert.Equal(t, uint8(0), quitOpposite)

}

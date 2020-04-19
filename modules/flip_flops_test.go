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

func TestDTypeLevelTriggeredFlipFlopProcess(t *testing.T) {

	trigger := NewDTypeLevelTriggeredFlipFlop()

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

func TestDTypeEdgeTriggeredFlipFlopProcess(t *testing.T) {
	trigger := NewDTypeEdgeTriggeredFlipFlop()
	q, qo := trigger.Process(1, 1)
	assert.Equal(t, uint8(1), q)
	assert.Equal(t, uint8(0), qo)

	trigger = NewDTypeEdgeTriggeredFlipFlop()
	q, qo = trigger.Process(1, 0)
	assert.Equal(t, uint8(0), q)
	assert.Equal(t, uint8(1), qo)
}

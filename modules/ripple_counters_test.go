package modules

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"virtools/primitives"
)

func TestRippleCounter16bitProcess(t *testing.T) {
	c := NewRippleCounter16bit()
	o := primitives.NewOscillator(time.Millisecond)
	assert.Equal(t, []uint8{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, c.Process(<-o.Out()))
	assert.Equal(t, []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, c.Process(<-o.Out()))
	assert.Equal(t, []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, c.Process(<-o.Out()))
	assert.Equal(t, []uint8{1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, c.Process(<-o.Out()))

}

package cpu_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
	"virtools/high_level/cpu"
)

func TestCPU(t *testing.T) {
	t.Parallel()

	tickTime := time.Millisecond * 100

	ram := &cpu.RandomAccessMemory{
		0: "LOAD 6",
		1: "ADD 7",
		2: "STORE 6",
		3: "JUMP 1",
		4: "0",
		5: "0",
		6: "0",
		7: "1",
	}
	c := cpu.BuildCPU(tickTime, ram)

	stop := make(chan struct{})

	time.AfterFunc(tickTime*10, func() {
		stop <- struct{}{}
	})

	c.Run(stop)

	sum, err := strconv.Atoi((*ram)[6])
	require.NoError(t, err)
	assert.Greater(t, sum, 0)
}

package modules

import (
	"fmt"
	"strconv"
	"strings"
)

type ram8x1 struct {
	decoder  decoder3x8
	triggers []dTypeFlipFlop
	selector selector8x1
}

func NewRAM8x1() *ram8x1 {
	var triggers []dTypeFlipFlop
	for i := 0; i < 8; i++ {
		triggers = append(triggers, *NewDTypeFlipFlop())
	}
	return &ram8x1{
		decoder:  *NewDecoder3x8(),
		triggers: triggers,
		selector: *NewSelector8x1(),
	}
}

func (r *ram8x1) Process(addresses []uint8, data, write uint8) (out uint8) {
	decoderOut := r.decoder.Process(write, addresses)
	var triggersOuts []uint8
	for i, t := range r.triggers {
		q, _ := t.Process(0, decoderOut[i], data)
		triggersOuts = append(triggersOuts, q)
	}
	return r.selector.Process(triggersOuts, addresses)
}

//Impossible to imitate real RAM64Kx8, need more than 4.5million relay, not enough real RAM and CPU for that
//So create a stub...
type ram64Kx8Stub struct {
	memoryStub [][]uint8
}

func NewRam64Kx8Stub() *ram64Kx8Stub {
	return &ram64Kx8Stub{make([][]uint8, 65536)}
}

func (r *ram64Kx8Stub) Process(addr, data []uint8, write uint8) []uint8 {
	bin := strings.Replace(strings.Trim(fmt.Sprint(addr), "[]"), " ", "", -1)
	i, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	if write == 1 {
		r.memoryStub[i] = data
	}
	return r.memoryStub[i]
}

package modules

import (
	"virtools/gates"
)

type decoder1x2 struct{}

func NewDecoder1x2() *decoder1x2 {
	return &decoder1x2{}
}

func (d *decoder1x2) Process(data uint8, selectInputs []uint8) (outs []uint8) {
	return []uint8{
		gates.AND(data, gates.NOT(selectInputs[0])),
		gates.AND(data, selectInputs[0]),
	}
}

type decoder3x8 struct{}

func NewDecoder3x8() *decoder3x8 {
	return &decoder3x8{}
}

func (d *decoder3x8) Process(data uint8, selectInputs []uint8) (outs []uint8) {

	sel0 := selectInputs[0]
	sel1 := selectInputs[1]
	sel2 := selectInputs[2]

	selInv0 := gates.NOT(sel0)
	selInv1 := gates.NOT(sel1)
	selInv2 := gates.NOT(sel2)

	return []uint8{
		gates.MultipleAND(data, selInv0, selInv1, selInv2),
		gates.MultipleAND(data, sel0, selInv1, selInv2),
		gates.MultipleAND(data, sel1, selInv0, selInv2),
		gates.MultipleAND(data, sel0, sel1, selInv2),
		gates.MultipleAND(data, sel2, selInv0, selInv1),
		gates.MultipleAND(data, sel0, sel2, selInv1),
		gates.MultipleAND(data, sel1, sel2, selInv0),
		gates.MultipleAND(data, sel0, sel1, sel2),
	}
}

type decoder4x16 struct {
	decoder3x8  decoder3x8
	decoders1x2 []decoder1x2
}

func NewDecoder4x16() *decoder4x16 {
	var decoders1x2 []decoder1x2
	for i := 0; i < 8; i++ {
		decoders1x2 = append(decoders1x2, *NewDecoder1x2())
	}
	return &decoder4x16{decoder3x8: *NewDecoder3x8(), decoders1x2: decoders1x2}
}

func (d *decoder4x16) Process(data uint8, selectInputs []uint8) (outs []uint8) {
	result := d.decoder3x8.Process(data, selectInputs[1:])
	for i, o := range result {
		outs = append(outs, d.decoders1x2[i].Process(o, selectInputs[:1])...)
	}
	return
}

type decoder8x256 struct {
	frontDecoder4x16 decoder4x16
	backDecoders4x16 []decoder4x16
}

func NewDecoder8x256() *decoder8x256 {
	var decoders4x16 []decoder4x16
	for i := 0; i < 16; i++ {
		decoders4x16 = append(decoders4x16, *NewDecoder4x16())
	}
	return &decoder8x256{frontDecoder4x16: *NewDecoder4x16(), backDecoders4x16: decoders4x16}
}

func (d *decoder8x256) Process(data uint8, selectInputs []uint8) (outs []uint8) {
	result := d.frontDecoder4x16.Process(data, selectInputs[4:])
	for i, o := range result {
		outs = append(outs, d.backDecoders4x16[i].Process(o, selectInputs[:4])...)
	}
	return
}

type decoder16x65536 struct {
	frontDecoder8x256 decoder8x256
	backDecoders8x256 []decoder8x256
}

func NewDecoder16x65536() *decoder16x65536 {
	var backDecoders8x256 []decoder8x256
	for i := 0; i < 256; i++ {
		backDecoders8x256 = append(backDecoders8x256, *NewDecoder8x256())
	}
	return &decoder16x65536{frontDecoder8x256: *NewDecoder8x256(), backDecoders8x256: backDecoders8x256}
}

func (d *decoder16x65536) Process(data uint8, selectInputs []uint8) (outs []uint8) {
	result := d.frontDecoder8x256.Process(data, selectInputs[8:])
	for i, o := range result {
		outs = append(outs, d.backDecoders8x256[i].Process(o, selectInputs[:8])...)
	}
	return
}

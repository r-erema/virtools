package modules

import "virtools/gates"

type ThreeToEightDecoder struct{}

func NewThreeToEightDecoder() *ThreeToEightDecoder {
	return &ThreeToEightDecoder{}
}

func (d *ThreeToEightDecoder) Process(data uint8, selectInputs []uint8) (outs []uint8) {

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
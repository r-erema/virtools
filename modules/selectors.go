package modules

import "virtools/gates"

type selector8x1 struct{}

func NewSelector8x1() *selector8x1 {
	return &selector8x1{}
}

func (s *selector8x1) Process(dataInputs []uint8, selectInputs []uint8) uint8 {

	sel0 := selectInputs[0]
	sel1 := selectInputs[1]
	sel2 := selectInputs[2]

	selInv0 := gates.NOT(sel0)
	selInv1 := gates.NOT(sel1)
	selInv2 := gates.NOT(sel2)

	d0 := gates.MultipleAND(dataInputs[0], selInv0, selInv1, selInv2)
	d1 := gates.MultipleAND(dataInputs[1], sel0, selInv1, selInv2)
	d2 := gates.MultipleAND(dataInputs[2], sel1, selInv0, selInv2)
	d3 := gates.MultipleAND(dataInputs[3], sel0, sel1, selInv2)
	d4 := gates.MultipleAND(dataInputs[4], sel2, selInv0, selInv1)
	d5 := gates.MultipleAND(dataInputs[5], sel0, sel2, selInv1)
	d6 := gates.MultipleAND(dataInputs[6], sel1, sel2, selInv0)
	d7 := gates.MultipleAND(dataInputs[7], sel0, sel1, sel2)

	return gates.MultipleOR(d0, d1, d2, d3, d4, d5, d6, d7)
}

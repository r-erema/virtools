package modules

import "virtools/gates"

type EightToOneSelector struct{}

func NewEightToOneSelector() *EightToOneSelector {
	return &EightToOneSelector{}
}

func (s *EightToOneSelector) Process(dataInputs []uint8, selectInputs []uint8) uint8 {

	si0 := selectInputs[0]
	si1 := selectInputs[1]
	si2 := selectInputs[2]

	siInv0 := gates.NOT(si0)
	siInv1 := gates.NOT(si1)
	siInv2 := gates.NOT(si2)

	d0 := gates.MultipleAND(dataInputs[0], siInv0, siInv1, siInv2)
	d1 := gates.MultipleAND(dataInputs[1], si0, siInv1, siInv2)
	d2 := gates.MultipleAND(dataInputs[2], si1, siInv0, siInv2)
	d3 := gates.MultipleAND(dataInputs[3], si0, si1, siInv2)
	d4 := gates.MultipleAND(dataInputs[4], si1, siInv0, siInv1)
	d5 := gates.MultipleAND(dataInputs[5], si0, si2, siInv1)
	d6 := gates.MultipleAND(dataInputs[6], si1, si2, siInv0)
	d7 := gates.MultipleAND(dataInputs[7], si0, si1, si2)

	return gates.MultipleOR(d0, d1, d2, d3, d4, d5, d6, d7)
}

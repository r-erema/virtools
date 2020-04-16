package modules

import "virtools/gates"

type eightBitOnesComplement struct{}

func NewEightBitOnesComplement() *eightBitOnesComplement {
	return &eightBitOnesComplement{}
}

func (oc *eightBitOnesComplement) Process(isSubtraction uint8, input []uint8) []uint8 {
	return []uint8{
		gates.XOR(isSubtraction, input[0]),
		gates.XOR(isSubtraction, input[1]),
		gates.XOR(isSubtraction, input[2]),
		gates.XOR(isSubtraction, input[3]),
		gates.XOR(isSubtraction, input[4]),
		gates.XOR(isSubtraction, input[5]),
		gates.XOR(isSubtraction, input[6]),
		gates.XOR(isSubtraction, input[7]),
	}
}

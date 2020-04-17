package modules

import "virtools/gates"

type OnesComplement8Bit struct{}

func NewOnesComplement8Bit() *OnesComplement8Bit {
	return &OnesComplement8Bit{}
}

func (oc *OnesComplement8Bit) Process(isSubtraction uint8, input []uint8) []uint8 {
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

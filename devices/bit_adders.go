package devices

import (
	"virtools/gates"
	"virtools/modules"
)

type eightBitAdder struct{}

func NewEightBitAdder() *eightBitAdder {
	return &eightBitAdder{}
}

func (a *eightBitAdder) Add(isSubtraction uint8, input1, input2 []uint8) (result []uint8) {
	adder1 := modules.NewFullAdder()
	adder2 := modules.NewFullAdder()
	adder3 := modules.NewFullAdder()
	adder4 := modules.NewFullAdder()
	adder5 := modules.NewFullAdder()
	adder6 := modules.NewFullAdder()
	adder7 := modules.NewFullAdder()
	adder8 := modules.NewFullAdder()

	complement := modules.NewEightBitOnesComplement()
	input2 = complement.Process(isSubtraction, input2)

	out1, carry := adder1.Add(isSubtraction, input1[7], input2[7])
	out2, carry := adder2.Add(carry, input1[6], input2[6])
	out3, carry := adder3.Add(carry, input1[5], input2[5])
	out4, carry := adder4.Add(carry, input1[4], input2[4])
	out5, carry := adder5.Add(carry, input1[3], input2[3])
	out6, carry := adder6.Add(carry, input1[2], input2[2])
	out7, carry := adder7.Add(carry, input1[1], input2[1])
	out8, carry := adder8.Add(carry, input1[0], input2[0])

	overflowOrUnderflow := gates.XOR(carry, isSubtraction)

	return []uint8{overflowOrUnderflow, out8, out7, out6, out5, out4, out3, out2, out1}
}

type eightBitLatchedAdder struct {
	adder *eightBitAdder
	latch *modules.EightBitLatch
}

func NewEightBitLatchedAdder() *eightBitLatchedAdder {
	return &eightBitLatchedAdder{
		NewEightBitAdder(),
		modules.NewEightBitLatch(),
	}
}

func (a *eightBitLatchedAdder) Add(isSubtraction, clear uint8, input []uint8) []uint8 {
	result := a.adder.Add(isSubtraction, input, a.latch.GetCurrentQuits())[1:]
	a.latch.Process(1, clear, result)
	return a.latch.GetCurrentQuits()
}

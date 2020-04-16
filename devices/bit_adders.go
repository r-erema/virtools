package devices

import "virtools/modules"

type eightBitAdder struct{}

func NewEightBitAdder() *eightBitAdder {
	return &eightBitAdder{}
}

func (a *eightBitAdder) Add(carry uint8, input1, input2 []uint8) (result []uint8) {
	adder1 := modules.NewFullAdder()
	adder2 := modules.NewFullAdder()
	adder3 := modules.NewFullAdder()
	adder4 := modules.NewFullAdder()
	adder5 := modules.NewFullAdder()
	adder6 := modules.NewFullAdder()
	adder7 := modules.NewFullAdder()
	adder8 := modules.NewFullAdder()

	out1, carry := adder1.Add(carry, input1[7], input2[7])
	out2, carry := adder2.Add(carry, input1[6], input2[6])
	out3, carry := adder3.Add(carry, input1[5], input2[5])
	out4, carry := adder4.Add(carry, input1[4], input2[4])
	out5, carry := adder5.Add(carry, input1[3], input2[3])
	out6, carry := adder6.Add(carry, input1[2], input2[2])
	out7, carry := adder7.Add(carry, input1[1], input2[1])
	out8, carry := adder8.Add(carry, input1[0], input2[0])

	return []uint8{carry, out8, out7, out6, out5, out4, out3, out2, out1}
}

type sixteenBitAdder struct{}

func NewSixteenBitAdder() *sixteenBitAdder {
	return &sixteenBitAdder{}
}

func (a *sixteenBitAdder) Add(carry uint8, input1, input2 []uint8) (result []uint8) {
	adder1 := NewEightBitAdder()
	adder2 := NewEightBitAdder()

	adder1Input1 := input1[8:]
	adder1Input2 := input2[8:]
	adder2Input1 := input1[:8]
	adder2Input2 := input2[:8]

	result1 := adder1.Add(carry, adder1Input1, adder1Input2)
	result2 := adder2.Add(result1[0], adder2Input1, adder2Input2)

	return append(result2, result1[1:]...)
}

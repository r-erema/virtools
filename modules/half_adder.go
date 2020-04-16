package modules

import "virtools/gates"

type halfAdder struct{}

func NewHalfAdder() *halfAdder {
	return &halfAdder{}
}

func (ha *halfAdder) Add(input1, input2 uint8) (sum, carry uint8) {
	sum = gates.XOR(input1, input2)
	carry = gates.AND(input1, input2)
	return
}

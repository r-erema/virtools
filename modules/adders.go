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

type fullAdder struct{}

func NewFullAdder() *fullAdder {
    return &fullAdder{}
}

func (ha *fullAdder) Add(carryInput, input1, input2 uint8) (sum, carry uint8) {
    ha1 := NewHalfAdder()
    ha2 := NewHalfAdder()
    sum2, carry2 := ha2.Add(input1, input2)
    sum1, carry1 := ha1.Add(carryInput, sum2)
    return sum1, gates.OR(carry1, carry2)
}

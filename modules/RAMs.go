package modules

type ram8x1 struct {
	decoder  decoder3x8
	triggers []dTypeFlipFlop
	selector selector8x1
}

func NewRAM8x1() *ram8x1 {
	var triggers []dTypeFlipFlop
	for i := 0; i < 8; i++ {
		triggers = append(triggers, *NewDTypeFlipFlop())
	}
	return &ram8x1{
		decoder:  *NewDecoder3x8(),
		triggers: triggers,
		selector: *NewSelector8x1(),
	}
}

func (r *ram8x1) Process(addresses []uint8, data, write uint8) (out uint8) {
	decoderOut := r.decoder.Process(write, addresses)
	var triggersOuts []uint8
	for i, t := range r.triggers {
		q, _ := t.Process(0, decoderOut[i], data)
		triggersOuts = append(triggersOuts, q)
	}
	return r.selector.Process(triggersOuts, addresses)
}

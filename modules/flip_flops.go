package modules

import "virtools/gates"

type rsFlipFlop struct {
	quit, quitOpposite uint8
}

func (ff *rsFlipFlop) Quit() uint8 {
	return ff.quit
}

func NewRsFlipFlop() *rsFlipFlop {
	return &rsFlipFlop{0, 0}
}

func (ff *rsFlipFlop) Process(reset, set uint8) (quit, quitOpposite uint8) {
	ff.quitOpposite = gates.NOR(ff.quit, set)
	ff.quit = gates.NOR(ff.quitOpposite, reset)
	return ff.quit, ff.quitOpposite
}

type dTypeFlipFlop struct {
	rsTrigger *rsFlipFlop
}

func NewDTypeFlipFlop() *dTypeFlipFlop {
	return &dTypeFlipFlop{NewRsFlipFlop()}
}

func (t *dTypeFlipFlop) Process(clear, clock, data uint8) (quit, quitOpposite uint8) {
	return t.rsTrigger.Process(
		gates.OR(clear, gates.AND(gates.NOT(data), clock)),
		gates.AND(clock, data),
	)
}

func (t *dTypeFlipFlop) Quit() uint8 {
	return t.rsTrigger.Quit()
}

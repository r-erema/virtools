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

type dTypeLevelTriggeredFlipFlop struct {
	rsTrigger *rsFlipFlop
}

func NewDTypeLevelTriggeredFlipFlop() *dTypeLevelTriggeredFlipFlop {
	return &dTypeLevelTriggeredFlipFlop{NewRsFlipFlop()}
}

func (t *dTypeLevelTriggeredFlipFlop) Process(clear, clock, dataInput uint8) (quit, quitOpposite uint8) {
	return t.rsTrigger.Process(
		gates.OR(clear, gates.AND(gates.NOT(dataInput), clock)),
		gates.AND(clock, dataInput),
	)
}

func (t *dTypeLevelTriggeredFlipFlop) Quit() uint8 {
	return t.rsTrigger.Quit()
}

type dTypeEdgeTriggeredFlipFlop struct {
	rsTriggerStep1 *rsFlipFlop
	rsTriggerStep2 *rsFlipFlop
}

func NewDTypeEdgeTriggeredFlipFlop() *dTypeEdgeTriggeredFlipFlop {
	return &dTypeEdgeTriggeredFlipFlop{rsTriggerStep1: NewRsFlipFlop(), rsTriggerStep2: NewRsFlipFlop()}
}

func (t *dTypeEdgeTriggeredFlipFlop) Process(clock, dataInput uint8) (quit, quitOpposite uint8) {
	step1Result1 := gates.AND(dataInput, gates.NOT(clock))
	step1Result2 := gates.AND(gates.NOT(clock), gates.NOT(dataInput))
	qLevel1, qOppositeLevel1 := t.rsTriggerStep1.Process(step1Result1, step1Result2)
	step2Result1 := gates.AND(qLevel1, clock)
	step2Result2 := gates.AND(qOppositeLevel1, clock)
	return t.rsTriggerStep2.Process(step2Result1, step2Result2)
}

func (t *dTypeEdgeTriggeredFlipFlop) Quits() (uint8, uint8) {
	return t.rsTriggerStep2.quit, t.rsTriggerStep2.quitOpposite
}

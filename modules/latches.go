package modules

type Latch8Bit struct {
	trigger1,
	trigger2,
	trigger3,
	trigger4,
	trigger5,
	trigger6,
	trigger7,
	trigger8 *dTypeLevelTriggeredFlipFlop
}

func NewLatch8Bit() *Latch8Bit {
	l := &Latch8Bit{}
	l.trigger1 = NewDTypeLevelTriggeredFlipFlop()
	l.trigger2 = NewDTypeLevelTriggeredFlipFlop()
	l.trigger3 = NewDTypeLevelTriggeredFlipFlop()
	l.trigger4 = NewDTypeLevelTriggeredFlipFlop()
	l.trigger5 = NewDTypeLevelTriggeredFlipFlop()
	l.trigger6 = NewDTypeLevelTriggeredFlipFlop()
	l.trigger7 = NewDTypeLevelTriggeredFlipFlop()
	l.trigger8 = NewDTypeLevelTriggeredFlipFlop()
	return l
}

func (l *Latch8Bit) Process(clear, clock uint8, input []uint8) []uint8 {
	quit1, _ := l.trigger1.Process(clear, clock, input[0])
	quit2, _ := l.trigger2.Process(clear, clock, input[1])
	quit3, _ := l.trigger3.Process(clear, clock, input[2])
	quit4, _ := l.trigger4.Process(clear, clock, input[3])
	quit5, _ := l.trigger5.Process(clear, clock, input[4])
	quit6, _ := l.trigger6.Process(clear, clock, input[5])
	quit7, _ := l.trigger7.Process(clear, clock, input[6])
	quit8, _ := l.trigger8.Process(clear, clock, input[7])
	return []uint8{quit1, quit2, quit3, quit4, quit5, quit6, quit7, quit8}
}

func (l *Latch8Bit) GetCurrentQuits() []uint8 {
	return []uint8{
		l.trigger1.Quit(),
		l.trigger2.Quit(),
		l.trigger3.Quit(),
		l.trigger4.Quit(),
		l.trigger5.Quit(),
		l.trigger6.Quit(),
		l.trigger7.Quit(),
		l.trigger8.Quit(),
	}
}

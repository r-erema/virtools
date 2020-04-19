package primitives

import "time"

type oscillator struct {
	out   chan uint8
	delay time.Duration
}

func NewOscillator(delay time.Duration) *oscillator {
	o := &oscillator{out: make(chan uint8), delay: delay}
	leverInSignals := make(chan LeverInSignals)
	lever := NewLever(leverInSignals, LeverTypeClosed)
	battery := NewSignalSource()
	battery.Out()

	magnetInSignals := make(chan uint8)
	magnet := NewMagnet(magnetInSignals)

	var magnetOutSignal uint8 = 0
	go func() {
		for {
			<-time.After(o.delay)
			batteryOutSignal := <-battery.Out()
			leverInSignals <- LeverInSignals{inSignal: batteryOutSignal, magnetSignal: magnetOutSignal}
			leverOut := <-lever.Out()
			magnetInSignals <- leverOut
			o.out <- leverOut
			magnetOutSignal = <-magnet.Out()
		}
	}()

	return o
}

func (o *oscillator) Out() chan uint8 {
	return o.out
}

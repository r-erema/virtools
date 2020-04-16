package primitives

type RelayInSignals struct {
	InSignal, LeverSignal uint8
}

type relay struct {
	lever *lever
}

func NewRelay(inSignals chan RelayInSignals, leverType string) *relay {

	leverInSignals := make(chan LeverInSignals)
	lever := NewLever(leverInSignals, leverType)
	r := &relay{lever: lever}

	magnetInSignals := make(chan uint8)
	magnet := NewMagnet(magnetInSignals)
	go func() {
		for signals := range inSignals {
			magnetInSignals <- signals.InSignal
			magnetOut := <-magnet.Out()
			leverInSignals <- LeverInSignals{inSignal: signals.LeverSignal, magnetSignal: magnetOut}
		}
	}()
	return r
}

func (r *relay) Out() chan uint8 {
	return r.lever.Out()
}

package primitives

const (
	LeverTypeOpened = "opened"
	LeverTypeClosed = "closed"
)

type LeverInSignals struct {
	inSignal, magnetSignal uint8
}

type lever struct {
	out       chan uint8
	leverType string
}

func NewLever(inSignals chan LeverInSignals, leverType string) *lever {
	l := &lever{out: make(chan uint8), leverType: leverType}
	go func() {
		zero := uint8(0)
		for signals := range inSignals {
			if (signals.magnetSignal > zero && l.leverType == LeverTypeOpened) ||
				(signals.magnetSignal <= zero && l.leverType == LeverTypeClosed) {
				l.out <- signals.inSignal
			} else {
				l.out <- 0
			}
		}
	}()
	return l
}

func (l *lever) Out() chan uint8 {
	return l.out
}

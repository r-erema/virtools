package primitives

type magnet struct {
	out chan uint8
}

func NewMagnet(inSignals chan uint8) *magnet {
	m := &magnet{out: make(chan uint8)}
	go func() {
		for signal := range inSignals {
			m.out <- signal
		}
	}()
	return m
}

func (m *magnet) Out() chan uint8 {
	return m.out
}

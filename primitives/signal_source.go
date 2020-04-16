package primitives

type signalSource struct {
    out chan uint8
}

func NewSignalSource() *signalSource {
    s := &signalSource{out: make(chan uint8)}
    go func() {
        for {
            s.out <- 1
        }
    }()

    return s
}

func (s *signalSource) Out() chan uint8 {
    return s.out
}

package modules

import "virtools/gates"

type rippleCounter16bit struct {
	triggers []*dTypeEdgeTriggeredFlipFlop
}

func NewRippleCounter16bit() *rippleCounter16bit {
	triggers := make([]*dTypeEdgeTriggeredFlipFlop, 15)
	for i := 0; i < 15; i++ {
		triggers[i] = NewDTypeEdgeTriggeredFlipFlop()
	}
	return &rippleCounter16bit{triggers: triggers}
}

func (c *rippleCounter16bit) Process(clock uint8) (out []uint8) {
	nextClock := clock
	out = append(out, gates.NOT(clock))
	for _, trigger := range c.triggers {
		q, qo := trigger.Quits()

		if q == qo {
			if q == 1 {
				qo = 0
			} else if q == 0 {
				qo = 1
			}
		}

		q, _ = trigger.Process(nextClock, qo)
		nextClock = qo
		out = append(out, q)
	}
	return
}

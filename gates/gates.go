package gates

import "virtools/primitives"

func AND(input1, input2 uint8) uint8 {

	relayIn1, relayIn2 := make(chan primitives.RelayInSignals), make(chan primitives.RelayInSignals)
	relay1, relay2 := primitives.NewRelay(relayIn1, primitives.LeverTypeOpened), primitives.NewRelay(relayIn2, primitives.LeverTypeOpened)

	battery := primitives.NewSignalSource()
	go func() {
		for batterySignal := range battery.Out() {
			relayIn1 <- primitives.RelayInSignals{InSignal: input1, LeverSignal: batterySignal}
			relayIn2 <- primitives.RelayInSignals{InSignal: input2, LeverSignal: <-relay1.Out()}
		}
	}()

	return <-relay2.Out()
}

func MultipleAND(inputs ...uint8) uint8 {
	currentResult := inputs[0]
	for i := 1; i < len(inputs); i++ {
		currentResult = AND(currentResult, inputs[i])
	}
	return currentResult
}

func NAND(input1, input2 uint8) uint8 {

	relayIn1, relayIn2 := make(chan primitives.RelayInSignals), make(chan primitives.RelayInSignals)
	relay1, relay2 := primitives.NewRelay(relayIn1, primitives.LeverTypeClosed), primitives.NewRelay(relayIn2, primitives.LeverTypeClosed)

	battery1 := primitives.NewSignalSource()
	go func() {
		for batterySignal := range battery1.Out() {
			relayIn1 <- primitives.RelayInSignals{InSignal: input1, LeverSignal: batterySignal}
		}
	}()

	battery2 := primitives.NewSignalSource()
	go func() {
		for batterySignal := range battery2.Out() {
			relayIn2 <- primitives.RelayInSignals{InSignal: input2, LeverSignal: batterySignal}
		}
	}()

	result1 := <-relay1.Out()
	zero := uint8(0)
	if result1 > zero {
		return result1
	}

	result2 := <-relay2.Out()
	if result2 > zero {
		return result2
	}

	return 0
}

func OR(input1, input2 uint8) uint8 {

	relayIn1, relayIn2 := make(chan primitives.RelayInSignals), make(chan primitives.RelayInSignals)
	relay1, relay2 := primitives.NewRelay(relayIn1, primitives.LeverTypeOpened), primitives.NewRelay(relayIn2, primitives.LeverTypeOpened)

	battery1 := primitives.NewSignalSource()
	go func() {
		for batterySignal := range battery1.Out() {
			relayIn1 <- primitives.RelayInSignals{InSignal: input1, LeverSignal: batterySignal}
		}
	}()

	battery2 := primitives.NewSignalSource()
	go func() {
		for batterySignal := range battery2.Out() {
			relayIn2 <- primitives.RelayInSignals{InSignal: input2, LeverSignal: batterySignal}
		}
	}()

	result1 := <-relay1.Out()
	zero := uint8(0)
	if result1 > zero {
		return result1
	}

	result2 := <-relay2.Out()
	if result2 > zero {
		return result2
	}

	return 0
}

func MultipleOR(inputs ...uint8) uint8 {
	currentResult := inputs[0]
	for i := 1; i < len(inputs); i++ {
		currentResult = OR(currentResult, inputs[i])
	}
	return currentResult
}

func NOR(input1, input2 uint8) uint8 {
	relayIn1, relayIn2 := make(chan primitives.RelayInSignals), make(chan primitives.RelayInSignals)
	relay1, relay2 := primitives.NewRelay(relayIn1, primitives.LeverTypeClosed), primitives.NewRelay(relayIn2, primitives.LeverTypeClosed)

	battery := primitives.NewSignalSource()
	go func() {
		for batterySignal := range battery.Out() {
			relayIn1 <- primitives.RelayInSignals{InSignal: input1, LeverSignal: batterySignal}
			relayIn2 <- primitives.RelayInSignals{InSignal: input2, LeverSignal: <-relay1.Out()}
		}
	}()

	return <-relay2.Out()
}

func XOR(input1, input2 uint8) uint8 {
	return AND(
		OR(input1, input2),
		NAND(input1, input2),
	)
}

func NOT(input uint8) uint8 {
	relayIn := make(chan primitives.RelayInSignals)
	relay := primitives.NewRelay(relayIn, primitives.LeverTypeClosed)

	battery := primitives.NewSignalSource()
	go func() {
		for batterySignal := range battery.Out() {
			relayIn <- primitives.RelayInSignals{InSignal: input, LeverSignal: batterySignal}
		}
	}()

	return <-relay.Out()
}

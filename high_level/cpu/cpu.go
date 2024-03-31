package cpu

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type programCounter struct {
	initialized   bool
	memoryAddress int
}

/*type instructionRegister struct {
	instruction
}*/

type accumulator struct {
	value int
}

type RandomAccessMemory []string

type CentralProcessingUnit struct {
	clock
	programCounter
	instructionRegister string
	accumulator

	ram *RandomAccessMemory
}

var errBadCycleIndex = errors.New("bad CPU cycle index")
var errBadInstruction = errors.New("bad CPU instruction")

func NewCentralProcessingUnit(clock clock, programCounter programCounter, instructionRegister string, accumulator accumulator, ram *RandomAccessMemory) *CentralProcessingUnit {
	return &CentralProcessingUnit{clock: clock, programCounter: programCounter, instructionRegister: instructionRegister, accumulator: accumulator, ram: ram}
}

func BuildCPU(clockFrequency time.Duration, ram *RandomAccessMemory) *CentralProcessingUnit {
	return NewCentralProcessingUnit(
		clock{ticker: time.NewTicker(clockFrequency)},
		programCounter{},
		"",
		accumulator{},
		ram,
	)
}

func (cpu *CentralProcessingUnit) Run(stop <-chan struct{}) {
	fetch := func() {
		if !cpu.programCounter.initialized {
			cpu.programCounter.memoryAddress = 0
			cpu.programCounter.initialized = true
		} else {
			if !strings.Contains(cpu.instructionRegister, "JUMP") {
				cpu.programCounter.memoryAddress++
			}
		}

		cpu.instructionRegister = (*cpu.ram)[cpu.programCounter.memoryAddress]
	}

	var execute func()

	decode := func() {
		memValues := strings.Split(cpu.instructionRegister, " ")

		currInstructionMemAddr, err := strconv.Atoi(memValues[1])
		if err != nil {
			panic(err)
		}

		switch memValues[0] {
		case "LOAD":
			execute = func() {
				val, err := strconv.Atoi((*cpu.ram)[currInstructionMemAddr])
				if err != nil {
					panic(err)
				}
				cpu.accumulator.value = val
			}
		case "ADD":
			execute = func() {
				val, err := strconv.Atoi((*cpu.ram)[currInstructionMemAddr])
				if err != nil {
					panic(err)
				}
				cpu.accumulator.value += val
			}
		case "STORE":
			execute = func() {
				(*cpu.ram)[currInstructionMemAddr] = strconv.Itoa(cpu.accumulator.value)
			}
		case "JUMP":
			execute = func() {
				cpu.programCounter.memoryAddress = currInstructionMemAddr
				log.Printf("(*cpu.ram)[6]: %s", (*cpu.ram)[6])
			}
		default:
			panic(fmt.Sprintf("%s: %s", errBadInstruction, memValues[0]))
		}

	}

	i := 0
	for {
		select {
		case <-cpu.tick():
			switch i {
			case 0:
				fetch()
			case 1:
				decode()
			case 2:
				execute()
			default:
				panic(fmt.Sprintf("%s: %d", errBadCycleIndex, i))
			}

			i++

			if i > 2 {
				i = 0
			}
		case <-stop:
			return
		}
	}
}

type clock struct {
	ticker *time.Ticker
}

func (c clock) tick() <-chan time.Time {
	return c.ticker.C
}

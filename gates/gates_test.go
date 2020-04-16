package gates

import (
	"fmt"
	"testing"
)

func TestAND(t *testing.T) {
	tests := []struct {
		input1, input2, wantOut uint8
	}{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := AND(tt.input1, tt.input2); got != tt.wantOut {
				t.Errorf("AND() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}

func TestMultipleAND(t *testing.T) {
	tests := []struct {
		inputs  []uint8
		wantOut uint8
	}{
		{[]uint8{0, 0, 0, 0}, 0},
		{[]uint8{0, 1, 1, 1, 0}, 0},
		{[]uint8{1, 1, 1, 1, 1}, 1},
		{[]uint8{1, 0, 1, 0, 1}, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := MultipleAND(tt.inputs...); got != tt.wantOut {
				t.Errorf("MultipleAND() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}

func TestNAND(t *testing.T) {
	tests := []struct {
		input1, input2, wantOut uint8
	}{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 0, 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := NAND(tt.input1, tt.input2); got != tt.wantOut {
				t.Errorf("NAND() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}

func TestOR(t *testing.T) {
	tests := []struct {
		input1, input2, wantOut uint8
	}{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := OR(tt.input1, tt.input2); got != tt.wantOut {
				t.Errorf("OR() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}

func TestMultipleOR(t *testing.T) {
	tests := []struct {
		inputs  []uint8
		wantOut uint8
	}{
		{[]uint8{0, 1, 1}, 1},
		{[]uint8{0, 1, 1, 0}, 1},
		{[]uint8{0, 0}, 0},
		{[]uint8{0, 0, 1, 0, 0}, 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := MultipleOR(tt.inputs...); got != tt.wantOut {
				t.Errorf("MultipleOR() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}

func TestNOR(t *testing.T) {
	tests := []struct {
		input1, input2, wantOut uint8
	}{
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := NOR(tt.input1, tt.input2); got != tt.wantOut {
				t.Errorf("NOR() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}

func TestXOR(t *testing.T) {
	tests := []struct {
		input1, input2, wantOut uint8
	}{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
		{0, 0, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := XOR(tt.input1, tt.input2); got != tt.wantOut {
				t.Errorf("XOR() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}

func TestNOT(t *testing.T) {
	tests := []struct {
		input, wantOut uint8
	}{
		{1, 0},
		{0, 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := NOT(tt.input); got != tt.wantOut {
				t.Errorf("NOT() = %v, want %v", got, tt.wantOut)
			}
		})
	}
}

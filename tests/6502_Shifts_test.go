package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

func expectUnaffectedFlagsShift(t *testing.T, before, after byte) {
	t.Helper()
	const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB | CPU6502.FlagV
	if before&unaffectedMask != after&unaffectedMask {
		t.Errorf("unaffected flags changed: before=%08b after=%08b", before, after)
	}
}

// Asl Accumulator
func TestASLAccumulatorCanShiftLeftAccumulator(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x20, false, false, false},
		{"ZeroValue", 0x80, true, false, true},
		{"NegativeValue", 0x40, false, true, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x0A // ASL Accumulator

			cpu.Reset()
			cpu.A = tc.val
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val << 1
			if cpu.A != tc.val {
				t.Errorf("cpu.A = 0x%02X, want 0x%02X", cpu.A, tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// Asl Zero Page
func TestASLZeroPageCanShiftLeftMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x20, false, false, false},
		{"ZeroValue", 0x80, true, false, true},
		{"NegativeValue", 0x40, false, true, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x010
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x06 // ASL Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val << 1
			if bus.RAM[ADDR] != tc.val {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// Asl Zero Page X
func TestASLZeroPageXCanShiftLeftMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x20, false, false, false},
		{"ZeroValue", 0x80, true, false, true},
		{"NegativeValue", 0x40, false, true, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x014
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x16 // ASL Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.X = 0x04
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val << 1
			if bus.RAM[ADDR] != tc.val {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// Asl Absolute
func TestASLAbsoluteCanShiftLeftMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x20, false, false, false},
		{"ZeroValue", 0x80, true, false, true},
		{"NegativeValue", 0x40, false, true, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52F0
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x0E // ASL Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val << 1
			if bus.RAM[ADDR] != tc.val {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// Asl Absolute X
func TestASLAbsoluteXCanShiftLeftMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x20, false, false, false},
		{"ZeroValue", 0x80, true, false, true},
		{"NegativeValue", 0x40, false, true, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52FC
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x1E // ASL Absolute X
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.X = 0x0C
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val << 1
			if bus.RAM[ADDR] != tc.val {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// LSR Accumulator
func TestLSRAccumulatorCanShiftRightAccumulator(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x41, false, false, true},
		{"ZeroValue", 0x01, true, false, true},
		{"NegativeValue", 0x41, false, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x4A // LSR Accumulator

			cpu.Reset()
			cpu.A = tc.val
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val >> 1
			if cpu.A != tc.val {
				t.Errorf("cpu.A = 0x%02X, want 0x%02X", cpu.A, tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// LSR Zero Page
func TestLSRZeroPageCanShiftRightMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x41, false, false, true},
		{"ZeroValue", 0x01, true, false, true},
		{"NegativeValue", 0x41, false, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x010
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x46 // LSR Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val >> 1
			if bus.RAM[ADDR] != tc.val {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// LSR Zero Page X
func TestLSRZeroPageXCanShiftRightMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x41, false, false, true},
		{"ZeroValue", 0x01, true, false, true},
		{"NegativeValue", 0x41, false, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x014
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x56 // LSR Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.X = 0x04
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val >> 1
			if bus.RAM[ADDR] != tc.val {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// LSR Absolute
func TestLSRAbsoluteCanShiftRightMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x41, false, false, true},
		{"ZeroValue", 0x01, true, false, true},
		{"NegativeValue", 0x41, false, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52F0
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x4E // LSR Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val >> 1
			if bus.RAM[ADDR] != tc.val {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// LSR Absolute X
func TestLSRAbsoluteXCanShiftRightMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNonNegativeValue", 0x41, false, false, true},
		{"ZeroValue", 0x01, true, false, true},
		{"NegativeValue", 0x41, false, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52FC
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x5E // LSR Absolute X
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.X = 0x0C
			cpu.SetFlag(CPU6502.FlagC, !tc.wantC)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.val = tc.val >> 1
			if bus.RAM[ADDR] != tc.val {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.val)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("Z flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROL Accumulator
func TestROLAccumulatorCanRotateLeftAccumulator(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantA               byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x20, false, 0x40, false, false, false},
		{"CarryInSetsBit0", 0x20, true, 0x41, false, false, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x01, false, false, false},
		{"NegativeResult", 0x40, false, 0x80, false, true, false},
		{"CarryOut", 0x80, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x81, true, 0x03, false, false, true},
		{"CarryOutNegativeResult", 0xC0, true, 0x81, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x2A // ROL Accumulator

			cpu.Reset()
			cpu.A = tc.val
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROL Zero Page
func TestROLZeroPageCanRotateLeftMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantVal             byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x20, false, 0x40, false, false, false},
		{"CarryInSetsBit0", 0x20, true, 0x41, false, false, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x01, false, false, false},
		{"NegativeResult", 0x40, false, 0x80, false, true, false},
		{"CarryOut", 0x80, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x81, true, 0x03, false, false, true},
		{"CarryOutNegativeResult", 0xC0, true, 0x81, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x010
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x26 // ROL Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			if bus.RAM[ADDR] != tc.wantVal {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.wantVal)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROL Zero Page X
func TestROLZeroPageXCanRotateLeftValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantVal             byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x20, false, 0x40, false, false, false},
		{"CarryInSetsBit0", 0x20, true, 0x41, false, false, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x01, false, false, false},
		{"NegativeResult", 0x40, false, 0x80, false, true, false},
		{"CarryOut", 0x80, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x81, true, 0x03, false, false, true},
		{"CarryOutNegativeResult", 0xC0, true, 0x81, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x014
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x36 // ROL Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.X = 0x04
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			if bus.RAM[ADDR] != tc.wantVal {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.wantVal)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROL Absolute
func TestROLAbsoluteCanRotateLeftMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantVal             byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x20, false, 0x40, false, false, false},
		{"CarryInSetsBit0", 0x20, true, 0x41, false, false, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x01, false, false, false},
		{"NegativeResult", 0x40, false, 0x80, false, true, false},
		{"CarryOut", 0x80, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x81, true, 0x03, false, false, true},
		{"CarryOutNegativeResult", 0xC0, true, 0x81, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52F0
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x2E // ROL Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if bus.RAM[ADDR] != tc.wantVal {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.wantVal)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROL Absolute X
func TestROLAbsoluteXCanRotateLeftMemoryValue(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantVal             byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x20, false, 0x40, false, false, false},
		{"CarryInSetsBit0", 0x20, true, 0x41, false, false, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x01, false, false, false},
		{"NegativeResult", 0x40, false, 0x80, false, true, false},
		{"CarryOut", 0x80, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x81, true, 0x03, false, false, true},
		{"CarryOutNegativeResult", 0xC0, true, 0x81, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52F0
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x3E // ROL Absolute X
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 15 {
				cpu.Clock()
			}

			// then:
			if bus.RAM[ADDR] != tc.wantVal {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.wantVal)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROR Accumulator
func TestRORAccumulatorCanRotateRightAccumulator(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantA               byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x04, false, 0x02, false, false, false},
		{"CarryInSetsBit7", 0x04, true, 0x82, false, true, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x80, false, true, false},
		{"PositiveResult", 0x40, false, 0x20, false, false, false},
		{"CarryOut", 0x01, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x03, true, 0x81, false, true, true},
		{"CarryOutNegativeResult", 0x81, true, 0xC0, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x6A // ROR Accumulator

			cpu.Reset()
			cpu.A = tc.val
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROR Zero Page
func TestRORZeroPageCanRotateRightAccumulator(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantVal             byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x04, false, 0x02, false, false, false},
		{"CarryInSetsBit7", 0x04, true, 0x82, false, true, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x80, false, true, false},
		{"PositiveResult", 0x40, false, 0x20, false, false, false},
		{"CarryOut", 0x01, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x03, true, 0x81, false, true, true},
		{"CarryOutNegativeResult", 0x81, true, 0xC0, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x010
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x66 // ROR Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.A = 0xAA
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			if bus.RAM[ADDR] != tc.wantVal {
				t.Errorf("RAM[0x%04X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.wantVal)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROR Zero Page X
func TestRORZeroPageXCanRotateRightAccumulator(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantVal             byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x04, false, 0x02, false, false, false},
		{"CarryInSetsBit7", 0x04, true, 0x82, false, true, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x80, false, true, false},
		{"PositiveResult", 0x40, false, 0x20, false, false, false},
		{"CarryOut", 0x01, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x03, true, 0x81, false, true, true},
		{"CarryOutNegativeResult", 0x81, true, 0xC0, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x014
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x76 // ROR Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.X = 0x04
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if bus.RAM[ADDR] != tc.wantVal {
				t.Errorf("RAM[0x%04X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.wantVal)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROR Absolute
func TestRORAbsoluteCanRotateRightAccumulator(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantVal             byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x04, false, 0x02, false, false, false},
		{"CarryInSetsBit7", 0x04, true, 0x82, false, true, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x80, false, true, false},
		{"PositiveResult", 0x40, false, 0x20, false, false, false},
		{"CarryOut", 0x01, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x03, true, 0x81, false, true, true},
		{"CarryOutNegativeResult", 0x81, true, 0xC0, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52F0
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x6E // ROR Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if bus.RAM[ADDR] != tc.wantVal {
				t.Errorf("RAM[0x%04X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.wantVal)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

// ROR Absolute X
func TestRORAbsoluteXCanRotateRightAccumulator(t *testing.T) {
	tests := []struct {
		name                string
		val                 byte
		carryIn             bool
		wantVal             byte
		wantZ, wantN, wantC bool
	}{
		{"NonZeroNoCarry", 0x04, false, 0x02, false, false, false},
		{"CarryInSetsBit7", 0x04, true, 0x82, false, true, false},
		{"ZeroValue", 0x00, false, 0x00, true, false, false},
		{"CarryInOnly", 0x00, true, 0x80, false, true, false},
		{"PositiveResult", 0x40, false, 0x20, false, false, false},
		{"CarryOut", 0x01, false, 0x00, true, false, true},
		{"CarryOutAndIn", 0x03, true, 0x81, false, true, true},
		{"CarryOutNegativeResult", 0x81, true, 0xC0, false, true, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52FC
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x7E // ROR Absolute X
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.val

			cpu.Reset()
			cpu.X = 0x0C
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if bus.RAM[ADDR] != tc.wantVal {
				t.Errorf("RAM[0x%04X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.wantVal)
			}

			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
			}

			expectUnaffectedFlagsShift(t, statusBefore, cpu.Status)
		})
	}
}

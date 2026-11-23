package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// Inc Zero Page
func TestINCZeroPageCanIncrementMemoryValue(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x0010
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xE6 // Inc Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = tc.valMem

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			tc.valMem++
			if bus.RAM[ADDR] != tc.valMem {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.valMem)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// Inc Zero Page X
func TestINCZeroPageXCanIncrementMemoryValue(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x0013
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xF6 // Inc Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.valMem

			cpu.Reset()
			cpu.X = 0x03
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.valMem++
			if bus.RAM[ADDR] != tc.valMem {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.valMem)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// Inc Absolute
func TestINCAbsoluteCanIncrementMemoryValue(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
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
			bus.RAM[0x0000] = 0xEE // Inc Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.valMem

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.valMem++
			if bus.RAM[ADDR] != tc.valMem {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.valMem)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// Inc Absolute X
func TestINCAbsoluteXCanIncrementMemoryValue(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52F5
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xFE // Inc Absolute X
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.valMem

			cpu.Reset()
			cpu.X = 0x05
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 15 {
				cpu.Clock()
			}

			// then:
			tc.valMem++
			if bus.RAM[ADDR] != tc.valMem {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.valMem)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// INX
func TestINXCanIncrementXRegister(t *testing.T) {
	tests := []struct {
		name         string
		val          byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xE8 // INX Implied

			cpu.Reset()
			cpu.X = tc.val
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			tc.val++
			if cpu.X != tc.val {
				t.Errorf("cpu.X = 0x%02X, want 0x%02X", cpu.X, tc.val)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// INY
func TestINYCanIncrementYRegister(t *testing.T) {
	tests := []struct {
		name         string
		val          byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xC8 // INY Implied

			cpu.Reset()
			cpu.Y = tc.val
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			tc.val++
			if cpu.Y != tc.val {
				t.Errorf("cpu.Y = 0x%02X, want 0x%02X", cpu.Y, tc.val)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// Dec Zero Page
func TestDecZeroPageCanDecrementMemoryValue(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0x01, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x0010
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xC6 // Dec Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = tc.valMem

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			tc.valMem--
			if bus.RAM[ADDR] != tc.valMem {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.valMem)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// Dec Zero Page X
func TestDECZeroPageXCanDecrementMemoryValue(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0x01, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x0013
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xD6 // Dec Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[ADDR] = tc.valMem

			cpu.Reset()
			cpu.X = 0x03
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.valMem--
			if bus.RAM[ADDR] != tc.valMem {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.valMem)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// Dec Absolute
func TestDECAbsoluteCanDecrementMemoryValue(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0x01, true, false},
		{"NegativeValue", 0xA9, false, true},
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
			bus.RAM[0x0000] = 0xCE // Dec Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.valMem

			cpu.Reset()
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			tc.valMem--
			if bus.RAM[ADDR] != tc.valMem {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.valMem)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// Dec Absolute X
func TestDECAbsoluteXCanDecrementMemoryValue(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			const ADDR = 0x52F5
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xDE // Dec Absolute X
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[ADDR] = tc.valMem

			cpu.Reset()
			cpu.X = 0x05
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 15 {
				cpu.Clock()
			}

			// then:
			tc.valMem--
			if bus.RAM[ADDR] != tc.valMem {
				t.Errorf("bus.RAM[0x%02X] = 0x%02X, want 0x%02X", ADDR, bus.RAM[ADDR], tc.valMem)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// DEX
func TestDEXCanDecreaseXRegister(t *testing.T) {
	tests := []struct {
		name         string
		val          byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0xCA // DEX Implied

			cpu.Reset()
			cpu.X = tc.val
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			tc.val--
			if cpu.X != tc.val {
				t.Errorf("cpu.X = 0x%02X, want 0x%02X", cpu.X, tc.val)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

// DEY
func TestDEYCanDecreaseYRegister(t *testing.T) {
	tests := []struct {
		name         string
		val          byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x08, false, false},
		{"ZeroValue", 0xFF, true, false},
		{"NegativeValue", 0xA9, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x88 // DEY Implied

			cpu.Reset()
			cpu.Y = tc.val
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			tc.val++
			if cpu.Y != tc.val {
				t.Errorf("cpu.Y = 0x%02X, want 0x%02X", cpu.Y, tc.val)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			expectUnaffectedFlags(t, statusBefore, cpu.Status)
		})
	}
}

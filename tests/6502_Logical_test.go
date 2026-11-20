package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// And Immediate
func TestANDImmediateCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x0F, 0x2C, 0x0F & 0x2C, false, false},
		{"ZeroValue", 0x00, 0xFF, 0x00 & 0xFF, true, false},
		{"NegativeValue", 0xA9, 0x83, 0xA9 & 0x83, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x29 // And Immidiate
			bus.RAM[0x0001] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
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

// And Zero Page
func TestANDZeroPageCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x0F, 0x2C, 0x0F & 0x2C, false, false},
		{"ZeroValue", 0x00, 0xFF, 0x00 & 0xFF, true, false},
		{"NegativeValue", 0xA9, 0x83, 0xA9 & 0x83, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x25 // And Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 11 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// And Zero Page X
func TestANDZeroPageXCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x0F, 0x2C, 0x0F & 0x2C, false, false},
		{"ZeroValue", 0x00, 0xFF, 0x00 & 0xFF, true, false},
		{"NegativeValue", 0xA9, 0x83, 0xA9 & 0x83, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x35 // And Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0014] = tc.valMem

			cpu.Reset()
			cpu.X = 0x04
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 11 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// And Absolute
func TestANDbsoluteCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x0F, 0x2C, 0x0F & 0x2C, false, false},
		{"ZeroValue", 0x00, 0xFF, 0x00 & 0xFF, true, false},
		{"NegativeValue", 0xA9, 0x83, 0xA9 & 0x83, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x2D // And Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 12 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// And Absolute X
func TestANDAbsoluteXCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", false, 0x0F, 0x2C, 0x0F & 0x2C, false, false},
		{"ZeroValue", false, 0x00, 0xFF, 0x00 & 0xFF, true, false},
		{"NegativeValue", false, 0xA9, 0x83, 0xA9 & 0x83, false, true},
		{"PageCross", true, 0x0F, 0x2C, 0x0F & 0x2C, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x3D // And Absolute X
				bus.RAM[0x0001] = 0xFF
				bus.RAM[0x0002] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x3D // And Absolute X
				bus.RAM[0x0001] = 0xF0
				bus.RAM[0x0002] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.X = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 12

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// And Absolute Y
func TestANDAbsoluteYCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", false, 0x0F, 0x2C, 0x0F & 0x2C, false, false},
		{"ZeroValue", false, 0x00, 0xFF, 0x00 & 0xFF, true, false},
		{"NegativeValue", false, 0xA9, 0x83, 0xA9 & 0x83, false, true},
		{"PageCross", true, 0x0F, 0x2C, 0x0F & 0x2C, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x39 // And Absolute Y
				bus.RAM[0x0001] = 0xFF
				bus.RAM[0x0002] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x39 // And Absolute Y
				bus.RAM[0x0001] = 0xF0
				bus.RAM[0x0002] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.Y = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 12

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// And Indirect X
func TestANDAIndirectXCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x0F, 0x2C, 0x0F & 0x2C, false, false},
		{"ZeroValue", 0x00, 0xFF, 0x00 & 0xFF, true, false},
		{"NegativeValue", 0xA9, 0x83, 0xA9 & 0x83, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x21 // And Indirect X
			bus.RAM[0x0001] = 0x02
			bus.RAM[0x0006] = 0xF0
			bus.RAM[0x0007] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.X = 0x04
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// And Indirect Y
func TestANDIndirectYCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", false, 0x0F, 0x2C, 0x0F & 0x2C, false, false},
		{"ZeroValue", false, 0x00, 0xFF, 0x00 & 0xFF, true, false},
		{"NegativeValue", false, 0xA9, 0x83, 0xA9 & 0x83, false, true},
		{"PageCross", true, 0x0F, 0x2C, 0x0F & 0x2C, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x31 // And Indirect Y
				bus.RAM[0x0001] = 0x10
				bus.RAM[0x0010] = 0xFF
				bus.RAM[0x0011] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x31 // And Indirect Y
				bus.RAM[0x0001] = 0x10
				bus.RAM[0x0010] = 0xF0
				bus.RAM[0x0011] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.Y = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 12

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// EOR Immediate
func TestEORImmediateCanLogicalXOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x4F, 0x78, 0x4F ^ 0x78, false, false},
		{"ZeroValue", 0xFF, 0xFF, 0xFF ^ 0xFF, true, false},
		{"NegativeValue", 0x9F, 0x79, 0x9F ^ 0x79, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x49 // EOR Immidiate
			bus.RAM[0x0001] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
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

// EOR Zero Page
func TestEORZeroPageCanLogicalXOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x4F, 0x78, 0x4F ^ 0x78, false, false},
		{"ZeroValue", 0xFF, 0xFF, 0xFF ^ 0xFF, true, false},
		{"NegativeValue", 0x9F, 0x79, 0x9F ^ 0x79, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x45 // EOR Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 11 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// EOR Zero Page X
func TestEORZeroPageXCanLogicalXOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x4F, 0x78, 0x4F ^ 0x78, false, false},
		{"ZeroValue", 0xFF, 0xFF, 0xFF ^ 0xFF, true, false},
		{"NegativeValue", 0x9F, 0x79, 0x9F ^ 0x79, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x55 // EOR Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0014] = tc.valMem

			cpu.Reset()
			cpu.X = 0x04
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 11 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// EOR Absolute
func TestEORAbsoluteCanLogicalXOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x4F, 0x78, 0x4F ^ 0x78, false, false},
		{"ZeroValue", 0xFF, 0xFF, 0xFF ^ 0xFF, true, false},
		{"NegativeValue", 0x9F, 0x79, 0x9F ^ 0x79, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x4D // EOR Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 12 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// EOR Absolute X
func TestEORAbsoluteXCanLogicalXOR(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", false, 0x4F, 0x78, 0x4F ^ 0x78, false, false},
		{"ZeroValue", false, 0xFF, 0xFF, 0xFF ^ 0xFF, true, false},
		{"NegativeValue", false, 0x9F, 0x79, 0x9F ^ 0x79, false, true},
		{"PageCross", true, 0x4F, 0x78, 0x4F ^ 0x78, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x5D // EOR Absolute X
				bus.RAM[0x0001] = 0xFF
				bus.RAM[0x0002] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x5D // EOR Absolute X
				bus.RAM[0x0001] = 0xF0
				bus.RAM[0x0002] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.X = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 12

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// EOR Absolute Y
func TestEORAbsoluteYCanLogicalXOR(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", false, 0x4F, 0x78, 0x4F ^ 0x78, false, false},
		{"ZeroValue", false, 0xFF, 0xFF, 0xFF ^ 0xFF, true, false},
		{"NegativeValue", false, 0x9F, 0x79, 0x9F ^ 0x79, false, true},
		{"PageCross", true, 0x4F, 0x78, 0x4F ^ 0x78, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x59 // And Absolute Y
				bus.RAM[0x0001] = 0xFF
				bus.RAM[0x0002] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x59 // And Absolute Y
				bus.RAM[0x0001] = 0xF0
				bus.RAM[0x0002] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.Y = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 12

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// EOR Indirect X
func TestEORIndirectXCanLogicalXOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x4F, 0x78, 0x4F ^ 0x78, false, false},
		{"ZeroValue", 0xFF, 0xFF, 0xFF ^ 0xFF, true, false},
		{"NegativeValue", 0x9F, 0x79, 0x9F ^ 0x79, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x41 // EOR Indirect X
			bus.RAM[0x0001] = 0x02
			bus.RAM[0x0006] = 0xF0
			bus.RAM[0x0007] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.X = 0x04
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// EOR Indirect Y
func TestEORIndirectYCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", false, 0x4F, 0x78, 0x4F ^ 0x78, false, false},
		{"ZeroValue", false, 0xFF, 0xFF, 0xFF ^ 0xFF, true, false},
		{"NegativeValue", false, 0x9F, 0x79, 0x9F ^ 0x79, false, true},
		{"PageCross", true, 0x4F, 0x78, 0x4F ^ 0x78, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x51 // EOR Indirect Y
				bus.RAM[0x0001] = 0x10
				bus.RAM[0x0010] = 0xFF
				bus.RAM[0x0011] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x51 // EOR Indirect Y
				bus.RAM[0x0001] = 0x10
				bus.RAM[0x0010] = 0xF0
				bus.RAM[0x0011] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.Y = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 13

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// ORA Immediate
func TestORAImmediateCanLogicalOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x3F, 0x30, 0x3F | 0x30, false, false},
		{"ZeroValue", 0x00, 0x00, 0x00 | 0x00, true, false},
		{"NegativeValue", 0xF0, 0x0F, 0xF0 | 0x0F, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x09 // ORA Immidiate
			bus.RAM[0x0001] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
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

// ORA Zero Page
func TestORAZeroPageCanLogicalOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x3F, 0x30, 0x3F | 0x30, false, false},
		{"ZeroValue", 0x00, 0x00, 0x00 | 0x00, true, false},
		{"NegativeValue", 0xF0, 0x0F, 0xF0 | 0x0F, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x05 // ORA Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 11 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// ORA Zero Page X
func TestORAZeroPageXCanLogicalOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x3F, 0x30, 0x3F | 0x30, false, false},
		{"ZeroValue", 0x00, 0x00, 0x00 | 0x00, true, false},
		{"NegativeValue", 0xF0, 0x0F, 0xF0 | 0x0F, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x15 // ORA Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0014] = tc.valMem

			cpu.Reset()
			cpu.X = 0x04
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 11 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// EOR Absolute
func TestORAAbsoluteCanLogicalOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x3F, 0x30, 0x3F | 0x30, false, false},
		{"ZeroValue", 0x00, 0x00, 0x00 | 0x00, true, false},
		{"NegativeValue", 0xF0, 0x0F, 0xF0 | 0x0F, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x0D // EOR Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 12 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// ORA Absolute X
func TestORAAbsoluteXCanLogicalOR(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", true, 0x3F, 0x30, 0x3F | 0x30, false, false},
		{"ZeroValue", true, 0x00, 0x00, 0x00 | 0x00, true, false},
		{"NegativeValue", true, 0xF0, 0x0F, 0xF0 | 0x0F, false, true},
		{"PageCross", true, 0x3F, 0x30, 0x3F | 0x30, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x1D // ORA Absolute X
				bus.RAM[0x0001] = 0xFF
				bus.RAM[0x0002] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x1D // ORA Absolute X
				bus.RAM[0x0001] = 0xF0
				bus.RAM[0x0002] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.X = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 12

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// ORA Absolute Y
func TestORAAbsoluteYCanLogicalOR(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", true, 0x3F, 0x30, 0x3F | 0x30, false, false},
		{"ZeroValue", true, 0x00, 0x00, 0x00 | 0x00, true, false},
		{"NegativeValue", true, 0xF0, 0x0F, 0xF0 | 0x0F, false, true},
		{"PageCross", true, 0x3F, 0x30, 0x3F | 0x30, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x19 // ORA Absolute Y
				bus.RAM[0x0001] = 0xFF
				bus.RAM[0x0002] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x19 // ORA Absolute Y
				bus.RAM[0x0001] = 0xF0
				bus.RAM[0x0002] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.Y = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 12

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// ORA Indirect X
func TestORAIndirectXCanLogicalXOR(t *testing.T) {
	tests := []struct {
		name         string
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x3F, 0x30, 0x3F | 0x30, false, false},
		{"ZeroValue", 0x00, 0x00, 0x00 | 0x00, true, false},
		{"NegativeValue", 0xF0, 0x0F, 0xF0 | 0x0F, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x01 // ORA Indirect X
			bus.RAM[0x0001] = 0x02
			bus.RAM[0x0006] = 0xF0
			bus.RAM[0x0007] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.X = 0x04
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// ORA Indirect Y
func TestORAIndirectYCanLogicalAnd(t *testing.T) {
	tests := []struct {
		name         string
		crossPage    bool
		valMem       byte
		valReg       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", true, 0x3F, 0x30, 0x3F | 0x30, false, false},
		{"ZeroValue", true, 0x00, 0x00, 0x00 | 0x00, true, false},
		{"NegativeValue", true, 0xF0, 0x0F, 0xF0 | 0x0F, false, true},
		{"PageCross", true, 0x3F, 0x30, 0x3F | 0x30, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)
			if tc.crossPage {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x11 // ORA Indirect Y
				bus.RAM[0x0001] = 0x10
				bus.RAM[0x0010] = 0xFF
				bus.RAM[0x0011] = 0x01
				bus.RAM[0x020B] = tc.valMem
			} else {
				bus.RAM[0xFFFC] = 0x00
				bus.RAM[0xFFFD] = 0x00
				bus.RAM[0x0000] = 0x11 // ORA Indirect Y
				bus.RAM[0x0001] = 0x10
				bus.RAM[0x0010] = 0xF0
				bus.RAM[0x0011] = 0x52
				bus.RAM[0x52FC] = tc.valMem
			}

			cpu.Reset()
			cpu.Y = 0x0C
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			statusBefore := cpu.Status

			cycle_count := 13

			if tc.crossPage {
				cycle_count++
			}

			// when:
			for range cycle_count {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
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

// Bit test Zero Page
func TestBITZeroPageCanLogicalBitTest(t *testing.T) {
	tests := []struct {
		name                string
		valMem              byte
		valReg              byte
		wantZ, wantN, wantV bool
	}{
		{"NonZeroResult", 0x0F, 0x0F, false, false, false},
		{"ZeroResult", 0x00, 0xFF, true, false, false},
		{"NegativeFromMemory", 0x80, 0xFF, false, true, false},
		{"OverflowFromMemory", 0x40, 0xFF, false, false, true},
		{"ZeroNegativeAndOverflow", 0xC0, 0x01, true, true, true},
		{"NandVComeFromMemoryNotFromA", 0x3F, 0xC0, true, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x24 // BIT Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg

			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 11 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.valReg {
				t.Errorf("A = 0x%02X, want unchanged 0x%02X", cpu.A, tc.valReg)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			if got := cpu.Status&CPU6502.FlagV != 0; got != tc.wantV {
				t.Errorf("V flag = %v, want %v", got, tc.wantV)
			}

			const unaffectedMask = CPU6502.FlagC | CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// Bit test Absolute
func TestBITAbsoluteCanLogicalBitTest(t *testing.T) {
	tests := []struct {
		name                string
		valMem              byte
		valReg              byte
		wantZ, wantN, wantV bool
	}{
		{"NonZeroResult", 0x0F, 0x0F, false, false, false},
		{"ZeroResult", 0x00, 0xFF, true, false, false},
		{"NegativeFromMemory", 0x80, 0xFF, false, true, false},
		{"OverflowFromMemory", 0x40, 0xFF, false, false, true},
		{"ZeroNegativeAndOverflow", 0xC0, 0x01, true, true, true},
		{"NandVComeFromMemoryNotFromA", 0x3F, 0xC0, true, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x2C // BIT Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg

			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 12 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.valReg {
				t.Errorf("A = 0x%02X, want unchanged 0x%02X", cpu.A, tc.valReg)
			}

			if got := cpu.Status&CPU6502.FlagZ != 0; got != tc.wantZ {
				t.Errorf("Z flag = %v, want %v", got, tc.wantZ)
			}

			if got := cpu.Status&CPU6502.FlagN != 0; got != tc.wantN {
				t.Errorf("N flag = %v, want %v", got, tc.wantN)
			}

			if got := cpu.Status&CPU6502.FlagV != 0; got != tc.wantV {
				t.Errorf("V flag = %v, want %v", got, tc.wantV)
			}

			const unaffectedMask = CPU6502.FlagC | CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

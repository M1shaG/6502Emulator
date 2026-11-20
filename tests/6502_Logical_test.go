package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// ORA Immediate
func TestORAImmediateCanLogicalORANonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x09 // ORA Immediate
	bus.RAM[0x0001] = 0x30

	cpu.Reset()
	cpu.A = 0x3F
	statusBefore := cpu.Status

	// when:
	for range 10 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x3F | 0x30
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestORAImmediateCanLogicalORAZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x09 // ORA Immediate
	bus.RAM[0x0001] = 0x00

	cpu.Reset()
	cpu.A = 0x00
	statusBefore := cpu.Status

	// when:
	for range 10 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x00 | 0x00
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestORAImmediateCanLogicalORANegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x09 // ORA Immediate
	bus.RAM[0x0001] = 0xF0

	cpu.Reset()
	cpu.A = 0x0F
	statusBefore := cpu.Status

	// when:
	for range 10 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0xF0 | 0x0F
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// ORA Zero Page
func TestORAZeroPageCanLogicalORANonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x05 // ORA Zero Page
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0x30

	cpu.Reset()
	cpu.A = 0x3F
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x30 | 0x3F
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestORAZeroPageCanLogicalORAZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x05 // ORA Zero Page
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0x00
	cpu.Reset()
	cpu.A = 0x00
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x00 | 0x00
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestORAZeroPageCanLogicalORANegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x05 // ORA Zero Page
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xF0

	cpu.Reset()
	cpu.A = 0x0F
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0xF0 | 0x0F
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// ORA Zero Page X
func TestORAZeroPageXCanLogicalAndNonZeroORANegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x15 // ORA Zero Page X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0014] = 0x30

	cpu.Reset()
	cpu.X = 0x04
	cpu.A = 0x3F
	statusBefore := cpu.Status

	// when:
	for range 12 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x30 | 0x3F
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestORAZeroPageXCanLogicalORAZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x15 // ORA Zero Page X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0014] = 0x00
	cpu.Reset()
	cpu.A = 0x00
	statusBefore := cpu.Status

	// when:
	for range 12 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x00 | 0x00
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestORAZeroPageXCanLogicalORANegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x15 // ORA Zero Page X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0014] = 0xF0

	cpu.Reset()
	cpu.X = 0x04
	cpu.A = 0x0F
	statusBefore := cpu.Status

	// when:
	for range 12 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0xF0 | 0x0F
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// ORA Absolute
func TestORAAbsoluteCanLogicalORANonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x0D // ORA Absolute
	bus.RAM[0x0001] = 0x52
	bus.RAM[0x0002] = 0xF0
	bus.RAM[0x52F0] = 0x30

	cpu.Reset()
	cpu.A = 0x3F
	statusBefore := cpu.Status

	// when:
	for range 12 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x30 | 0x3F
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestORAAbsoluteCanLogicalORAZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x0D // ORA Absolute
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52F0] = 0x00
	cpu.Reset()
	cpu.A = 0x00
	statusBefore := cpu.Status

	// when:
	for range 12 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x00 | 0x00
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestORAAbsoluteCanLogicalORANegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x0D // ORA Negative
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52F0] = 0x0F

	cpu.Reset()
	cpu.A = 0xF0
	statusBefore := cpu.Status

	// when:
	for range 12 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0x0F | 0xF0
	if cpu.A != ORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, ORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// ORA Absolute
func TestORAAbsoluteCanLogicalOR(t *testing.T) {
	tests := []struct {
		name         string
		valReg       byte
		valMem       byte
		wantA        byte
		wantZ, wantN bool
	}{
		{"NonZeroNonNegativeValue", 0x0F, 0x30, 0x3F, false, false},
		{"ZeroValue", 0x00, 0x00, 0x00, true, false},
		{"NegativeValue", 0x0F, 0xF0, 0xFF, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x0D // ORA Absolute
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

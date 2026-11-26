package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// BCC Relative
func TestBCCCanBranchForwardsWhenCarryFlagIsNotSet(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x90 // BCC Relative
	bus.RAM[0x0001] = 0x09

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagC, false)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0B {
		t.Errorf("cpu.PC = %#02X, want 0x0X0B", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// BCS Relative
func TestBCSCanBranchForwardsWhenCarryFlagIsSet(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0xB0 // BCS Relative
	bus.RAM[0x0001] = 0x09

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagC, true)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0B {
		t.Errorf("cpu.PC = %#02X, want 0x0X0B", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// BNE Relative
func TestBEQCanBranchForwardsWhenZeroFlagIsNotSet(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0xD0 // BNE Relative
	bus.RAM[0x0001] = 0x09

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagZ, false)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0B {
		t.Errorf("cpu.PC = %#02X, want 0x0X0B", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// BEQ Relative
func TestBEQCanBranchForwardsWhenCarryFlagIsSet(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0xF0 // BEQ Relative
	bus.RAM[0x0001] = 0x09

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagZ, true)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0B {
		t.Errorf("cpu.PC = %#02X, want 0x0X0B", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// BPL Relative
func TestBPLCanBranchForwardsWhenNegativeFlagIsNotSet(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x10 // BPL Relative
	bus.RAM[0x0001] = 0x09

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagN, false)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0B {
		t.Errorf("cpu.PC = %#02X, want 0x0X0B", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// BMI Relative
func TestBEQCanBranchForwardsWhenNegativeFlagIsSet(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x30 // BMI Relative
	bus.RAM[0x0001] = 0x09

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagN, true)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0B {
		t.Errorf("cpu.PC = %#02X, want 0x0X0B", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// BVC Relative
func TestBVCCanBranchForwardsWhenOverflowFlagIsNotSet(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x50 // BVC Relative
	bus.RAM[0x0001] = 0x09

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagV, false)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0B {
		t.Errorf("cpu.PC = %#02X, want 0x0X0B", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// BVS Relative
func TestBVSCanBranchForwardsWhenNegativeFlagIsSet(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x70 // BVS Relative
	bus.RAM[0x0001] = 0x09

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagV, true)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0B {
		t.Errorf("cpu.PC = %#02X, want 0x0X0B", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// New Page
func TestCanBranchForwardsWhenNewPage(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x70 // BVS Relative
	bus.RAM[0x0001] = 0xFF

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagV, true)
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x01 {
		t.Errorf("cpu.PC = %#02X, want 0x01", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

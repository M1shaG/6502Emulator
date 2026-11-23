package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

func expectUnaffectedFlagsAll(t *testing.T, before, after byte) {
	t.Helper()
	const unaffectedMask = CPU6502.FlagC | CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB | CPU6502.FlagV | CPU6502.FlagZ | CPU6502.FlagN
	if before&unaffectedMask != after&unaffectedMask {
		t.Errorf("unaffected flags changed: before=%08b after=%08b", before, after)
	}
}

// JMP Absolute
func TestJMPAbsoluteCanJumpToAnNewLocationInTheProgram(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x4C // JMP Absolute
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52

	cpu.Reset()
	statusBefore := cpu.Status

	// when:
	for range 11 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x52F0 {
		t.Errorf("cpu.PC = %#02x, want 0x52F0", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// JMP Indirect
func TestJMPIndirectCanJumpToAnNewLocationInTheProgram(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x6C
	bus.RAM[0x0001] = 0xFF
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FF] = 0x23
	bus.RAM[0x5200] = 0xFF
	bus.RAM[0x5300] = 0x11

	cpu.Reset()
	statusBefore := cpu.Status

	// when:
	for range 13 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0xFF23 {
		t.Errorf("cpu.PC = %#02x, want 0xFF23", cpu.PC)
	}
	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

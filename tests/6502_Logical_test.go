package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

func TestLogicalAnd(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x29 // AND Immediate
	bus.RAM[0x0001] = 0x2C

	cpu.Reset()
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	// then:
	if cpu.A != 0x29&0x2C {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, 0x29&0x2C)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

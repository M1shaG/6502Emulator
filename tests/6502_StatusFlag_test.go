package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// SEC, SED, SEI
func TestSetFlagInstructionsCanSetTheirFlag(t *testing.T) {
	tests := []struct {
		name        string
		flag        byte
		instruction byte
	}{
		{"SEC", CPU6502.FlagC, 0x38},
		{"SED", CPU6502.FlagD, 0xF8},
		{"SEI", CPU6502.FlagI, 0x78},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = tc.instruction

			cpu.Reset()
			cpu.SetFlag(tc.flag, false)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			if !cpu.GetFlag(tc.flag) {
				t.Errorf("flag 0x%02X should be set", tc.flag)
			}

			if (cpu.Status &^ tc.flag) != (statusBefore &^ tc.flag) {
				t.Errorf("other flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// CLC, CLD, CLI, CLV
func TestClearFlagInstructionsCanClearTheirFlag(t *testing.T) {
	tests := []struct {
		name        string
		flag        byte
		instruction byte
	}{
		{"CLC", CPU6502.FlagC, 0x18},
		{"CLD", CPU6502.FlagD, 0xD8},
		{"CLI", CPU6502.FlagI, 0x58},
		{"CLV", CPU6502.FlagV, 0xB8},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = tc.instruction

			cpu.Reset()
			cpu.SetFlag(tc.flag, true)
			statusBefore := cpu.Status

			// when:
			for range 10 { // 8 reset + 2 CLx
				cpu.Clock()
			}

			// then:
			if cpu.GetFlag(tc.flag) {
				t.Errorf("flag 0x%02X should be clear", tc.flag)
			}

			if (cpu.Status &^ tc.flag) != (statusBefore &^ tc.flag) {
				t.Errorf("other flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

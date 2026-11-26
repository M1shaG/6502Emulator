package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// BRK
func TestBRKPushesReturnAddressAndStatusAndJumpsToInterruptVector(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x34
	bus.RAM[0xFFFD] = 0x12
	bus.RAM[0xFFFE] = 0x00
	bus.RAM[0xFFFF] = 0x80
	bus.RAM[0x1234] = 0x00 // BRK

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagI, false)
	cpu.SetFlag(CPU6502.FlagC, true)
	cpu.SetFlag(CPU6502.FlagN, true)
	statusBefore := cpu.Status
	spBefore := cpu.SP

	// when:
	for range 15 {
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x8000 {
		t.Errorf("PC = 0x%04X, want 0x8000", cpu.PC)
	}

	if cpu.SP != spBefore-3 {
		t.Errorf("SP = 0x%02X, want 0x%02X", cpu.SP, spBefore-3)
	}

	hi := bus.RAM[0x0100+uint16(spBefore)]
	lo := bus.RAM[0x0100+uint16(spBefore)-1]
	if hi != 0x12 || lo != 0x36 {
		t.Errorf("pushed PC = 0x%02X%02X, want 0x1236", hi, lo)
	}

	pushed := bus.RAM[0x0100+uint16(spBefore)-2]
	wantPushed := statusBefore | CPU6502.FlagB | CPU6502.FlagU
	if pushed != wantPushed {
		t.Errorf("pushed status = 0x%02X, want 0x%02X", pushed, wantPushed)
	}

	if !cpu.GetFlag(CPU6502.FlagI) {
		t.Error("I flag should be set")
	}

	if (cpu.Status &^ CPU6502.FlagI) != (statusBefore &^ CPU6502.FlagI) {
		t.Errorf("other flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
	}
}

// NOP
func TestNOPDoesNothing(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0xEA // NOP

	cpu.Reset()
	cpu.A = 0x11
	cpu.X = 0x22
	cpu.Y = 0x33
	statusBefore := cpu.Status
	spBefore := cpu.SP

	// when:
	for range 10 { // 8 reset + 2 NOP
		cpu.Clock()
	}

	// then:
	if cpu.PC != 0x0001 {
		t.Errorf("PC = 0x%04X, want 0x0001", cpu.PC)
	}

	if cpu.A != 0x11 || cpu.X != 0x22 || cpu.Y != 0x33 {
		t.Errorf("A,X,Y = 0x%02X,0x%02X,0x%02X, want 0x11,0x22,0x33", cpu.A, cpu.X, cpu.Y)
	}

	if cpu.SP != spBefore {
		t.Errorf("SP = 0x%02X, want 0x%02X", cpu.SP, spBefore)
	}

	expectUnaffectedFlagsAll(t, statusBefore, cpu.Status)
}

// RTI
func TestRTIPullsStatusAndProgramCounterFromTheStack(t *testing.T) {
	tests := []struct {
		name       string
		onStack    byte
		wantStatus byte
	}{
		{"PullsFlags", 0xC3, 0xE3},
		{"IgnoresBFromStack", 0x30, 0x20},
		{"AllBitsSet", 0xFF, 0xEF},
		{"NoFlags", 0x00, 0x20},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x40 // RTI
			bus.RAM[0x01FD] = tc.onStack
			bus.RAM[0x01FE] = 0x34
			bus.RAM[0x01FF] = 0x12

			cpu.Reset()
			cpu.SP = 0xFC
			cpu.A = 0xAA

			// when:
			for range 14 { // 8 reset + 6 RTI
				cpu.Clock()
			}

			// then:
			if cpu.PC != 0x1234 {
				t.Errorf("PC = 0x%04X, want 0x1234", cpu.PC)
			}

			if cpu.SP != 0xFF {
				t.Errorf("SP = 0x%02X, want 0xFF", cpu.SP)
			}

			if cpu.Status != tc.wantStatus {
				t.Errorf("Status = 0x%02X, want 0x%02X", cpu.Status, tc.wantStatus)
			}

			if cpu.A != 0xAA {
				t.Errorf("A = 0x%02X, want unchanged 0xAA", cpu.A)
			}
		})
	}
}

// BRK + RTI
func TestCanBreakAndReturnFromInterruptAgain(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0xFFFE] = 0x00
	bus.RAM[0xFFFF] = 0x80
	bus.RAM[0x0000] = 0x00 // BRK
	bus.RAM[0x0001] = 0xE8
	bus.RAM[0x0002] = 0xA9
	bus.RAM[0x0003] = 0x11
	bus.RAM[0x8000] = 0x40 // RTI

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagI, false)
	cpu.SetFlag(CPU6502.FlagC, true)
	statusBefore := cpu.Status
	spBefore := cpu.SP

	// when:
	for range 23 {
		cpu.Clock()
	}

	// then:
	if cpu.A != 0x11 {
		t.Errorf("A = 0x%02X, want 0x11", cpu.A)
	}

	if cpu.X != 0x00 {
		t.Errorf("X = 0x%02X, want 0x00 (padding byte was executed)", cpu.X)
	}

	if cpu.SP != spBefore {
		t.Errorf("SP = 0x%02X, want 0x%02X", cpu.SP, spBefore)
	}

	const mask = CPU6502.FlagC | CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagV
	if cpu.Status&mask != statusBefore&mask {
		t.Errorf("flags not restored: before=%08b after=%08b", statusBefore, cpu.Status)
	}
}

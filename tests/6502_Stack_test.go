package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// TSX
func TestTSXCanTransferTheStackPointerToXRegister(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0xBA // TSX

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagZ, true)
	cpu.SetFlag(CPU6502.FlagN, true)
	statusBefore := cpu.Status
	cpu.X = 0x00
	cpu.SP = 0x01

	// when:
	for i := 0; i < 10; i++ {
		cpu.Clock()
	}

	// then:
	if cpu.X != 0x01 {
		t.Errorf("cpu.X = %#02X, want 0x01", cpu.X)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestTSXCanTransferAZeroStackPointerToXRegister(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0xBA // TSX

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagZ, true)
	cpu.SetFlag(CPU6502.FlagN, true)
	statusBefore := cpu.Status
	cpu.X = 0x01
	cpu.SP = 0x00

	// when:
	for i := 0; i < 10; i++ {
		cpu.Clock()
	}

	// then:
	if cpu.X != 0x00 {
		t.Errorf("cpu.X = %#02X, want 0x00;", cpu.X)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestTSXCanTransferANegativeStackPointerToXRegister(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0xBA // TSX

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagZ, true)
	cpu.SetFlag(CPU6502.FlagN, true)
	statusBefore := cpu.Status
	cpu.X = 0x01
	cpu.SP = 0b10000000

	// when:
	for i := 0; i < 10; i++ {
		cpu.Clock()
	}

	// then:
	if cpu.X != 0b10000000 {
		t.Errorf("cpu.X = %#02X, want 10000000(bin);", cpu.X)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// TXS

func TestTXSCanTransferXRegisterToTheStackPointer(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x9A // TXS

	cpu.Reset()
	cpu.SetFlag(CPU6502.FlagZ, true)
	cpu.SetFlag(CPU6502.FlagN, true)
	statusBefore := cpu.Status
	cpu.X = 0xFF
	cpu.SP = 0

	// when:
	for i := 0; i < 10; i++ {
		cpu.Clock()
	}

	// then:
	if cpu.SP != 0xFF {
		t.Errorf("cpu.SP = %#02X, want 0xFF", cpu.X)
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// PHA
func TestPHACanPushARegsiterOntoTheStack(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x48 // PHA

	cpu.Reset()
	statusBefore := cpu.Status
	cpu.A = 0xFF

	// when:
	for i := 0; i < 11; i++ {
		cpu.Clock()
	}

	// then:
	if cpu.Read(Bus.Word(cpu.SP)+1) != cpu.A {
		t.Errorf("mem[c.SP] = %#02X, want 0xFF", Bus.Word(cpu.SP)+1)
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// PLA
func TestPLACanPullAValueFromTheStackIntoTheARegsiter(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x68 // PLA

	cpu.Reset()
	statusBefore := cpu.Status
	cpu.A = 0x00
	cpu.SP = 0xFE
	bus.RAM[0x1FF] = 0x42

	// when:
	for i := 0; i < 12; i++ {
		cpu.Clock()
	}

	// then:
	if cpu.A != 0x00 {
		t.Errorf("cpu.A = %#02X, want 0x00", cpu.A)
	}

	if cpu.SP != 0xFF {
		t.Errorf("SP = %#02X, want 0xFF", cpu.SP)
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestPLACanPullAZeroValueFromTheStackIntoTheARegsiter(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x68 // PLA

	cpu.Reset()
	statusBefore := cpu.Status
	cpu.A = 0x00
	cpu.SP = 0xFE
	bus.RAM[0x1FF] = 0x00

	// when:
	for i := 0; i < 12; i++ {
		cpu.Clock()
	}

	// then:
	if cpu.A != 0x00 {
		t.Errorf("cpu.A = %#02X, want 0x00", cpu.A)
	}

	if cpu.SP != 0xFF {
		t.Errorf("SP = %#02X, want 0xFF", cpu.SP)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// PHP

func TestPHPCanPushProcessorStatusOntoTheStack(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x08 // PHP

	cpu.Reset()
	cpu.Status = 0xCC
	spBefore := cpu.SP

	// when:
	for i := 0; i < 11; i++ { // 8 reset + 3 PHP
		cpu.Clock()
	}

	// then:
	pushed := bus.RAM[0x0100+Bus.Word(spBefore)]
	if pushed != 0xFC { // B and U are set in the pushed copy
		t.Errorf("pushed = 0x%02X, want 0xFC", pushed)
	}

	if cpu.SP != spBefore-1 {
		t.Errorf("SP = 0x%02X, want 0x%02X", cpu.SP, spBefore-1)
	}

	if cpu.Status != 0xCC {
		t.Errorf("Status = 0x%02X, want unchanged 0xCC", cpu.Status)
	}
}

// TODO: complete tests

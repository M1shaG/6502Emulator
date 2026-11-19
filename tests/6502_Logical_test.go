package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// AND Immediate
func TestANDImmediateCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
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
	for range 10 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDImmediateCanLogicalAndZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x29 // AND Immediate
	bus.RAM[0x0001] = 0x00

	cpu.Reset()
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for range 10 {
		cpu.Clock()
	}

	var AndValue byte = 0x00 & 0xFF
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDImmediateCanLogicalAndNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x29 // AND Immediate
	bus.RAM[0x0001] = 0x80

	cpu.Reset()
	cpu.A = 0x83
	statusBefore := cpu.Status

	// when:
	for range 10 {
		cpu.Clock()
	}

	var AndValue byte = 0x80 & 0x83
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// AND Zero Page
func TestANDZeroPageCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x25 // AND Zero Page
	bus.RAM[0x0001] = 0x0F
	bus.RAM[0x000F] = 0x2C

	cpu.Reset()
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDZeroPageCanLogicalAndZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x25 // AND Zero Page
	bus.RAM[0x0001] = 0x0F
	bus.RAM[0x000F] = 0x00

	cpu.Reset()
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	var AndValue byte = 0x00 & 0xFF
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDZeroPageCanLogicalAndNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x25 // AND Zero Page
	bus.RAM[0x0001] = 0x0F
	bus.RAM[0x000F] = 0x80

	cpu.Reset()
	cpu.A = 0x83
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	var AndValue byte = 0x80 & 0x83
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// AND Zero Page X
func TestANDZeroPageXCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x35 // AND Zero Page X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x00FC] = 0x2C

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDZeroPageXCanLogicalAndZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x35 // AND Zero Page X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x00FC] = 0x00

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	var AndValue byte = 0x00 & 0xFF
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDZeroPageXCanLogicalAndNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x35 // AND Zero Page X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x00FC] = 0x80

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0x83
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	var AndValue byte = 0x80 & 0x83
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// And Absolute

func TestANDAbsoluteCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x2D // AND Absolute
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52F0] = 0x2C

	cpu.Reset()
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDAbsoluteCanLogicalAndZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x2D // AND Absolute
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52F0] = 0x00

	cpu.Reset()
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	var AndValue byte = 0x00 & 0xFF
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDAbsoluteCanLogicalAndNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x2D // AND Absolute
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52F0] = 0x80

	cpu.Reset()
	cpu.A = 0x83
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	var AndValue byte = 0x80 & 0x83
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// And Absolute X
func TestANDAbsoluteXCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x3D // AND Absolute X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x2C

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDAbsoluteXCanLogicalAndZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x3D // AND Absolute X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x00

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	var AndValue byte = 0x00 & 0xFF
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDAbsoluteXCanLogicalAndNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x3D // AND Absolute X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x80

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0x83
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	var AndValue byte = 0x80 & 0x83
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDAbsoluteXPageBoundaryCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x3D // AND Absolute X
	bus.RAM[0x0001] = 0xFF
	bus.RAM[0x0002] = 0x02
	bus.RAM[0x0300] = 0x2C

	cpu.Reset()
	cpu.X = 0x01
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// And Absolute Y
func TestANDAbsoluteYCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x39 // AND Absolute Y
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x2C

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDAbsoluteYCanLogicalAndZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x39 // AND Absolute Y
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x00

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	var AndValue byte = 0x00 & 0xFF
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDAbsoluteYCanLogicalAndNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x39 // AND Absolute Y
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x80

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0x83
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	var AndValue byte = 0x80 & 0x83
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDAbsoluteYPageBoundaryCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x39 // AND Absolute Y
	bus.RAM[0x0001] = 0xFF
	bus.RAM[0x0002] = 0x02
	bus.RAM[0x0300] = 0x2C

	cpu.Reset()
	cpu.Y = 0x01
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// And Indirect X
func TestANDIndirectXCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x21 // AND Indirect X
	bus.RAM[0x0001] = 0x02
	bus.RAM[0x0006] = 0xF0
	bus.RAM[0x0007] = 0x52
	bus.RAM[0x52F0] = 0x2C

	cpu.Reset()
	cpu.X = 0x04
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDIndirectXCanLogicalAndZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x21 // AND Indirect X
	bus.RAM[0x0001] = 0x02
	bus.RAM[0x0006] = 0xF0
	bus.RAM[0x0007] = 0x52
	bus.RAM[0x52F0] = 0x00

	cpu.Reset()
	cpu.X = 0x04
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	var AndValue byte = 0x00 & 0xFF
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDIndirectXCanLogicalAndNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x21 // AND Indirect X
	bus.RAM[0x0001] = 0x02
	bus.RAM[0x0006] = 0xF0
	bus.RAM[0x0007] = 0x52
	bus.RAM[0x52F0] = 0x80

	cpu.Reset()
	cpu.X = 0x04
	cpu.A = 0x83
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	var AndValue byte = 0x80 & 0x83
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// And Indirect Y
func TestANDIndirectYCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x31 // AND Indirect Y
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xF0
	bus.RAM[0x0011] = 0x52
	bus.RAM[0x52FC] = 0x2C

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDIndirectYCanLogicalAndZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x31 // AND Indirect Y
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xF0
	bus.RAM[0x0011] = 0x52
	bus.RAM[0x52FC] = 0x00

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	var AndValue byte = 0x00 & 0xFF
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDIndirectYCanLogicalAndNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x31 // AND Indirect Y
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xF0
	bus.RAM[0x0011] = 0x52
	bus.RAM[0x52FC] = 0x80

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0x83
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	var AndValue byte = 0x80 & 0x83
	// then:
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestANDIndirectYPageBoundaryCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x31 // AND Indirect Y
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xFF
	bus.RAM[0x0011] = 0x02
	bus.RAM[0x0300] = 0x2C

	cpu.Reset()
	cpu.Y = 0x01
	cpu.A = 0x28
	statusBefore := cpu.Status

	// when:
	for _ = range 14 {
		cpu.Clock()
	}

	// then:
	var AndValue byte = 0x28 & 0x2C
	if cpu.A != AndValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, AndValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// EOR Immediate
func TestEORImmediateCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x49 // EOR Immediate
	bus.RAM[0x0001] = 0x4F

	cpu.Reset()
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORImmediateCanLogicalEORZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x49 // EOR Immediate
	bus.RAM[0x0001] = 0xFF

	cpu.Reset()
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0xFF ^ 0xFF
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORImmediateCanLogicalEORNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x49 // EOR Immediate
	bus.RAM[0x0001] = 0x9F

	cpu.Reset()
	cpu.A = 0x79
	statusBefore := cpu.Status

	// when:
	for _ = range 10 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x9F ^ 0x79
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// EOR Zero Page
func TestEORZeroPageCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x45 // EOR Zero Page
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0x4F

	cpu.Reset()
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for _ = range 11 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORZeroPageCanLogicalEORZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x45 // EOR Zero Page
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xFF

	cpu.Reset()
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 11 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0xFF ^ 0xFF
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORZeroPageCanLogicalEORNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x45 // EOR Zero Page
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0x9F

	cpu.Reset()
	cpu.A = 0x79
	statusBefore := cpu.Status

	// when:
	for _ = range 11 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x9F ^ 0x79
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// EOR Zero Page X
func TestEORZeroPageXCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x55 // EOR Zero Page X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0018] = 0x4F

	cpu.Reset()
	cpu.X = 0x08
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORZeroPageXCanLogicalEORZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x55 // EOR Zero Page X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0018] = 0xFF

	cpu.Reset()
	cpu.X = 0x08
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0xFF ^ 0xFF
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORZeroPageXCanLogicalEORNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x55 // EOR Zero Page X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0018] = 0x9F

	cpu.Reset()
	cpu.X = 0x08
	cpu.A = 0x79
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x9F ^ 0x79
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// EOR Absolute
func TestEORAbsoluteCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x4D // EOR Absolute
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52F0] = 0x4F

	cpu.Reset()
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORAbsoluteCanLogicalEORZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x4D // EOR Absolute
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52F0] = 0xFF

	cpu.Reset()
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0xFF ^ 0xFF
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORAbsoluteCanLogicalEORNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x4D // EOR Zero Page X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52F0] = 0x9F

	cpu.Reset()
	cpu.A = 0x79
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x9F ^ 0x79
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// EOR Absolute X
func TestEORAbsoluteXCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x5D // EOR Absolute X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x4F

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORAbsoluteXCanLogicalEORZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x5D // EOR Absolute X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0xFF

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0xFF ^ 0xFF
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORAbsoluteXCanLogicalEORNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x5D // EOR Absolute X
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x9F

	cpu.Reset()
	cpu.X = 0x0C
	cpu.A = 0x79
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x9F ^ 0x79
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORAbsoluteXPageBoundaryCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x5D // EOR Absolute X
	bus.RAM[0x0001] = 0xFF
	bus.RAM[0x0002] = 0x02
	bus.RAM[0x0300] = 0x4F

	cpu.Reset()
	cpu.X = 0x01
	cpu.A = 0x0C
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x0C
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// EOR Absolute Y
func TestEORAbsoluteYCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x59 // EOR Absolute Y
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x4F

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORAbsoluteYCanLogicalEORZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x59 // EOR Absolute Y
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0xFF

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0xFF ^ 0xFF
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORAbsoluteYCanLogicalEORNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x59 // EOR Absolute Y
	bus.RAM[0x0001] = 0xF0
	bus.RAM[0x0002] = 0x52
	bus.RAM[0x52FC] = 0x9F

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0x79
	statusBefore := cpu.Status

	// when:
	for _ = range 12 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x9F ^ 0x79
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORAbsoluteYPageBoundaryCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x59 // EOR Absolute Y
	bus.RAM[0x0001] = 0xFF
	bus.RAM[0x0002] = 0x02
	bus.RAM[0x0300] = 0x4F

	cpu.Reset()
	cpu.Y = 0x01
	cpu.A = 0x0C
	statusBefore := cpu.Status

	// when:
	for _ = range 13 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x0C
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// EOR Indirect X
func TestEORIndirectXCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x41 // EOR Indirect X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0014] = 0xF0
	bus.RAM[0x0015] = 0x52
	bus.RAM[0x52F0] = 0x4F

	cpu.Reset()
	cpu.X = 0x04
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for _ = range 14 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORIndirectXCanLogicalEORZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x41 // EOR Indirect X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0014] = 0xF0
	bus.RAM[0x0015] = 0x52
	bus.RAM[0x52F0] = 0xFF

	cpu.Reset()
	cpu.X = 0x04
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for _ = range 14 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0xFF ^ 0xFF
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORIndirectXCanLogicalEORNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x41 // EOR Indirect X
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0014] = 0xF0
	bus.RAM[0x0015] = 0x52
	bus.RAM[0x52F0] = 0x9F

	cpu.Reset()
	cpu.X = 0x04
	cpu.A = 0x79
	statusBefore := cpu.Status

	// when:
	for range 14 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x9F ^ 0x79
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// EOR Indirect Y
func TestEORIndirectYCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x51 // EOR Indirect Y
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xF0
	bus.RAM[0x0011] = 0x52
	bus.RAM[0x52FC] = 0x4F

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for range 13 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORIndirectYCanLogicalEORZeroValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x51 // EOR Indirect Y
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xF0
	bus.RAM[0x0011] = 0x52
	bus.RAM[0x52FC] = 0xFF

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0xFF
	statusBefore := cpu.Status

	// when:
	for range 13 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0xFF ^ 0xFF
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ == 0 {
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORIndirectYCanLogicalEORNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x51 // EOR Indirect Y
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xF0
	bus.RAM[0x0011] = 0x52
	bus.RAM[0x52FC] = 0x9F

	cpu.Reset()
	cpu.Y = 0x0C
	cpu.A = 0x79
	statusBefore := cpu.Status

	// when:
	for range 13 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x9F ^ 0x79
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be set")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

func TestEORIndirectYPageBoundaryCanLogicalEORNonZeroNonNegativeValue(t *testing.T) {
	// given:
	bus := Bus.NewBus()
	cpu := &CPU6502.CPU{}
	cpu.ConnectBus(bus)

	bus.RAM[0xFFFC] = 0x00
	bus.RAM[0xFFFD] = 0x00
	bus.RAM[0x0000] = 0x51 // EOR Indirect Y
	bus.RAM[0x0001] = 0x10
	bus.RAM[0x0010] = 0xFF
	bus.RAM[0x0011] = 0x02
	bus.RAM[0x0300] = 0x4F

	cpu.Reset()
	cpu.Y = 0x01
	cpu.A = 0x78
	statusBefore := cpu.Status

	// when:
	for range 14 {
		cpu.Clock()
	}

	// then:
	var EORValue byte = 0x4F ^ 0x78
	if cpu.A != EORValue {
		t.Errorf("cpu.A = %#02X, want %#02X", cpu.A, EORValue)
	}

	if cpu.Status&CPU6502.FlagZ != 0 {
		t.Error("Z flag should be clear")
	}

	if cpu.Status&CPU6502.FlagN != 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// ORA Immediate
func TestORAImmediateCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
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

func TestORAImmediateCanLogicalAndZeroValue(t *testing.T) {
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

func TestORAImmediateCanLogicalAndNegativeValue(t *testing.T) {
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
		t.Error("Z flag should be set")
	}

	if cpu.Status&CPU6502.FlagN == 0 {
		t.Error("N flag should be clear")
	}

	expectUnaffectedFlags(t, statusBefore, cpu.Status)
}

// ORA Zero Page
func TestORAZeroPageCanLogicalAndNonZeroNonNegativeValue(t *testing.T) {
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

func TestORAZeroPageCanLogicalAndZeroValue(t *testing.T) {
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

func TestORAZeroPageCanLogicalAndNegativeValue(t *testing.T) {
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
	for range 10 {
		cpu.Clock()
	}

	// then:
	var ORValue byte = 0xF0 | 0x0F
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

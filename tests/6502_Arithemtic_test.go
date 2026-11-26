package tests

import (
	CPU6502 "nesem/6502"
	Bus "nesem/bus"
	"testing"
)

// ADC Immediate
func TestADCImmediateCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x69 // ADC Immediate
			bus.RAM[0x0001] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 10 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// ADC Zero Page
func TestADCZeroPageCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x65 // ADC Zero Page
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 11 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// ADC Zero Page X
func TestADCZeroPageXCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x75 // ADC Zero Page X
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0014] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.X = 0x04
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 12 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// ADC Absolute
func TestADCAbsoluteCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x6D // ADC Absolute
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 12 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// ADC Absolute X
func TestADCAbsoluteXCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x7D // ADC Absolute X
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x52FC] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.X = 0x0C
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 12 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

func TestADCAbsoluteXPageCrossedCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x7D // ADC Absolute X
			bus.RAM[0x0001] = 0xFF
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x5300] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.X = 0x01
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// ADC Absolute Y
func TestADCAbsoluteYCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x79 // ADC Absolute Y
			bus.RAM[0x0001] = 0xF0
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x52FC] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.Y = 0x0C
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 12 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

func TestADCAbsoluteYPageCrossedCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x79 // ADC Absolute Y
			bus.RAM[0x0001] = 0xFF
			bus.RAM[0x0002] = 0x52
			bus.RAM[0x5300] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.Y = 0x01
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// ADC Indirect X
func TestADCIndirectXCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x61 // ADC Indirect X
			bus.RAM[0x0001] = 0x02
			bus.RAM[0x0006] = 0xF0
			bus.RAM[0x0007] = 0x52
			bus.RAM[0x52F0] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.X = 0x04
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// ADC Indirect Y
func TestADCIndirectYCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x71 // ADC Indirect Y
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = 0xF0
			bus.RAM[0x0011] = 0x52
			bus.RAM[0x52F4] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.Y = 0x04
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 13 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

// ADC Indirect Y
func TestADCIndirectYPageCrossedCanADD(t *testing.T) {
	tests := []struct {
		name                       string
		carryIn                    bool
		valMem, valReg             byte
		wantA                      byte
		wantC, wantZ, wantN, wantV bool
	}{
		{"NonZeroNonNegativeValue", true, 0x20, 0x20, 0x41, false, false, false, false},
		{"NoCarryIn", false, 0x20, 0x10, 0x30, false, false, false, false},
		{"ZeroValue", false, 0x00, 0x00, 0x00, false, true, false, false},
		{"ZeroFromCarryOut", true, 0x00, 0xFF, 0x00, true, true, false, false},
		{"NegativeResultNoOverflow", false, 0x80, 0x01, 0x81, false, false, true, false},
		{"CarryOutAndOverflow", true, 0x83, 0xA9, 0x2D, true, false, false, true},
		{"PosPlusPosOverflow", false, 0x50, 0x50, 0xA0, false, false, true, true},
		{"CarryInCausesOverflow", true, 0x00, 0x7F, 0x80, false, false, true, true},
		{"NegPlusNegZeroOverflow", false, 0x80, 0x80, 0x00, true, true, false, true},
		{"NegPlusNegNoOverflow", true, 0xFF, 0xFF, 0xFF, true, false, true, false},
		{"PosPlusNegCarryNoOverflow", false, 0xD0, 0x50, 0x20, true, false, false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given:
			bus := Bus.NewBus()
			cpu := &CPU6502.CPU{}
			cpu.ConnectBus(bus)

			bus.RAM[0xFFFC] = 0x00
			bus.RAM[0xFFFD] = 0x00
			bus.RAM[0x0000] = 0x71 // ADC Indirect Y
			bus.RAM[0x0001] = 0x10
			bus.RAM[0x0010] = 0xFF
			bus.RAM[0x0011] = 0x52
			bus.RAM[0x5300] = tc.valMem

			cpu.Reset()
			cpu.A = tc.valReg
			cpu.Y = 0x01
			cpu.SetFlag(CPU6502.FlagC, tc.carryIn)
			cpu.SetFlag(CPU6502.FlagZ, !tc.wantZ)
			cpu.SetFlag(CPU6502.FlagN, !tc.wantN)
			cpu.SetFlag(CPU6502.FlagV, !tc.wantV)
			statusBefore := cpu.Status

			// when:
			for range 14 {
				cpu.Clock()
			}

			// then:
			if cpu.A != tc.wantA {
				t.Errorf("A = 0x%02X, want 0x%02X", cpu.A, tc.wantA)
			}
			if got := cpu.Status&CPU6502.FlagC != 0; got != tc.wantC {
				t.Errorf("C flag = %v, want %v", got, tc.wantC)
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

			const unaffectedMask = CPU6502.FlagI | CPU6502.FlagD | CPU6502.FlagB
			if statusBefore&unaffectedMask != cpu.Status&unaffectedMask {
				t.Errorf("unaffected flags changed: before=%08b after=%08b", statusBefore, cpu.Status)
			}
		})
	}
}

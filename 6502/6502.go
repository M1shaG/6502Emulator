// Useful links:
// 6502 Documentation:
// https://6502.org/users/obelisk/6502/instructions.html
package CPU6502

import (
	Bus "nesem/bus"
)

// CPU STRUCT
type CPU struct {
	// Information of registers:
	// https://6502.org/users/obelisk/6502/registers.html
	A, X, Y byte // Accumulator, X, Y
	SP      byte
	PC      Bus.Word

	// Processor Status
	// #7 #6 #5 #4 #3 #2 #1 #0 - List of bits
	// #0 (C) - Carry Flag
	// #1 (Z) - Zero Flag
	// #2 (I) - Interrupt Disable Flag
	// #3 (D) - Decimal Mode Flag
	// #4 (B) - Break Flag
	// #5 - unused
	// #6 (V) - Overflow Flag
	// #7 (N) - Negative Flag
	// https://www.middle-engine.com/images/2020-06-23-programming-the-nes-the-6502-in-detail/processor-status-register-2x.png
	Status byte

	// It is not part of the processor architecture, but it is necessary for the program to run
	bus *Bus.Bus

	addrAbs Bus.Word
	addrRel Bus.Word
	fetched byte
	temp    Bus.Word

	cycles       byte
	opcode       byte
	currentInstr Instruction
}

// Represent each status as one active bit for bit manipulation

const (
	FlagC byte = 1 << iota // 00000001
	FlagZ                  // 00000010
	FlagI                  // 00000100
	FlagD                  // 00001000
	FlagB                  // 00010000
	FlagU                  // 00100000
	FlagV                  // 01000000
	FlagN                  // 10000000
)

// Bit manipulation functions
// Logical and with CPU.Status AND FlagX gives answer if flag active
func (c *CPU) GetFlag(f byte) bool {
	return c.Status&f > 0
}

func (c *CPU) SetFlag(f byte, v bool) {
	if v {
		c.Status |= f
	} else {
		c.Status &^= f
	}
}

// Connect to Bus that represents 16KB RAM(simple array of 65536 elements)
func (c *CPU) ConnectBus(b *Bus.Bus) {
	c.bus = b
}

// Read from Bus (16bit address)
func (c *CPU) Read(addr Bus.Word) byte {
	return c.bus.Read(addr)
}

// Write to Bus (8bit value, 16bit address)
func (c *CPU) Write(addr Bus.Word, data byte) {
	c.bus.Write(addr, data)
}

// Reads the data required for the instruction
// Implied not read data for execution, so it ignored
func (c *CPU) fetch() byte {
	if !c.currentInstr.isImplied {
		c.fetched = c.Read(c.addrAbs)
	}
	return c.fetched

}

// Reset CPU (Initialize and start program execution)
// Registers are set to 0x00, the status register is cleared
// Absolute address is read from location 0xFFFC
// In this address(that and next one to read 16bit) contains address from where reed instruction
func (c *CPU) Reset() {
	// Get address to set program counter to
	c.addrAbs = 0xFFFC
	lo := c.Read(c.addrAbs + 0)
	hi := c.Read(c.addrAbs + 1)

	// Set Program Counter
	c.PC = (Bus.Word(hi) << 8) | Bus.Word(lo)

	// Reset registers
	c.A = 0
	c.X = 0
	c.Y = 0
	c.SP = 0
	c.Status = 0x00

	// Clear internal helper variables
	c.addrRel = 0x0000
	c.addrAbs = 0x0000
	c.fetched = 0x00

	// Reset takes time
	c.cycles = 8
}

// Perform one clock cycles worth of emulation
func (c *CPU) Clock() {
	// Each instruction requires a certain number of clock cycles
	// To remain compilant with connected devices
	// Implement cycles as simply counting down the cycles required for instruction

	// Time to get new instruction
	if c.cycles == 0 {
		// Read current instruction
		c.opcode = c.Read(c.PC)

		// Increase PC
		c.PC++

		// Handful current instruction
		c.currentInstr = opcodeTable[c.opcode]

		// Set cycles
		c.cycles = opcodeTable[c.opcode].Cycles

		// Perform operation
		// Get addressing mode
		additional_cycle1 := c.currentInstr.AddrMode(c)

		// Perform instruction
		additional_cycle2 := c.currentInstr.Operate(c)

		// Due to CPU architecture, some combination of instruction and address mode
		// Can require additional cycle so we check if instruction AND address mode
		// Requires additional cycle
		c.cycles += (additional_cycle1 & additional_cycle2)
	}

	// One clock cycle, one cycles goes
	c.cycles--
}

// Addresses

// Additional documentation:
// https://6502.org/users/obelisk/6502/addressing.html

// Address Mode: Implicit
// No need read value
func (c *CPU) IMP() byte {
	c.fetched = c.A
	return 0
}

// Address Mode; Immediate
// The value for instruction lays on next byte
func (c *CPU) IMM() byte {
	c.addrAbs = c.PC
	c.PC++
	return 0
}

// Address Mode: Zero Page
// Address to first 256 bytes $0000 - $00FF
func (c *CPU) ZP0() byte {
	c.addrAbs = Bus.Word(c.Read(c.PC))
	c.PC++
	return 0
}

// Address Mode: Zero Page X
// Same as Zero Page but we add X register
// if add not feet into 2 bytes we drop first byte #1FF => Byte with value 1 drops
func (c *CPU) ZPX() byte {
	c.addrAbs = Bus.Word(c.Read(c.PC) + c.X)
	c.PC++
	c.addrAbs &= 0x00FF
	return 0
}

// Address Mode: Zero Page Y
// Same as Zero Page X, but we add Y register
func (c *CPU) ZPY() byte {
	c.addrAbs = Bus.Word(c.Read(c.PC) + c.Y)
	c.PC++
	return 0
}

// Address Mode: Relative
// Exclusive for branch operations and used for
// add bytes to program counter if condition true
func (c *CPU) REL() byte {
	c.addrRel = Bus.Word(c.Read(c.PC))
	c.PC++
	if c.addrRel&0x80 != 0 {
		c.addrRel |= 0xFF00
	}
	return 0
}

// Address Mode: Absolute
// Addresses full 16bit value
func (c *CPU) ABS() byte {
	lo := c.Read(c.PC)
	c.PC++
	hi := c.Read(c.PC)
	c.PC++

	c.addrAbs = Bus.Word(hi)<<8 | Bus.Word(lo)
	return 0
}

// Address Mode: Absolute X
// Same as Absolute but we add X register
// If the resulting address changes the page, an additional clock cycle is required
func (c *CPU) ABX() byte {
	lo := c.Read(c.PC)
	c.PC++
	hi := c.Read(c.PC)
	c.PC++

	c.addrAbs = Bus.Word(hi)<<8 | Bus.Word(lo)
	c.addrAbs += Bus.Word(c.X)

	if (c.addrAbs & 0xFF00) != (Bus.Word(hi) << 8) {
		return 1
	} else {
		return 0
	}
}

// Address Mode: Absolute Y
// Same as Absolute X but we add Y register
func (c *CPU) ABY() byte {
	lo := c.Read(c.PC)
	c.PC++
	hi := c.Read(c.PC)
	c.PC++

	c.addrAbs = Bus.Word(hi)<<8 | Bus.Word(lo)
	c.addrAbs += Bus.Word(c.Y)

	if (c.addrAbs & 0xFF00) != (Bus.Word(hi) << 8) {
		return 1
	} else {
		return 0
	}
}

// Address Mode: Indirect
// Read next 2 bytes to get an address. There is hardware bug
// If low byte addr 0xFF, read high bite we need to create a page boundary.
// But this doesnt actually work on the chip as designed and yielding an invalid address.
func (c *CPU) IND() byte {
	ptr_lo := c.Read(c.PC)
	c.PC++
	ptr_hi := c.Read(c.PC)
	c.PC++

	ptr := Bus.Word(ptr_hi)<<8 | Bus.Word(ptr_lo)
	if ptr_lo == 0x00FF {
		c.addrAbs = Bus.Word(c.Read(ptr&0xFF00))<<8 | Bus.Word(ptr+0)
	} else {
		c.addrAbs = Bus.Word(c.Read(ptr+1&0xFF00))<<8 | Bus.Word(ptr+0)
	}
	return 0
}

// Address Mode: Indirect X
// Same as Indirect but we add X to each byte
func (c *CPU) IZX() byte {
	t := c.Read(c.PC)
	c.PC++

	lo := c.Read(Bus.Word(t+c.X) & 0x00FF)
	hi := c.Read(Bus.Word(t+c.X+1) & 0x00FF)

	c.addrAbs = Bus.Word(hi)<<8 | Bus.Word(lo)

	return 0
}

// Address Mode: Indexed Y
// Reads first address from next two bytes then
// adds Y registers to that address
func (c *CPU) IZY() byte {
	t := c.Read(c.PC)
	c.PC++

	lo := c.Read(Bus.Word(t) & 0x00FF)
	hi := c.Read(Bus.Word(t+1) & 0x00FF)

	c.addrAbs = Bus.Word(hi)<<8 | Bus.Word(lo)
	c.addrAbs += Bus.Word(c.Y)

	if (c.addrAbs & 0xFF00) != (Bus.Word(hi) << 8) {
		return 1
	} else {
		return 0
	}
}

// Load/Store Operations

// Instruction: Load The Accumulator
// Function: A = M
// Flags Out: N, Z
func (c *CPU) LDA() byte {
	c.fetch()
	c.A = c.fetched
	c.SetFlag(FlagZ, c.A == 0x00)
	c.SetFlag(FlagN, c.A&0x80 != 0)
	return 1
}

// Instruction: Load The X Registers
// Function: X = M
// Flags Out: N, Z
func (c *CPU) LDX() byte {
	c.fetch()
	c.X = c.fetched
	c.SetFlag(FlagZ, c.X == 0x00)
	c.SetFlag(FlagN, c.X&0x80 != 0)
	return 1
}

// Instruction: Load The Y Registers
// Function: Y = M
// Flags Out: N, Z
func (c *CPU) LDY() byte {
	c.fetch()
	c.Y = c.fetched
	c.SetFlag(FlagZ, c.Y == 0x00)
	c.SetFlag(FlagN, c.Y&0x80 != 0)
	return 1
}

// Store

// Instruction: Store Accumulator at Address
// Function: M = A
func (c *CPU) STA() byte {
	c.Write(c.addrAbs, c.A)
	return 0
}

// Instruction: Store X Register at Address
// Function: M = X
func (c *CPU) STX() byte {
	c.Write(c.addrAbs, c.X)
	return 0
}

// Instruction: Store Y Register at Address
// Function: M = Y
func (c *CPU) STY() byte {
	c.Write(c.addrAbs, c.Y)
	return 0
}

// Register Transfers

// Instruction: Transfer Accumulator to X Register
// Function: X = A
// Flags Out: N, Z
func (c *CPU) TAX() byte {
	c.X = c.A
	c.SetFlag(FlagZ, c.X == 0x00)
	c.SetFlag(FlagN, c.X&0x80 != 0)
	return 0
}

// Instruction: Transfer Accumulator to Y Register
// Function: Y = A
// Flags Out: N, Z
func (c *CPU) TAY() byte {
	c.Y = c.A
	c.SetFlag(FlagZ, c.Y == 0x00)
	c.SetFlag(FlagN, c.Y&0x80 != 0)
	return 0
}

// Instruction: Transfer X Register to Accumulator
// Function: A = X
// Flags Out: N, Z
func (c *CPU) TXA() byte {
	c.A = c.X
	c.SetFlag(FlagZ, c.A == 0x00)
	c.SetFlag(FlagN, c.A&0x80 != 0)
	return 0
}

// Instruction: Transfer Y Register to Accumulator
// Function: A = Y
// Flags Out: N, Z
func (c *CPU) TYA() byte {
	c.A = c.Y
	c.SetFlag(FlagZ, c.A == 0x00)
	c.SetFlag(FlagN, c.A&0x80 != 0)
	return 0
}

// Stack Operations

// Instruction: Transfer Stack Pointer to X Register
// Function: X = stack pointer
// Flags Out: N, Z
func (c *CPU) TSX() byte {
	c.X = c.SP
	c.SetFlag(FlagZ, c.X == 0x00)
	c.SetFlag(FlagN, c.X&0x80 != 0)
	return 0
}

// Instruction: Transfer X Register to Stack Pointer
// Function: stack pointer = X
func (c *CPU) TXS() byte {
	c.SP = c.X
	return 0
}

// Instruction: Push Accumulator to Stack
// Function: A -> stack
func (c *CPU) PHA() byte {
	c.Write(0x0100+Bus.Word(c.SP), c.A)
	c.SP--
	return 0
}

// Instruction: Push Status Register to Stack
// Function: status -> stack
// Note: Break flag is set to 1 before push
func (c *CPU) PHP() byte {
	c.Write(0x0100+Bus.Word(c.SP), c.Status|FlagB|FlagU)
	c.SetFlag(FlagB, false)
	c.SetFlag(FlagU, false)
	c.SP--
	return 0
}

// Instruction: Pop Accumulator off Stack
// Function: A <- stack
// Flags Out: N, Z
func (c *CPU) PLA() byte {
	c.SP++
	c.Read(0x0100 + Bus.Word(c.SP))
	c.SetFlag(FlagZ, c.A == 0x00)
	c.SetFlag(FlagN, c.A&0x80 != 0)
	return 0
}

// Instruction: Pop Status Register off Stack
// Function: Status <- stack
func (c *CPU) PLP() byte {
	c.SP++
	c.Status = c.Read(0x0100 + Bus.Word(c.SP))
	c.SetFlag(FlagU, true)
	return 0
}

// Logical

// Instruction: Bitwise Logic AND
// Function: A = A & M
// Flags Out: N, Z
func (c *CPU) AND() byte {
	c.fetch()
	c.A = c.A & c.fetched
	c.SetFlag(FlagZ, c.A == 0x00)
	c.SetFlag(FlagN, c.A&0x80 != 0)
	return 1
}

// Instruction: Bitwise Logic XOR
// Function: A = A xor M
// Flags Out: N, Z
func (c *CPU) EOR() byte {
	c.fetch()
	c.A = c.A ^ c.fetched
	c.SetFlag(FlagZ, c.A == 0x00)
	c.SetFlag(FlagN, c.A&0x80 != 0)
	return 1
}

// Instruction: Bitwise Logic OR
// Function: A = A | M
// Flags Out: N, Z
func (c *CPU) ORA() byte {
	c.fetch()
	c.A = c.A | c.fetched
	c.SetFlag(FlagZ, c.A == 0x00)
	c.SetFlag(FlagN, c.A&0x80 != 0)
	return 1
}

// Instruction: Bit Test
// Function: A & M, N = M7, V = M6
// Flags Out: Z, V, N
func (c *CPU) BIT() byte {
	c.fetch()
	c.temp = Bus.Word(c.A) & Bus.Word(c.fetched)
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x00)
	c.SetFlag(FlagN, c.fetched&(1<<7) != 0)
	c.SetFlag(FlagV, c.fetched&(1<<6) != 0)
	return 0
}

// Arithmetic

// Instruction: Add with Carry In
// Function:    A = A + M + C
// Flags Out:   C, V, N, Z
func (c *CPU) ADC() byte {
	c.fetch()
	c.temp = Bus.Word(c.A) + Bus.Word(c.fetched)

	if c.GetFlag(FlagC) {
		c.temp += 1
	}

	c.SetFlag(FlagC, c.temp > 255)
	c.SetFlag(FlagV, Bus.Word(c.A)&Bus.Word(c.fetched)&Bus.Word(c.A)^Bus.Word(c.temp) != 0)
	c.SetFlag(FlagN, c.temp&0x80 != 0)
	c.A = byte(c.temp & 0x00FF)
	return 1
}

// Instruction: Subtraction with Borrow In
// Function:    A = A - M - (1 - C)
// Flags Out:   C, V, N, Z
func (c *CPU) SBC() byte {
	c.fetch()

	value := Bus.Word(c.fetched) ^ 0x00FF
	c.temp = Bus.Word(c.A) + value
	if c.GetFlag(FlagC) {
		c.temp += 1
	}
	c.SetFlag(FlagC, c.temp&0xFF00 != 0)
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0)
	c.SetFlag(FlagV, (c.temp^Bus.Word(c.A))&(c.temp&value) != 0)
	c.SetFlag(FlagN, c.temp&0x0080 != 0)
	c.A = byte(c.temp & 0x00FF)
	return 1
}

// Instruction: Compare Accumulator
// Function: C <- A >= M     Z <- (A - M) == 0
// Flags Out: N, C, Z
func (c *CPU) CMP() byte {
	c.fetch()
	c.temp = Bus.Word(c.A) - Bus.Word(c.fetched)
	c.SetFlag(FlagC, c.A >= c.fetched)
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x000)
	c.SetFlag(FlagN, c.temp&0x0080 != 0)
	return 1
}

// Instruction: Compare X Register
// Function: C <- X >= M      Z <- (X - M) == 0
// Flags Out: N, C, Z
func (c *CPU) CPX() byte {
	c.fetch()
	c.temp = Bus.Word(c.X) - Bus.Word(c.fetched)
	c.SetFlag(FlagC, c.X >= c.fetched)
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x0000)
	c.SetFlag(FlagN, c.temp&0x0080 != 0)
	return 0
}

// Instruction: Compare Y Register
// Function: C <- Y >= M      Z <- (Y - M) == 0
// Flags Out: N, C, Z
func (c *CPU) CPY() byte {
	c.fetch()
	c.temp = Bus.Word(c.Y) - Bus.Word(c.fetched)
	c.SetFlag(FlagC, c.Y >= c.fetched)
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x0000)
	c.SetFlag(FlagN, c.temp&0x0080 != 0)
	return 0
}

// Increments & Decrements

// Instruction: Increment Value at Memory Location
// Function: M = M + 1
// Flags Out: N, Z
func (c *CPU) INC() byte {
	c.fetch()
	c.temp = Bus.Word(c.fetched) + 1
	c.Write(c.addrAbs, byte(c.temp&0x00FF))
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x0000)
	c.SetFlag(FlagN, c.temp&0x0080 != 0)
	return 0
}

// Instruction: Increment X Register
// Function: X = X + 1
// Flags Out: N, Z
func (c *CPU) INX() byte {
	c.X++
	c.SetFlag(FlagZ, c.X == 0x00)
	c.SetFlag(FlagN, c.X&0x80 != 0)
	return 0
}

// Instruction: Increment Y Register
// Function: Y = Y + 1
// Flags Out: N, Z
func (c *CPU) INY() byte {
	c.Y++
	c.SetFlag(FlagZ, c.Y == 0x00)
	c.SetFlag(FlagN, c.Y&0x80 != 0)
	return 0
}

// Instruction: Decrement Value at Memory Location
// Function: M = M - 1
// Flags Out: N, Z
func (c *CPU) DEC() byte {
	c.fetch()
	c.temp = Bus.Word(c.fetched) - 1
	c.Write(c.addrAbs, byte(c.temp&0x00FF))
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x0000)
	c.SetFlag(FlagN, c.temp&0x0080 != 0)
	return 0
}

// Instruction: Decrement X Register
// Function: X = X - 1
// Flags Out: N, Z
func (c *CPU) DEX() byte {
	c.X--
	c.SetFlag(FlagZ, c.X == 0x00)
	c.SetFlag(FlagN, c.X&0x80 != 0)
	return 0
}

// Instruction: Decrement Y Register
// Function: Y = Y - 1
// Flags Out: N, Z
func (c *CPU) DEY() byte {
	c.Y--
	c.SetFlag(FlagZ, c.Y == 0x00)
	c.SetFlag(FlagN, c.Y&0x80 != 0)
	return 0
}

// Shifts
// Instruction: Arithmetic Shift Left
// Function:    A = C <- (A << 1) <- 0
// Flags Out:   N, Z, C
func (c *CPU) ASL() byte {
	c.fetch()
	c.temp = Bus.Word(c.fetched) << 1
	c.SetFlag(FlagC, (c.temp&0xFF00) > 0)
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x00)
	c.SetFlag(FlagN, c.temp&0x80 != 0)
	if c.currentInstr.isImplied {
		c.A = byte(c.temp & 0x00FF)
	} else {
		c.Write(c.addrAbs, byte(c.temp&0x00FF))
	}
	return 0
}

// Instruction: Logical Shift Right
// Function:    A = 0 -> (A >> 1) -> C
// Flags Out: Z, N
func (c *CPU) LSR() byte {
	c.fetch()
	c.SetFlag(FlagC, c.fetched&0x0001 != 0)
	c.temp = Bus.Word(c.fetched) >> 1
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x00)
	c.SetFlag(FlagN, c.temp&0x80 != 0)
	if c.currentInstr.isImplied {
		c.A = byte(c.temp & 0x00FF)
	} else {
		c.Write(c.addrAbs, byte(c.temp&0x00FF))
	}
	return 0
}

// Instruction: Rotate Left
// Function: A = (A << 1) <- C
// Flags Out: Z, N
func (c *CPU) ROL() byte {
	c.fetch()
	c.temp = Bus.Word(c.fetched << 1)
	if c.GetFlag(FlagC) {
		c.temp |= 1
	}
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x00)
	c.SetFlag(FlagN, c.temp&0x0080 != 0)
	if c.currentInstr.isImplied {
		c.A = byte(c.temp & 0x00FF) // Probably & 0x00FF it is useless
	} else {
		c.Write(c.addrAbs, byte(c.temp&0x00FF))
	}
	return 0
}

// Instruction: Rotate Left
// Function: A = C -> (A >> 1)
// Flags Out: Z, N
func (c *CPU) ROR() byte {
	c.fetch()
	if c.GetFlag(FlagC) {
		c.temp = 1 << 7
	}
	c.temp |= Bus.Word(c.fetched >> 1)
	c.SetFlag(FlagZ, (c.temp&0x00FF) == 0x00)
	c.SetFlag(FlagN, c.temp&0x0080 != 0)
	if c.currentInstr.isImplied {
		c.A = byte(c.temp & 0x00FF)
	} else {
		c.Write(c.addrAbs, byte(c.temp&0x00FF))
	}
	return 0
}

// Jumps & Calls
// Instruction: Jump To Location
// Function: pc = address
func (c *CPU) JMP() byte {
	c.PC = c.addrAbs
	return 0
}

// Instruction: Jump To Sub-Routine
// Function: Push current pc to stack, pc = address
func (c *CPU) JSR() byte {
	c.PC--
	c.Write(0x0100+Bus.Word(c.SP), byte(c.PC>>8)&0x00FF)
	c.SP--
	c.Write(0x0100+Bus.Word(c.SP), byte(c.PC>>8)&0x00FF)
	c.SP--

	c.PC = c.addrAbs
	return 0
}

// Instruction: Return from Subroutine
// Function: Pulls pc from stack
func (c *CPU) RTS() byte {
	c.SP++
	c.PC = Bus.Word(c.Read(0x0100 + Bus.Word(c.SP)))
	c.SP++
	c.PC |= Bus.Word(c.Read(0x0100 + Bus.Word(c.SP)))

	c.PC++
	return 0
}

// Branches
// Instruction: Branch if Carry Clear
// Function: if(C == 0) pc = address
func (c *CPU) BCC() byte {
	if !c.GetFlag(FlagC) {
		c.cycles++
		c.addrAbs = c.PC + c.addrRel

		if c.addrAbs&0x00FF != (c.PC & 0xFF00) {
			c.cycles++
		}
		c.PC = c.addrAbs

	}
	return 0
}

// Instruction: Branch if Carry Set
// Function: if(C == 1) pc = address
func (c *CPU) BCS() byte {
	if c.GetFlag(FlagC) {
		c.cycles++
		c.addrAbs = c.PC + c.addrRel

		if c.addrAbs&0x00FF != (c.PC & 0xFF00) {
			c.cycles++
		}
		c.PC = c.addrAbs

	}
	return 0
}

// Instruction: Branch if Equal
// Function: if(Z == 1) pc = address
func (c *CPU) BEQ() byte {
	if c.GetFlag(FlagZ) {
		c.cycles++
		c.addrAbs = c.PC + c.addrRel

		if c.addrAbs&0x00FF != (c.PC & 0xFF00) {
			c.cycles++
		}
		c.PC = c.addrAbs

	}
	return 0
}

// Instruction: Branch if Negative
// Function: if(N == 1) pc = address
func (c *CPU) BMI() byte {
	if c.GetFlag(FlagN) {
		c.cycles++
		c.addrAbs = c.PC + c.addrRel

		if c.addrAbs&0x00FF != (c.PC & 0xFF00) {
			c.cycles++
		}
		c.PC = c.addrAbs

	}
	return 0
}

// Instruction: Branch if Not Equal
// Function: if(Z == 0) pc = address
func (c *CPU) BNE() byte {
	if !c.GetFlag(FlagZ) {
		c.cycles++
		c.addrAbs = c.PC + c.addrRel

		if c.addrAbs&0x00FF != (c.PC & 0xFF00) {
			c.cycles++
		}
		c.PC = c.addrAbs

	}
	return 0
}

// Instruction: Branch if Not Equal
// Function: if(Z == 0) pc = address
func (c *CPU) BPL() byte {
	if !c.GetFlag(FlagN) {
		c.cycles++
		c.addrAbs = c.PC + c.addrRel

		if c.addrAbs&0x00FF != (c.PC & 0xFF00) {
			c.cycles++
		}
		c.PC = c.addrAbs

	}
	return 0
}

// Instruction: Branch if Overflow Clear
// Function: if(V == 0) pc = address
func (c *CPU) BVC() byte {
	if !c.GetFlag(FlagV) {
		c.cycles++
		c.addrAbs = c.PC + c.addrRel

		if c.addrAbs&0x00FF != (c.PC & 0xFF00) {
			c.cycles++
		}
		c.PC = c.addrAbs

	}
	return 0
}

// Instruction: Branch if Overflow Set
// Function: if(V == 1) pc = address
func (c *CPU) BVS() byte {
	if c.GetFlag(FlagV) {
		c.cycles++
		c.addrAbs = c.PC + c.addrRel

		if c.addrAbs&0x00FF != (c.PC & 0xFF00) {
			c.cycles++
		}
		c.PC = c.addrAbs

	}
	return 0
}

// Status Flag Changes

// Instruction: Clear Carry Flag
// Function: C = 0
func (c *CPU) CLC() byte {
	c.SetFlag(FlagC, false)
	return 0
}

// Instruction: Clear Decimal Flag
// Function: D = 0
func (c *CPU) CLD() byte {
	c.SetFlag(FlagD, false)
	return 0
}

// Instruction: Disable Interrupts / Clear Interrupt Flag
// Function: I = 0
func (c *CPU) CLI() byte {
	c.SetFlag(FlagI, false)
	return 0
}

// Instruction: Clear Overflow Flag
// Function: V = 0
func (c *CPU) CLV() byte {
	c.SetFlag(FlagV, false)
	return 0
}

// Instruction: Set Carry Flag
// Function: C = 1
func (c *CPU) SEC() byte {
	c.SetFlag(FlagC, false)
	return 0
}

// Instruction: Set Decimal Flag
// Function: D = 1
func (c *CPU) SED() byte {
	c.SetFlag(FlagD, false)
	return 0
}

// Instruction: Set Interrupt Flag / Enable Interrupts
// Function: I = 1
func (c *CPU) SEI() byte {
	c.SetFlag(FlagI, false)
	return 0
}

// Instruction: Break
// Function: Program Sourced Interrupt
func (c *CPU) BRK() byte {
	c.PC++

	c.SetFlag(FlagI, true)
	c.Write(0x0100+Bus.Word(c.SP), byte(Bus.Word(c.PC)>>8))
	c.SP--
	c.Write(0x0100+Bus.Word(c.SP), byte(Bus.Word(c.PC)))
	c.SP--

	c.SetFlag(FlagB, true)
	c.Write(0x0100+Bus.Word(c.SP), c.Status)
	c.SP--
	c.SetFlag(FlagB, false)

	c.PC = Bus.Word(Bus.Word(c.Read(0xFFFE)) | Bus.Word(c.Read(0xFFFF))<<8)

	return 0
}

// Instruction: No Operation
func (c *CPU) NOP() byte {
	// Architecture that the all NOPS are equal.
	// More info:
	// https://www.nesdev.org/wiki/CPU_unofficial_opcodes
	switch c.opcode {
	case 0x1C:
	case 0x3C:
	case 0x5C:
	case 0x7C:
	case 0xDC:
	case 0xFC:
		return 1
	}
	return 0
}

// Instruction: Return from Interrupt
// Function: Pulls the processor flags from the stack
func (c *CPU) RTI() byte {
	c.SP++
	c.Status = c.Read(0x0100 + Bus.Word(c.SP))
	c.Status &= FlagB
	c.Status &= FlagU

	c.SP++
	c.PC = Bus.Word(c.Read(0x0100 + Bus.Word(c.SP)))
	c.SP++
	c.PC |= Bus.Word(c.Read(0x0100 + Bus.Word(c.SP)<<8))
	return 0
}

// This function captures illegal opcodees
func (c *CPU) XXX() byte {
	return 0
}

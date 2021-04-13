package cpu

import (
	"fmt"
	"reflect"
	"runtime"
)

type instruction struct {
	I func(*CPU, Memory)
	C int
	B uint16
}

func handler(f func(*CPU, Memory, MemoryMode), mode MemoryMode) func(*CPU, Memory) {
	return func(c *CPU, m Memory) {
		f(c, m, mode)
	}
}

func (i instruction) String() string {
	return fmt.Sprintf("Instruction %v\n", runtime.FuncForPC(reflect.ValueOf(i.I).Pointer()).Name())
}

func (i instruction) Execute(c *CPU, clock *Clock, m Memory) bool {
	c.PC++
	i.I(c, m)
	if !clock.Tick(i.C) {
		fmt.Printf("No more clock ticks\n")
		return false
	}
	return true
}

var fastLookup = make([]*instruction, 256)

func init() {
	for i := 0; i <= 255; i++ {
		fastLookup[i] = instructionTable[uint8(i)]
	}
	if EnableLogging {
		defaultLogger = &logger{
			buffer: make([]*state, 100),
			len:    100,
			head:   0,
		}

	}
}

var instructionTable = map[byte]*instruction{
	0xA9: {handler(LDA, addrI), 2, 2},
	0xA5: {handler(LDA, addrZP), 3, 2},
	0xB5: {handler(LDA, addrZPX), 4, 2},
	0xAD: {handler(LDA, addrA), 4, 3},
	0xBD: {handler(LDA, addrAX), 4, 3},
	0xB9: {handler(LDA, addrAY), 4, 3},
	0xA1: {handler(LDA, addrIX), 6, 2},
	0xB1: {handler(LDA, addrIY), 5, 2},
	0xA2: {handler(LDX, addrI), 2, 2},
	0xA6: {handler(LDX, addrZP), 3, 2},
	0xB6: {handler(LDX, addrZPY), 4, 2},
	0xAE: {handler(LDX, addrA), 4, 3},
	0xBE: {handler(LDX, addrAY), 4, 3},
	0xA0: {handler(LDY, addrI), 2, 2},
	0xA4: {handler(LDY, addrZP), 3, 2},
	0xB4: {handler(LDY, addrZPX), 4, 2},
	0xAC: {handler(LDY, addrA), 4, 3},
	0xBC: {handler(LDY, addrAX), 4, 3},
	0x85: {handler(STA, addrZP), 3, 2},
	0x95: {handler(STA, addrZPX), 4, 2},
	0x8D: {handler(STA, addrA), 4, 3},
	0x9D: {handler(STA, addrAX), 5, 3},
	0x99: {handler(STA, addrAY), 5, 3},
	0x81: {handler(STA, addrIX), 6, 2},
	0x91: {handler(STA, addrIY), 6, 2},
	0x86: {handler(STX, addrZP), 3, 2},
	0x96: {handler(STX, addrZPY), 4, 2},
	0x8E: {handler(STX, addrA), 4, 3},
	0x84: {handler(STY, addrZP), 3, 2},
	0x94: {handler(STY, addrZPX), 4, 2},
	0x8C: {handler(STY, addrA), 4, 3},
	0x69: {handler(ADC, addrI), 2, 2},
	0x65: {handler(ADC, addrZP), 3, 2},
	0x75: {handler(ADC, addrZPX), 4, 2},
	0x6D: {handler(ADC, addrA), 4, 3},
	0x7D: {handler(ADC, addrAX), 4, 3},
	0x79: {handler(ADC, addrAY), 4, 3},
	0x61: {handler(ADC, addrIX), 6, 2},
	0x71: {handler(ADC, addrIY), 5, 2},
	0x29: {handler(AND, addrI), 2, 2},
	0x25: {handler(AND, addrZP), 3, 2},
	0x35: {handler(AND, addrZPX), 4, 2},
	0x2D: {handler(AND, addrA), 4, 3},
	0x3D: {handler(AND, addrAX), 4, 3},
	0x39: {handler(AND, addrAY), 4, 3},
	0x21: {handler(AND, addrIX), 6, 2},
	0x31: {handler(AND, addrIY), 5, 2},
	0x0A: {ASLAccumulator, 2, 2},
	0x06: {handler(ASL, addrZP), 5, 2},
	0x16: {handler(ASL, addrZPX), 6, 2},
	0x0E: {handler(ASL, addrA), 6, 3},
	0x1E: {handler(ASL, addrAX), 7, 3},
	0x90: {BCC, 2, 2},
	0xB0: {BCS, 2, 2},
	0xF0: {BEQ, 2, 2},
	0x24: {handler(BIT, addrZP), 3, 2},
	0x2C: {handler(BIT, addrA), 4, 3},
	0x30: {BMI, 2, 2},
	0xD0: {BNE, 2, 2},
	0x10: {BPL, 2, 2},
	0x50: {BVC, 2, 2},
	0x70: {BVS, 2, 2},
	0x00: {BRK, 7, 1},
	0x18: {CLC, 2, 1},
	0xD8: {CLD, 2, 1},
	0x58: {CLI, 2, 1},
	0xB8: {CLV, 2, 1},
	0x38: {SEC, 2, 1},
	0xF8: {SED, 2, 1},
	0x78: {SEI, 2, 1},
	0xC9: {handler(CMP, addrI), 2, 2},
	0xC5: {handler(CMP, addrZP), 3, 2},
	0xD5: {handler(CMP, addrZPX), 3, 2},
	0xCD: {handler(CMP, addrA), 4, 3},
	0xDD: {handler(CMP, addrAX), 4, 3},
	0xD9: {handler(CMP, addrAY), 4, 3},
	0xC1: {handler(CMP, addrIX), 6, 2},
	0xD1: {handler(CMP, addrIY), 5, 2},
	0xE0: {handler(CPX, addrI), 2, 2},
	0xE4: {handler(CPX, addrZP), 3, 2},
	0xEC: {handler(CPX, addrA), 4, 3},
	0xC0: {handler(CPY, addrI), 2, 2},
	0xC4: {handler(CPY, addrZP), 3, 2},
	0xCC: {handler(CPY, addrA), 4, 3},
	0xC6: {handler(DEC, addrZP), 5, 2},
	0xD6: {handler(DEC, addrZPX), 6, 2},
	0xCE: {handler(DEC, addrA), 6, 3},
	0xDE: {handler(DEC, addrAX), 7, 3},
	0xCA: {DEX, 2, 1},
	0x88: {DEY, 2, 1},
	0xE6: {handler(INC, addrZP), 5, 2},
	0xF6: {handler(INC, addrZPX), 6, 2},
	0xEE: {handler(INC, addrA), 6, 3},
	0xFE: {handler(INC, addrAX), 7, 3},
	0xE8: {INX, 2, 1},
	0xC8: {INY, 2, 1},
	0x4C: {JMPA, 3, 3},
	0x6C: {JMPI, 5, 3},
	0x20: {JSR, 6, 3},
	0x49: {handler(EOR, addrI), 2, 2},
	0x45: {handler(EOR, addrZP), 3, 2},
	0x55: {handler(EOR, addrZPX), 4, 2},
	0x4D: {handler(EOR, addrA), 4, 3},
	0x5D: {handler(EOR, addrAX), 4, 3},
	0x59: {handler(EOR, addrAY), 4, 3},
	0x41: {handler(EOR, addrIX), 6, 2},
	0x51: {handler(EOR, addrIY), 5, 2},
	0x4A: {LSRAccumulator, 2, 1},
	0x46: {handler(LSR, addrZP), 5, 2},
	0x56: {handler(LSR, addrZPX), 6, 2},
	0x4E: {handler(LSR, addrA), 6, 3},
	0x5E: {handler(LSR, addrAX), 7, 3},
	0xEA: {func(*CPU, Memory) {}, 2, 1}, // NOP
	0x09: {handler(ORA, addrI), 2, 2},
	0x05: {handler(ORA, addrZP), 3, 2},
	0x15: {handler(ORA, addrZPX), 4, 2},
	0x0D: {handler(ORA, addrA), 4, 3},
	0x1D: {handler(ORA, addrAX), 4, 3},
	0x19: {handler(ORA, addrAY), 4, 3},
	0x01: {handler(ORA, addrIX), 6, 2},
	0x11: {handler(ORA, addrIY), 5, 2},
	0x48: {PHA, 3, 1},
	0x08: {PHP, 3, 1},
	0x68: {PLA, 4, 1},
	0x28: {PLP, 4, 1},
	0x2A: {ROLAccumulator, 2, 1},
	0x26: {handler(ROL, addrZP), 5, 2},
	0x36: {handler(ROL, addrZPX), 6, 2},
	0x2E: {handler(ROL, addrA), 6, 3},
	0x3E: {handler(ROL, addrAX), 7, 3},
	0x6A: {RORAccumulator, 2, 1},
	0x66: {handler(ROR, addrZP), 5, 2},
	0x76: {handler(ROR, addrZPX), 6, 2},
	0x6E: {handler(ROR, addrA), 6, 3},
	0x7E: {handler(ROR, addrAX), 7, 3},
	0x40: {RTI, 6, 1},
	0x60: {RTS, 6, 1},
	0xE9: {handler(SUB, addrI), 2, 2},
	0xE5: {handler(SUB, addrZP), 3, 2},
	0xF5: {handler(SUB, addrZPX), 4, 2},
	0xED: {handler(SUB, addrA), 4, 3},
	0xFD: {handler(SUB, addrAX), 4, 3},
	0xF9: {handler(SUB, addrAY), 4, 3},
	0xE1: {handler(SUB, addrIX), 6, 2},
	0xF1: {handler(SUB, addrIY), 5, 2},
	0xAA: {TAX, 2, 1},
	0xA8: {TAY, 2, 1},
	0xBA: {TSX, 2, 1},
	0x8A: {TXA, 2, 1},
	0x9A: {TXS, 2, 1},
	0x98: {TYA, 2, 1},
}

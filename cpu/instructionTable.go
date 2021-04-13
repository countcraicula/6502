package cpu

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type instruction struct {
	I func(*CPU, Memory, MemoryMode)
	M MemoryMode
	C int
	B uint16
}

func (i *instruction) String() string {
	return fmt.Sprintf("Instruction %v %v\n", i.Name(), i.Mode())
}

func (i *instruction) Name() string {
	if i == nil {
		return "ILL"
	}
	s := strings.Split(runtime.FuncForPC(reflect.ValueOf(i.I).Pointer()).Name(), ".")
	return s[1]
}

func (i *instruction) Mode() string {
	if i == nil {
		return "ILL"
	}
	s := strings.Split(runtime.FuncForPC(reflect.ValueOf(i.M).Pointer()).Name(), ".")
	return s[1]
}

func (i instruction) Execute(c *CPU, clock *Clock, m Memory) bool {
	c.PC++
	i.I(c, m, i.M)
	if !clock.Tick(i.C) {
		fmt.Printf("No more clock ticks\n")
		return false
	}
	return true
}

var fastLookup = make([]*instruction, 256)

var illegalInstruction *instruction

func init() {
	for i := 0; i <= 255; i++ {
		fastLookup[i] = illegalInstruction
	}
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
	0x00: {BRK, addrNone, 7, 1},
	0x01: {ORA, addrIX, 6, 2},
	0x02: {CLE, addrNone, 2, 1},
	0x03: {SEE, addrNone, 2, 1},
	0x04: {TSB, addrZP, 5, 2},
	0x05: {ORA, addrZP, 3, 2},
	0x06: {ASL, addrZP, 5, 2},
	0x07: {RMB, addrNone, 2, 3},
	0x08: {PHP, addrNone, 3, 1},
	0x09: {ORA, addrI, 2, 2},
	0x0A: {ASLAccumulator, addrNone, 2, 2},
	0x0B: {TSY, addrNone, 2, 1},
	0x0C: {TSB, addrA, 6, 3},
	0x0D: {ORA, addrA, 4, 3},
	0x0E: {ASL, addrA, 6, 3},
	0x0F: {BBR, addrNone, 2, 3},
	0x10: {BPL, addrNone, 2, 2},
	0x11: {ORA, addrIY, 5, 2},
	0x12: {ORA, addrIZ, 5, 2},
	0x13: {BPL16, addrNone, 2, 3},
	0x14: {TRB, addrZP, 5, 2},
	0x15: {ORA, addrZPX, 4, 2},
	0x16: {ASL, addrZPX, 6, 2},
	0x17: {RMB, addrNone, 2, 3},
	0x18: {CLC, addrNone, 2, 1},
	0x19: {ORA, addrAY, 4, 3},
	0x1A: {INCAccumulator, addrNone, 2, 1},
	0x1B: {INZ, addrNone, 2, 1},
	0x1C: {TRB, addrA, 6, 3},
	0x1D: {ORA, addrAX, 4, 3},
	0x1E: {ASL, addrAX, 7, 3},
	0x1F: {BBR, addrNone, 2, 3},
	0x20: {JSR, addrI, 6, 3},
	0x21: {AND, addrIX, 6, 2},
	0x22: {JSR, addrIN, 6, 2},
	0x23: {JSR, addrAIX, 6, 2},
	0x24: {BIT, addrZP, 3, 2},
	0x25: {AND, addrZP, 3, 2},
	0x26: {ROL, addrZP, 5, 2},
	0x27: {RMB, addrNone, 2, 3},
	0x28: {PLP, addrNone, 4, 1},
	0x29: {AND, addrI, 2, 2},
	0x2A: {ROLAccumulator, addrNone, 2, 1},
	0x2B: {TYS, addrNone, 2, 1},
	0x2C: {BIT, addrA, 4, 3},
	0x2D: {AND, addrA, 4, 3},
	0x2E: {ROL, addrA, 6, 3},
	0x2F: {BBR, addrNone, 2, 3},
	0x30: {BMI, addrNone, 2, 2},
	0x31: {AND, addrIY, 5, 2},
	0x32: {AND, addrIZ, 5, 2},
	0x33: {BMI16, addrNone, 2, 3},
	0x34: {BIT, addrZPX, 4, 2},
	0x35: {AND, addrZPX, 4, 2},
	0x36: {ROL, addrZPX, 6, 2},
	0x37: {RMB, addrNone, 2, 3},
	0x38: {SEC, addrNone, 2, 1},
	0x39: {AND, addrAY, 4, 3},
	0x3A: {DECAccumulator, addrNone, 2, 1},
	0x3B: {DEZ, addrNone, 2, 1},
	0x3C: {BIT, addrAX, 4, 3},
	0x3D: {AND, addrAX, 4, 3},
	0x3E: {ROL, addrAX, 7, 3},
	0x3F: {BBR, addrNone, 2, 3},
	0x40: {RTI, addrNone, 6, 1},
	0x41: {EOR, addrIX, 6, 2},
	0x42: {NEG, addrNone, 2, 1},
	0x43: {ASRAcculuator, addrNone, 2, 1},
	0x44: {ASR, addrZP, 5, 2},
	0x45: {EOR, addrZP, 3, 2},
	0x46: {LSR, addrZP, 5, 2},
	0x47: {RMB, addrNone, 2, 3},
	0x48: {PHA, addrNone, 3, 1},
	0x49: {EOR, addrI, 2, 2},
	0x4A: {LSRAccumulator, addrNone, 2, 1},
	0x4B: {TAZ, addrNone, 2, 1},
	0x4C: {JMP, addrA, 3, 3},
	0x4D: {EOR, addrA, 4, 3},
	0x4E: {LSR, addrA, 6, 3},
	0x4F: {BBR, addrNone, 2, 3},
	0x50: {BVC, addrNone, 2, 2},
	0x51: {EOR, addrIY, 5, 2},
	0x52: {EOR, addrIZ, 5, 2},
	0x53: {BVC16, addrNone, 2, 3},
	0x54: {ASR, addrZPX, 6, 2},
	0x55: {EOR, addrZPX, 4, 2},
	0x56: {LSR, addrZPX, 6, 2},
	0x57: {RMB, addrNone, 2, 3},
	0x58: {CLI, addrNone, 2, 1},
	0x59: {EOR, addrAY, 4, 3},
	0x5A: {PHY, addrNone, 4, 1},
	0x5B: {TAB, addrNone, 2, 1},
	0x5C: {NOP, addrNone, 2, 1},
	0x5D: {EOR, addrAX, 4, 3},
	0x5E: {LSR, addrAX, 7, 3},
	0x5F: {BBR, addrNone, 2, 3},
	0x60: {RTS, addrNone, 6, 1},
	0x61: {ADC, addrIX, 6, 2},
	0x62: {RTN, addrI, 6, 2},
	0x63: {JSR, addrRel, 3, 3},
	0x64: {STZ, addrZP, 3, 2},
	0x65: {ADC, addrZP, 3, 2},
	0x66: {ROR, addrZP, 5, 2},
	0x67: {RMB, addrNone, 2, 3},
	0x68: {PLA, addrNone, 4, 1},
	0x69: {ADC, addrI, 2, 2},
	0x6A: {RORAccumulator, addrNone, 2, 1},
	0x6B: {TZA, addrNone, 2, 1},
	0x6C: {JMP, addrIN, 5, 3},
	0x6D: {ADC, addrA, 4, 3},
	0x6E: {ROR, addrA, 6, 3},
	0x6F: {BBR, addrNone, 2, 3},
	0x70: {BVS, addrNone, 2, 2},
	0x71: {ADC, addrIY, 5, 2},
	0x72: {ADC, addrIZ, 5, 2},
	0x73: {BVS16, addrNone, 2, 3},
	0x74: {STZ, addrZPX, 3, 2},
	0x75: {ADC, addrZPX, 4, 2},
	0x76: {ROR, addrZPX, 6, 2},
	0x77: {RMB, addrNone, 2, 3},
	0x78: {SEI, addrNone, 2, 1},
	0x79: {ADC, addrAY, 4, 3},
	0x7A: {PLY, addrNone, 4, 1},
	0x7B: {TBA, addrNone, 2, 1},
	0x7C: {JMP, addrAIX, 6, 3},
	0x7D: {ADC, addrAX, 4, 3},
	0x7E: {ROR, addrAX, 7, 3},
	0x7F: {BBR, addrNone, 2, 3},
	0x80: {BRA, addrNone, 3, 2},
	0x81: {STA, addrIX, 6, 2},
	0x82: {STA, addrISY, 6, 2},
	0x83: {BRA16, addrNone, 4, 3},
	0x84: {STY, addrZP, 3, 2},
	0x85: {STA, addrZP, 3, 2},
	0x86: {STX, addrZP, 3, 2},
	0x87: {SMB, addrNone, 2, 3},
	0x88: {DEY, addrNone, 2, 1},
	0x89: {BIT, addrI, 2, 2},
	0x8A: {TXA, addrNone, 2, 1},
	0x8B: {STY, addrAX, 4, 3},
	0x8C: {STY, addrA, 4, 3},
	0x8D: {STA, addrA, 4, 3},
	0x8E: {STX, addrA, 4, 3},
	0x8F: {BBS, addrNone, 2, 3},
	0x90: {BCC, addrNone, 2, 2},
	0x91: {STA, addrIY, 6, 2},
	0x92: {STA, addrIZ, 6, 2},
	0x93: {BCC16, addrNone, 2, 2},
	0x94: {STY, addrZPX, 4, 2},
	0x95: {STA, addrZPX, 4, 2},
	0x96: {STX, addrZPY, 4, 2},
	0x97: {SMB, addrNone, 2, 3},
	0x98: {TYA, addrNone, 2, 1},
	0x99: {STA, addrAY, 5, 3},
	0x9A: {TXS, addrNone, 2, 1},
	0x9B: {STX, addrAY, 4, 3},
	0x9C: {STZ, addrA, 3, 2},
	0x9D: {STA, addrAX, 5, 3},
	0x9E: {STZ, addrAX, 3, 2},
	0x9F: {BBS, addrNone, 2, 3},
	0xA0: {LDY, addrI, 2, 2},
	0xA1: {LDA, addrIX, 6, 2},
	0xA2: {LDX, addrI, 2, 2},
	0xA3: {LDZ, addrI, 2, 2},
	0xA4: {LDY, addrZP, 3, 2},
	0xA5: {LDA, addrZP, 3, 2},
	0xA6: {LDX, addrZP, 3, 2},
	0xA7: {SMB, addrNone, 2, 3},
	0xA8: {TAY, addrNone, 2, 1},
	0xA9: {LDA, addrI, 2, 2},
	0xAA: {TAX, addrNone, 2, 1},
	0xAB: {LDZ, addrA, 4, 3},
	0xAC: {LDY, addrA, 4, 3},
	0xAD: {LDA, addrA, 4, 3},
	0xAE: {LDX, addrA, 4, 3},
	0xAF: {BBS, addrNone, 2, 3},
	0xB0: {BCS, addrNone, 2, 2},
	0xB1: {LDA, addrIY, 5, 2},
	0xB2: {LDA, addrIZ, 5, 2},
	0xB3: {BCS16, addrNone, 2, 2},
	0xB4: {LDY, addrZPX, 4, 2},
	0xB5: {LDA, addrZPX, 4, 2},
	0xB6: {LDX, addrZPY, 4, 2},
	0xB7: {SMB, addrNone, 2, 3},
	0xB8: {CLV, addrNone, 2, 1},
	0xB9: {LDA, addrAY, 4, 3},
	0xBA: {TSX, addrNone, 2, 1},
	0xBB: {LDZ, addrAX, 4, 3},
	0xBC: {LDY, addrAX, 4, 3},
	0xBD: {LDA, addrAX, 4, 3},
	0xBE: {LDX, addrAY, 4, 3},
	0xBF: {BBS, addrNone, 2, 3},
	0xC0: {CPY, addrI, 2, 2},
	0xC1: {CMP, addrIX, 6, 2},
	0xC2: {CPZ, addrI, 2, 2},
	0xC3: {DEW, addrZP, 6, 2},
	0xC4: {CPY, addrZP, 3, 2},
	0xC5: {CMP, addrZP, 3, 2},
	0xC6: {DEC, addrZP, 5, 2},
	0xC7: {SMB, addrNone, 2, 3},
	0xC8: {INY, addrNone, 2, 1},
	0xC9: {CMP, addrI, 2, 2},
	0xCA: {DEX, addrNone, 2, 1},
	0xCB: {ASW, addrA, 7, 3},
	0xCC: {CPY, addrA, 4, 3},
	0xCD: {CMP, addrA, 4, 3},
	0xCE: {DEC, addrA, 6, 3},
	0xCF: {BBS, addrNone, 2, 3},
	0xD0: {BNE, addrNone, 2, 2},
	0xD1: {CMP, addrIY, 5, 2},
	0xD2: {CMP, addrIZ, 5, 2},
	0xD3: {BNE16, addrNone, 2, 3},
	0xD4: {CPZ, addrZP, 3, 2},
	0xD5: {CMP, addrZPX, 3, 2},
	0xD6: {DEC, addrZPX, 6, 2},
	0xD7: {SMB, addrNone, 2, 3},
	0xD8: {CLD, addrNone, 2, 1},
	0xD9: {CMP, addrAY, 4, 3},
	0xDA: {PHX, addrNone, 4, 1},
	0xDB: {PHZ, addrNone, 4, 1},
	0xDC: {CPZ, addrA, 4, 3},
	0xDD: {CMP, addrAX, 4, 3},
	0xDE: {DEC, addrAX, 7, 3},
	0xDF: {BBS, addrNone, 2, 3},
	0xE0: {CPX, addrI, 2, 2},
	0xE1: {SUB, addrIX, 6, 2},
	0xE2: {LDA, addrISY, 5, 2},
	0xE3: {INW, addrZP, 6, 2},
	0xE4: {CPX, addrZP, 3, 2},
	0xE5: {SUB, addrZP, 3, 2},
	0xE6: {INC, addrZP, 5, 2},
	0xE7: {SMB, addrNone, 2, 3},
	0xE8: {INX, addrNone, 2, 1},
	0xE9: {SUB, addrI, 2, 2},
	0xEA: {NOP, addrNone, 2, 1},
	0xEB: {ROW, addrA, 7, 3},
	0xEC: {CPX, addrA, 4, 3},
	0xED: {SUB, addrA, 4, 3},
	0xEE: {INC, addrA, 6, 3},
	0xEF: {BBS, addrNone, 2, 3},
	0xF0: {BEQ, addrNone, 2, 2},
	0xF1: {SUB, addrIY, 5, 2},
	0xF2: {SUB, addrIZ, 5, 2},
	0xF4: {PHW, addrI16, 4, 3},
	0xF5: {SUB, addrZPX, 4, 2},
	0xF6: {INC, addrZPX, 6, 2},
	0xF7: {SMB, addrNone, 2, 3},
	0xF8: {SED, addrNone, 2, 1},
	0xF9: {SUB, addrAY, 4, 3},
	0xFA: {PLX, addrNone, 4, 1},
	0xFB: {PLZ, addrNone, 4, 1},
	0xFC: {PHW, addrA, 4, 3},
	0xFD: {SUB, addrAX, 4, 3},
	0xFE: {INC, addrAX, 7, 3},
	0xFF: {BBS, addrNone, 2, 3},
}

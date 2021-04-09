package cpu

type instruction struct {
	I func(*CPU, Memory)
	C uint
	B uint16
}

func (i instruction) Execute(c *CPU, clock Clock, m Memory) bool {
	if !clock.Tick(i.C) {
		return false
	}
	i.I(c, m)
	c.PC += i.B
	return true
}

var instructionTable = map[byte]instruction{
	0xA9: {LDAImmediate, 2, 2},
	0xA5: {LDAZP, 3, 2},
	0xB5: {LDAZPX, 4, 2},
	0xAD: {LDAA, 4, 3},
	0xBD: {LDAAX, 4, 3},
	0xB9: {LDAAY, 4, 3},
	0xA1: {LDAIX, 6, 2},
	0xB1: {LDAIY, 5, 2},
	0xA2: {LDXImmediate, 2, 2},
	0xA6: {LDXZP, 3, 2},
	0xB6: {LDXZPY, 4, 2},
	0xAE: {LDXA, 4, 3},
	0xBE: {LDXAY, 4, 3},
	0xA0: {LDYImmediate, 2, 2},
	0xA4: {LDYZP, 3, 2},
	0xB4: {LDYZPX, 4, 2},
	0xAC: {LDYA, 4, 3},
	0xBC: {LDYAX, 4, 3},
	0x85: {STAZP, 3, 2},
	0x95: {STAZPX, 4, 2},
	0x8D: {STAA, 4, 3},
	0x9D: {STAAX, 5, 3},
	0x99: {STAAY, 5, 3},
	0x81: {STAIX, 6, 2},
	0x91: {STAIY, 6, 2},
	0x86: {STXZP, 3, 2},
	0x96: {STXZPY, 4, 2},
	0x8E: {STXA, 4, 3},
	0x84: {STYZP, 3, 2},
	0x94: {STYZPX, 4, 2},
	0x8C: {STYA, 4, 3},
	0x69: {ADCImmediate, 2, 2},
	0x65: {ADCZP, 3, 2},
	0x75: {ADCZPX, 4, 2},
	0x6D: {ADCA, 4, 3},
	0x7D: {ADCAX, 4, 3},
	0x79: {ADCAY, 4, 3},
	0x61: {ADCIX, 6, 2},
	0x71: {ADCIY, 5, 2},
	0x29: {ANDImmediate, 2, 2},
	0x25: {ANDZP, 3, 2},
	0x35: {ANDZPX, 4, 2},
	0x2D: {ANDA, 4, 3},
	0x3D: {ANDAX, 4, 3},
	0x39: {ANDAY, 4, 3},
	0x21: {ANDIX, 6, 2},
	0x31: {ANDIY, 5, 2},
	0x0A: {ASLAccumulator, 2, 2},
	0x06: {ASLZP, 5, 2},
	0x16: {ASLZPX, 6, 2},
	0x0E: {ASLA, 6, 3},
	0x1E: {ASLAX, 7, 3},
	0x90: {BCC, 2, 2},
	0xB0: {BCS, 2, 2},
	0xF0: {BEQ, 2, 2},
	0x24: {BITZP, 3, 2},
	0x2C: {BITA, 4, 3},
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
	0xC9: {CMPImmediate, 2, 2},
	0xC5: {CMPZP, 3, 2},
	0xD5: {CMPZPX, 3, 2},
	0xCD: {CMPA, 4, 3},
	0xDD: {CMPAX, 4, 3},
	0xD9: {CMPAY, 4, 3},
	0xC1: {CMPIX, 6, 2},
	0xD1: {CMPIY, 5, 2},
	0xE0: {CPXImmediate, 2, 2},
	0xE4: {CPXZP, 3, 2},
	0xEC: {CPXA, 4, 3},
	0xC0: {CPYImmediate, 2, 2},
	0xC4: {CPYZP, 3, 2},
	0xCC: {CPYA, 4, 3},
	0xC6: {DECZP, 5, 2},
	0xD6: {DECZPX, 6, 2},
	0xCE: {DECA, 6, 3},
	0xDE: {DECAX, 7, 3},
	0xCA: {DEX, 2, 1},
	0x88: {DEY, 2, 1},
	0xE6: {INCZP, 5, 2},
	0xF6: {INCZPX, 6, 2},
	0xEE: {INCA, 6, 3},
	0xFE: {INCAX, 7, 3},
	0xE8: {INX, 2, 1},
	0xC8: {INY, 2, 1},
	0x4C: {JMPA, 3, 3},
	0x6C: {JMPI, 5, 3},
	0x20: {JSR, 6, 3},
	0x49: {EORImmediate, 2, 2},
	0x45: {EORZP, 3, 2},
	0x55: {EORZPX, 4, 2},
	0x4D: {EORA, 4, 3},
	0x5D: {EORAX, 4, 3},
	0x59: {EORAY, 4, 3},
	0x41: {EORIX, 6, 2},
	0x51: {EORIY, 5, 2},
	0x4A: {LSRAccumulator, 2, 1},
	0x46: {LSRZP, 5, 2},
	0x56: {LSRZPX, 6, 2},
	0x4E: {LSRA, 6, 3},
	0x5E: {LSRAX, 7, 3},
	0xEA: {func(*CPU, Memory) {}, 2, 1}, // NOP
	0x09: {ORAImmediate, 2, 2},
	0x05: {ORAZP, 3, 2},
	0x15: {ORAZPX, 4, 2},
	0x0D: {ORAA, 4, 3},
	0x1D: {ORAAX, 4, 3},
	0x19: {ORAAY, 4, 3},
	0x01: {ORAIX, 6, 2},
	0x11: {ORAIY, 5, 2},
	0x48: {PHA, 3, 1},
	0x08: {PHP, 3, 1},
	0x68: {PLA, 4, 1},
	0x28: {PLP, 4, 1},
	0x2A: {ROLAccumulator, 2, 1},
	0x26: {ROLZP, 5, 2},
	0x36: {ROLZPX, 6, 2},
	0x2E: {ROLA, 6, 3},
	0x3E: {ROLAX, 7, 3},
	0x6A: {RORAccumulator, 2, 1},
	0x66: {RORZP, 5, 2},
	0x76: {RORZPX, 6, 2},
	0x6E: {RORA, 6, 3},
	0x7E: {RORAX, 7, 3},
	0x40: {RTI, 6, 1},
	0x60: {RTS, 6, 1},
	0xE9: {SECImmediate, 2, 2},
	0xE5: {SECZP, 3, 2},
	0xF5: {SECZPX, 4, 2},
	0xED: {SECA, 4, 3},
	0xFD: {SECAX, 4, 3},
	0xF9: {SECAY, 4, 3},
	0xE1: {SECIX, 6, 2},
	0xF1: {SECIY, 5, 2},
}

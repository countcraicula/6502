package cpu

type instruction struct {
	I func(*CPU, Memory)
	C uint
}

func (i instruction) Execute(c *CPU, clock Clock, m Memory) bool {
	if !clock.Tick(i.C) {
		return false
	}
	i.I(c, m)
	return true
}

var instructionTable = map[byte]instruction{
	0xA9: {LDAImmediate, 2},
	0xA5: {LDAZP, 3},
	0xB5: {LDAZPX, 4},
	0xAD: {LDAA, 4},
	0xBD: {LDAAX, 4},
	0xB9: {LDAAY, 4},
	0xA1: {LDAIX, 6},
	0xB1: {LDAIY, 5},
	0xA2: {LDXImmediate, 2},
	0xA6: {LDXZP, 3},
	0xB6: {LDXZPY, 4},
	0xAE: {LDXA, 4},
	0xBE: {LDXAY, 4},
	0xA0: {LDYImmediate, 2},
	0xA4: {LDYZP, 3},
	0xB8: {LDYZPX, 4},
	0xAC: {LDYA, 4},
	0xBC: {LDYAX, 4},
	0x85: {STAZP, 3},
	0x95: {STAZPX, 4},
	0x8D: {STAA, 4},
	0x9D: {STAAX, 5},
	0x99: {STAAY, 5},
	0x81: {STAIX, 6},
	0x91: {STAIY, 6},
	0x86: {STXZP, 3},
	0x96: {STXZPY, 4},
	0x8E: {STXA, 4},
	0x84: {STYZP, 3},
	0x94: {STYZPX, 4},
	0x8C: {STYA, 4},
	0x69: {ADCImmediate, 2},
	0x65: {ADCZP, 3},
	0x75: {ADCZPX, 4},
	0x6D: {ADCA, 4},
	0x7D: {ADCAX, 4},
	0x79: {ADCAY, 4},
	0x61: {ADCIX, 6},
	0x71: {ADCIY, 5},
	0x29: {ANDImmediate, 2},
	0x25: {ANDZP, 3},
	0x35: {ANDZPX, 4},
	0x2D: {ANDA, 4},
	0x3D: {ANDAX, 4},
	0x39: {ANDAY, 4},
	0x21: {ANDIX, 6},
	0x31: {ANDIY, 5},
	0x0A: {ASLAccumulator, 2},
	0x06: {ASLZP, 5},
	0x16: {ASLZPX, 6},
	0x0E: {ASLA, 6},
	0x1E: {ASLAX, 7},
}

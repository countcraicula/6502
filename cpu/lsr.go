package cpu

func shiftRight(c *CPU, v uint8) uint8 {
	c.C = v&0x01 > 0
	a := v >> 1
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func LSR(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, shiftRight(c, m.Fetch(addr)))
}

func LSRAccumulator(c *CPU, m Memory) {
	c.A = shiftRight(c, c.A)
}

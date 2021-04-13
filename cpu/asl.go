package cpu

func ASL(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, shiftLeft(c, m.Fetch(addr)))
}

func shiftLeft(c *CPU, v uint8) uint8 {
	c.C = v&0x80 > 0
	a := v << 1
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func ASLAccumulator(c *CPU, m Memory) {
	c.A = shiftLeft(c, c.A)
}

package cpu

func AND(c *CPU, m Memory, mode MemoryMode) {
	bitwiseAnd(c, m.Fetch(mode(c, m)))
}

func bitwiseAnd(c *CPU, v uint8) {
	a := c.A & v
	c.A = uint8(a)
	c.Z = c.A == 0
	c.N = c.A&0x80 > 0
}

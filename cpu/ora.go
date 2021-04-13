package cpu

func inclusiveOR(c *CPU, v uint8) {
	a := c.A | v
	c.Z = a == 0
	c.N = a&0x80 > 0
	c.A = a
}

func ORA(c *CPU, m Memory, mode MemoryMode) {
	inclusiveOR(c, m.Fetch(mode(c, m)))
}

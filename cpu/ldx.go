package cpu

func setFlagsLDX(c *CPU) {
	c.N = c.X&0x80 > 0
	c.Z = c.X == 0
}

func LDX(c *CPU, m Memory, mode MemoryMode) {
	c.X = m.Fetch(mode(c, m))
	c.N = c.X&0x80 > 0
	c.Z = c.X == 0
}

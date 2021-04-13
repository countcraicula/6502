package cpu

func exclusiveOR(c *CPU, v uint8) {
	a := c.A ^ v
	c.Z = a == 0
	c.N = a&0x80 > 0
	c.A = a
}

func EOR(c *CPU, m Memory, mode MemoryMode) {
	exclusiveOR(c, m.Fetch(mode(c, m)))
}

package cpu

func BIT(c *CPU, m Memory, mode MemoryMode) {
	bitTest(c, m.Fetch(mode(c, m)))
}

func bitTest(c *CPU, v uint8) {
	c.Z = c.A&v == 0
	c.N = v&0x80 > 0
	c.V = v&0x40 > 0
}

package cpu

func bitTest(c *CPU, v uint8) {
	c.Z = c.A&v == 0
	c.N = v&0x80 > 0
	c.V = v&0x40 > 0
}

func BITZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	bitTest(c, m.Fetch(addr))
}

func BITA(c *CPU, m Memory) {
	addr := addrA(c, m)
	bitTest(c, m.Fetch(addr))
}

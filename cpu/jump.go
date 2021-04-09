package cpu

func JMPA(c *CPU, m Memory) {
	addr := addrA(c, m)
	c.PC = addr
}

func JMPI(c *CPU, m Memory) {
	addr := m.Fetch16(addrA(c, m))
	c.PC = addr
}

func JSR(c *CPU, m Memory) {
	pc := c.PC - 1
	push(c, m, uint8(pc&0xFF))
	push(c, m, uint8(pc>>8))
	c.PC = addrA(c, m)
}

package cpu

func STXZP(c *CPU, m Memory) {
	m.Store(addrZP(c, m), c.X)
	c.PC++
}

func STXZPY(c *CPU, m Memory) {
	m.Store(addrZPY(c, m), c.X)
	c.PC++
}

func STXA(c *CPU, m Memory) {
	m.Store(addrA(c, m), c.X)
	c.PC += 2
}

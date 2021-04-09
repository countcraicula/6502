package cpu

func STXZP(c *CPU, m Memory) {
	m.Store(addrZP(c, m), c.X)
}

func STXZPY(c *CPU, m Memory) {
	m.Store(addrZPY(c, m), c.X)
}

func STXA(c *CPU, m Memory) {
	m.Store(addrA(c, m), c.X)
}

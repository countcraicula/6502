package cpu

func STYZP(c *CPU, m Memory) {
	m.Store(addrZP(c, m), c.Y)
	c.PC++
}

func STYZPX(c *CPU, m Memory) {
	m.Store(addrZPX(c, m), c.Y)
	c.PC++
}

func STYA(c *CPU, m Memory) {
	m.Store(addrA(c, m), c.Y)
	c.PC += 2
}

package cpu

func STYZP(c *CPU, m Memory) {
	m.Store(addrZP(c, m), c.Y)
}

func STYZPX(c *CPU, m Memory) {
	m.Store(addrZPX(c, m), c.Y)
}

func STYA(c *CPU, m Memory) {
	m.Store(addrA(c, m), c.Y)
}

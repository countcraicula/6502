package cpu

func setFlagsLDY(c *CPU) {
	c.N = c.Y&0x80 > 0
	c.Z = c.Y == 0
}

func LDYImmediate(c *CPU, m Memory) {
	c.Y = m.Fetch(addrI(c, m))
	setFlagsLDY(c)
}

func LDYZP(c *CPU, m Memory) {
	c.Y = m.Fetch(addrZP(c, m))
	setFlagsLDY(c)
}

func LDYZPX(c *CPU, m Memory) {
	c.Y = m.Fetch(addrZPX(c, m))
	setFlagsLDY(c)
}

func LDYA(c *CPU, m Memory) {
	c.Y = m.Fetch(addrA(c, m))
	setFlagsLDY(c)
}

func LDYAX(c *CPU, m Memory) {
	c.Y = m.Fetch(addrAX(c, m))
	setFlagsLDY(c)
}

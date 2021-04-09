package cpu

func setFlagsLDY(c *CPU) {
	if c.Y > 127 {
		c.N = true
	}
	if c.Y == 0 {
		c.Z = true
	}
}

func LDYImmediate(c *CPU, m Memory) {
	c.Y = m.Fetch(c.PC)
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

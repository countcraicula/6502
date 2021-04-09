package cpu

func setFlagsLDA(c *CPU) {
	if c.A > 127 {
		c.N = true
	}
	if c.A == 0 {
		c.Z = true
	}
}

func LDAImmediate(c *CPU, m Memory) {
	c.A = m.Fetch(c.PC)
	setFlagsLDA(c)
}

func LDAZP(c *CPU, m Memory) {
	c.A = m.Fetch(addrZP(c, m))
	setFlagsLDA(c)
}

func LDAZPX(c *CPU, m Memory) {
	c.A = m.Fetch(addrZPX(c, m))
	setFlagsLDA(c)
}

func LDAA(c *CPU, m Memory) {
	c.A = m.Fetch(addrA(c, m))
	setFlagsLDA(c)
}

func LDAAX(c *CPU, m Memory) {
	c.A = m.Fetch(addrAX(c, m))
	setFlagsLDA(c)
}

func LDAAY(c *CPU, m Memory) {
	c.A = m.Fetch(addrAY(c, m))
	setFlagsLDA(c)
}

func LDAIX(c *CPU, m Memory) {
	c.A = m.Fetch(addrIX(c, m))
	setFlagsLDA(c)
}

func LDAIY(c *CPU, m Memory) {
	c.A = m.Fetch(addrIY(c, m))
	setFlagsLDA(c)
}

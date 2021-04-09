package cpu

func setFlagsLDA(c *CPU) {
	c.N = c.A&0x80 > 0
	c.Z = c.A == 0
}

func LDAImmediate(c *CPU, m Memory) {
	c.A = m.Fetch(addrI(c, m))
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

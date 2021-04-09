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
	c.PC++
	setFlagsLDA(c)
}

func LDAZP(c *CPU, m Memory) {
	c.A = m.Fetch(addrZP(c, m))
	c.PC++
	setFlagsLDA(c)
}

func LDAZPX(c *CPU, m Memory) {
	c.A = m.Fetch(addrZPX(c, m))
	c.PC++
	setFlagsLDA(c)
}

func LDAA(c *CPU, m Memory) {
	c.A = m.Fetch(addrA(c, m))
	c.PC += 2
	setFlagsLDA(c)
}

func LDAAX(c *CPU, m Memory) {
	c.A = m.Fetch(addrAX(c, m))
	c.PC += 2
	setFlagsLDA(c)
}

func LDAAY(c *CPU, m Memory) {
	c.A = m.Fetch(addrAY(c, m))
	c.PC += 2
	setFlagsLDA(c)
}

func LDAIX(c *CPU, m Memory) {
	c.A = m.Fetch(addrIX(c, m))
	c.PC++
	setFlagsLDA(c)
}

func LDAIY(c *CPU, m Memory) {
	c.A = m.Fetch(addrIY(c, m))
	c.PC++
	setFlagsLDA(c)
}

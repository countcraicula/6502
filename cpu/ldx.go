package cpu

func setFlagsLDX(c *CPU) {
	if c.X > 127 {
		c.N = true
	}
	if c.X == 0 {
		c.Z = true
	}
}

func LDXImmediate(c *CPU, m Memory) {
	c.X = m.Fetch(c.PC)
	setFlagsLDX(c)
}

func LDXZP(c *CPU, m Memory) {
	c.X = m.Fetch(addrZP(c, m))
	setFlagsLDX(c)
}

func LDXZPY(c *CPU, m Memory) {
	c.X = m.Fetch(addrZPY(c, m))
	setFlagsLDX(c)
}

func LDXA(c *CPU, m Memory) {
	c.X = m.Fetch(addrA(c, m))
	setFlagsLDX(c)
}

func LDXAY(c *CPU, m Memory) {
	c.X = m.Fetch(addrZPY(c, m))
	setFlagsLDX(c)
}

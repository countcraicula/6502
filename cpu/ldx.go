package cpu

func setFlagsLDX(c *CPU) {
	c.N = c.X&0x80 > 0
	c.Z = c.X == 0
}

func LDXImmediate(c *CPU, m Memory) {
	c.X = m.Fetch(addrI(c, m))
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
	c.X = m.Fetch(addrAY(c, m))
	setFlagsLDX(c)
}

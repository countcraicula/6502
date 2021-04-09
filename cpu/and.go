package cpu

func bitwiseAnd(c *CPU, v uint8) {
	a := c.A & v
	c.A = uint8(a)
	c.Z = c.A == 0
	c.N = c.A&0x80 > 0
}

func ANDImmediate(c *CPU, m Memory) {
	v := m.Fetch(addrI(c, m))
	bitwiseAnd(c, v)
}

func ANDZP(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
	bitwiseAnd(c, v)
}

func ANDZPX(c *CPU, m Memory) {
	v := m.Fetch(addrZPX(c, m))
	bitwiseAnd(c, v)
}

func ANDA(c *CPU, m Memory) {
	v := m.Fetch(addrA(c, m))
	bitwiseAnd(c, v)
}

func ANDAX(c *CPU, m Memory) {
	v := m.Fetch(addrAX(c, m))
	bitwiseAnd(c, v)
}

func ANDAY(c *CPU, m Memory) {
	v := m.Fetch(addrAY(c, m))
	bitwiseAnd(c, v)
}

func ANDIX(c *CPU, m Memory) {
	v := m.Fetch(addrIX(c, m))
	bitwiseAnd(c, v)
}

func ANDIY(c *CPU, m Memory) {
	v := m.Fetch(addrIY(c, m))
	bitwiseAnd(c, v)
}

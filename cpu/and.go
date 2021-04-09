package cpu

func bitwiseAnd(c *CPU, v uint8) {
	a := c.A & v
	c.A = uint8(a)
	if c.A == 0 {
		c.Z = true
	}
	if c.A > 127 {
		c.N = true
	}
}

func ANDImmediate(c *CPU, m Memory) {
	v := m.Fetch(c.PC)
	bitwiseAnd(c, v)
}

func ANDZP(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
	bitwiseAnd(c, v)
}

func ANDZPX(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
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

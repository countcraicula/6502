package cpu

func subWithCarry(c *CPU, v uint8) {
	a := uint16(c.A) - uint16(v)
	if c.C {
		a--
	}
	if a < 127 && v < 127 {
		if a > 127 {
			c.V = true
		}
	}
	if a > 0xFF {
		c.C = true
	}
	c.A = uint8(a)
	if c.A == 0 {
		c.Z = true
	}
	if c.A > 127 {
		c.N = true
	}
}

func SECImmediate(c *CPU, m Memory) {
	v := m.Fetch(c.PC)
	subWithCarry(c, v)
}

func SECZP(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
	subWithCarry(c, v)
}

func SECZPX(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
	subWithCarry(c, v)
}

func SECA(c *CPU, m Memory) {
	v := m.Fetch(addrA(c, m))
	subWithCarry(c, v)
}

func SECAX(c *CPU, m Memory) {
	v := m.Fetch(addrAX(c, m))
	subWithCarry(c, v)
}

func SECAY(c *CPU, m Memory) {
	v := m.Fetch(addrAY(c, m))
	subWithCarry(c, v)
}

func SECIX(c *CPU, m Memory) {
	v := m.Fetch(addrIX(c, m))
	subWithCarry(c, v)
}

func SECIY(c *CPU, m Memory) {
	v := m.Fetch(addrIY(c, m))
	subWithCarry(c, v)
}

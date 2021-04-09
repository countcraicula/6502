package cpu

func subWithCarry(c *CPU, v uint8) {
	if c.D {
		subBCD(c, v)
	} else {
		subBinary(c, v)
	}
	if c.A == 0 {
		c.Z = true
	}
	c.N = c.A&0x80 > 0
}

func subBCD(c *CPU, v uint8) {
	ad1 := c.A & 0xF
	ad2 := c.A >> 4
	vd1 := v & 0xF
	vd2 := v >> 4
	carry := uint8(0)
	rd1 := ad1 - vd1
	if rd1 > 9 {
		rd1 -= 10
		carry = 1
	}
	rd2 := ad2 - vd2 + carry
	carry = 0
	if rd2 > 9 {
		rd2 -= 10
		carry = 1
	}
	c.A = rd2<<4 | rd1
	c.C = carry > 0
}

func subBinary(c *CPU, v uint8) {
	a := uint16(c.A) + uint16(v)
	if c.C {
		a++
		c.C = false
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

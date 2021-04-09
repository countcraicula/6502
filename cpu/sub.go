package cpu

func subWithCarry(c *CPU, v uint8) {
	if c.D {
		subBCD(c, v)
	} else {
		subBinary(c, v)
	}
	c.Z = c.A == 0
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
	a := int8(c.A) - int8(v) - 1
	if c.C {
		a++
	}
	c.V = c.A > 127 && v > 127 && a >= 0
	c.C = uint8(a) > 0xFF
	c.A = uint8(a)
}

func SUBImmediate(c *CPU, m Memory) {
	v := m.Fetch(addrI(c, m))
	subWithCarry(c, v)
}

func SUBZP(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
	subWithCarry(c, v)
}

func SUBZPX(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
	subWithCarry(c, v)
}

func SUBA(c *CPU, m Memory) {
	v := m.Fetch(addrA(c, m))
	subWithCarry(c, v)
}

func SUBAX(c *CPU, m Memory) {
	v := m.Fetch(addrAX(c, m))
	subWithCarry(c, v)
}

func SUBAY(c *CPU, m Memory) {
	v := m.Fetch(addrAY(c, m))
	subWithCarry(c, v)
}

func SUBIX(c *CPU, m Memory) {
	v := m.Fetch(addrIX(c, m))
	subWithCarry(c, v)
}

func SUBIY(c *CPU, m Memory) {
	v := m.Fetch(addrIY(c, m))
	subWithCarry(c, v)
}

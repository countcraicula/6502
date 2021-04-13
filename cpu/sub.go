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
	ad := int16(ad2*10 + ad1)
	vd := int16(vd2*10 + vd1)
	rd := ad - vd
	if !c.C {
		rd--
	}
	c.C = rd >= 0
	if rd < 0 {
		rd += 100
	}
	c.A = uint8((rd/10)<<4 | (rd % 10))
}

func subBinary(c *CPU, v uint8) {
	a := uint16(c.A) + uint16(v^0xFF)
	if c.C {
		a++
	}
	c.V = (c.A^uint8(a))&(v^0xFF^uint8(a))&0x80 > 0
	c.C = a > 0xFF
	c.A = uint8(a)
}

func SUB(c *CPU, m Memory, mode MemoryMode) {
	v := m.Fetch(mode(c, m))
	subWithCarry(c, v)
}

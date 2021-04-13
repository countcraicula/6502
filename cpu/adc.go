package cpu

func ADC(c *CPU, m Memory, mode MemoryMode) {
	v := m.Fetch(mode(c, m))
	addWithCarry(c, v)
}

func addWithCarry(c *CPU, v uint8) {
	if c.D {
		addBCD(c, v)
	} else {
		addBinary(c, v)
	}
	c.Z = c.A == 0
	c.N = c.A&0x80 > 0
}

func addBCD(c *CPU, v uint8) {
	ad1 := c.A & 0xF
	ad2 := c.A >> 4
	vd1 := v & 0xF
	vd2 := v >> 4
	ad := ad2*10 + ad1
	vd := vd2*10 + vd1
	rd := ad + vd
	if c.C {
		rd++
	}
	c.A = ((rd/10)%10)<<4 | (rd % 10)
	c.C = rd > 99
}

func addBinary(c *CPU, v uint8) {
	a := uint16(c.A) + uint16(v)
	if c.C {
		a++
		c.C = false
	}
	c.V = (c.A^uint8(a))&(v^uint8(a))&0x80 > 0
	c.C = a > 0xFF
	c.A = uint8(a)
}

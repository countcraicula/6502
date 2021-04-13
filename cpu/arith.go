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

func rotateRight(c *CPU, v uint8) uint8 {
	carry := v&0x01 > 0
	a := v >> 1
	if c.C {
		a |= 0x80
	}
	c.C = carry
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func rotateLeft(c *CPU, v uint8) uint8 {
	carry := v&0x80 > 0
	a := v << 1
	if c.C {
		a |= 0x1
	}
	c.C = carry
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func inclusiveOR(c *CPU, v uint8) {
	a := c.A | v
	c.Z = a == 0
	c.N = a&0x80 > 0
	c.A = a
}

func shiftRight(c *CPU, v uint8) uint8 {
	c.C = v&0x01 > 0
	a := v >> 1
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func increment(c *CPU, v uint8) uint8 {
	v++
	c.Z = v == 0
	c.N = v&0x80 > 0
	return v
}

func exclusiveOR(c *CPU, v uint8) {
	a := c.A ^ v
	c.Z = a == 0
	c.N = a&0x80 > 0
	c.A = a
}

func decrement(c *CPU, v uint8) uint8 {
	v--
	c.Z = v == 0
	c.N = v&0x80 > 0
	return v
}

func compare(c *CPU, i, v uint8) {
	r := int8(i - v)
	c.C = r >= 0
	c.Z = r == 0
	c.N = r < 0
}

func shiftLeft(c *CPU, v uint8) uint8 {
	c.C = v&0x80 > 0
	a := v << 1
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func bitwiseAnd(c *CPU, v uint8) {
	a := c.A & v
	c.A = uint8(a)
	c.Z = c.A == 0
	c.N = c.A&0x80 > 0
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

func bitTest(c *CPU, v uint8) {
	c.Z = c.A&v == 0
	c.N = v&0x80 > 0
	c.V = v&0x40 > 0
}

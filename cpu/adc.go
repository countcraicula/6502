package cpu

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
	carry := uint8(0)
	rd1 := ad1 + vd1
	if rd1 > 9 {
		rd1 -= 10
		carry = 1
	}
	rd2 := ad2 + vd2 + carry
	carry = 0
	if rd2 > 9 {
		rd2 -= 10
		carry = 1
	}
	c.A = rd2<<4 | rd1
	c.C = carry > 0
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

func ADCImmediate(c *CPU, m Memory) {
	v := m.Fetch(addrI(c, m))
	addWithCarry(c, v)
}

func ADCZP(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
	addWithCarry(c, v)
}

func ADCZPX(c *CPU, m Memory) {
	v := m.Fetch(addrZPX(c, m))
	addWithCarry(c, v)
}

func ADCA(c *CPU, m Memory) {
	v := m.Fetch(addrA(c, m))
	addWithCarry(c, v)
}

func ADCAX(c *CPU, m Memory) {
	v := m.Fetch(addrAX(c, m))
	addWithCarry(c, v)
}

func ADCAY(c *CPU, m Memory) {
	v := m.Fetch(addrAY(c, m))
	addWithCarry(c, v)
}

func ADCIX(c *CPU, m Memory) {
	v := m.Fetch(addrIX(c, m))
	addWithCarry(c, v)
}

func ADCIY(c *CPU, m Memory) {
	v := m.Fetch(addrIY(c, m))
	addWithCarry(c, v)
}

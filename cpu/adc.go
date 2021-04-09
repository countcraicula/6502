package cpu

func addWithCarry(c *CPU, v uint8) {
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
	if c.A == 0 {
		c.Z = true
	}
	if c.A > 127 {
		c.N = true
	}
}

func ADCImmediate(c *CPU, m Memory) {
	v := m.Fetch(c.PC)
	addWithCarry(c, v)
}

func ADCZP(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
	addWithCarry(c, v)
}

func ADCZPX(c *CPU, m Memory) {
	v := m.Fetch(addrZP(c, m))
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

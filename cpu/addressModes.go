package cpu

type MemoryMode func(*CPU, Memory) uint16

func addrNone(*CPU, Memory) uint16 {
	return 0
}

func addrI(c *CPU, m Memory) uint16 {
	v := c.PC
	c.PC++
	return v
}

func addrI16(c *CPU, m Memory) uint16 {
	v := c.PC
	c.PC += 2
	return v
}

func addrZP(c *CPU, m Memory) uint16 {
	v := uint16(m.Fetch(c.PC)) + uint16(c.B)<<8
	c.PC++
	return v
}

func addrZPX(c *CPU, m Memory) uint16 {
	v := uint16(m.Fetch(c.PC)+c.X) + uint16(c.B)<<8
	c.PC++
	return v
}

func addrZPY(c *CPU, m Memory) uint16 {
	v := uint16(m.Fetch(c.PC)+c.Y) + uint16(c.B)<<8
	c.PC++
	return v
}

func addrA(c *CPU, m Memory) uint16 {
	v := m.Fetch16(c.PC)
	c.PC += 2
	return v
}

func addrAX(c *CPU, m Memory) uint16 {
	return addrA(c, m) + uint16(c.X)
}

func addrAY(c *CPU, m Memory) uint16 {
	return addrA(c, m) + uint16(c.Y)
}

func addrIZ(c *CPU, m Memory) uint16 {
	v := uint16(m.Fetch(c.PC)+c.ZR) + (uint16(c.B) << 8)
	c.PC++
	return m.Fetch16(v)
}

func addrIX(c *CPU, m Memory) uint16 {
	v := uint16(m.Fetch(c.PC)+c.X) + (uint16(c.B) << 8)
	c.PC++
	return m.Fetch16(v)
}

func addrIY(c *CPU, m Memory) uint16 {
	z := uint16(m.Fetch(c.PC)) + (uint16(c.B) << 8)
	v := m.Fetch16(z) + uint16(c.Y)
	c.PC++
	return v
}

func addrISY(c *CPU, m Memory) uint16 {
	z := uint16(m.Fetch(c.PC)+c.SP) + (uint16(c.B) << 8)
	v := m.Fetch16(z) + uint16(c.Y)
	c.PC++
	return v
}

func addrIN(c *CPU, m Memory) uint16 {
	return m.Fetch16(addrA(c, m))
}

func addrAIX(c *CPU, m Memory) uint16 {
	return m.Fetch16(addrAX(c, m))
}

func addrRel(c *CPU, m Memory) uint16 {
	v := int(c.PC) + int(int16(m.Fetch16(c.PC)))
	c.PC += 2
	return uint16(v & 0xFFFF)
}

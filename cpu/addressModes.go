package cpu

func addrZP(c *CPU, m Memory) uint16 {
	return uint16(m.Fetch(c.PC))
}

func addrZPX(c *CPU, m Memory) uint16 {
	return uint16(m.Fetch(c.PC) + c.X)
}

func addrZPY(c *CPU, m Memory) uint16 {
	return uint16(m.Fetch(c.PC) + c.Y)
}

func addrA(c *CPU, m Memory) uint16 {
	return m.Fetch16(uint16(c.PC) + uint16(c.PC+1)<<8)
}

func addrAX(c *CPU, m Memory) uint16 {
	return addrA(c, m) + uint16(c.X)
}

func addrAY(c *CPU, m Memory) uint16 {
	return addrA(c, m) + uint16(c.Y)
}

func addrIX(c *CPU, m Memory) uint16 {
	return uint16(m.Fetch(c.PC) + c.X)
}

func addrIY(c *CPU, m Memory) uint16 {
	return uint16(m.Fetch(uint16(m.Fetch(c.PC)))) + uint16(c.Y)
}

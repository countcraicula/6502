package cpu

func inclusiveOR(c *CPU, v uint8) {
	a := c.A | v
	c.Z = a == 0
	c.N = a&0x80 > 0
	c.A = a
}

func ORAImmediate(c *CPU, m Memory) {
	inclusiveOR(c, m.Fetch(c.PC))
}

func ORAZP(c *CPU, m Memory) {
	inclusiveOR(c, m.Fetch(addrZP(c, m)))
}

func ORAZPX(c *CPU, m Memory) {
	inclusiveOR(c, m.Fetch(addrZPX(c, m)))
}

func ORAA(c *CPU, m Memory) {
	inclusiveOR(c, m.Fetch(addrA(c, m)))
}

func ORAAX(c *CPU, m Memory) {
	inclusiveOR(c, m.Fetch(addrAX(c, m)))
}

func ORAAY(c *CPU, m Memory) {
	inclusiveOR(c, m.Fetch(addrAY(c, m)))
}

func ORAIX(c *CPU, m Memory) {
	inclusiveOR(c, m.Fetch(addrIX(c, m)))
}

func ORAIY(c *CPU, m Memory) {
	inclusiveOR(c, m.Fetch(addrIY(c, m)))
}

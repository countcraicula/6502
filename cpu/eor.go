package cpu

func exclusiveOR(c *CPU, v uint8) {
	a := c.A ^ v
	c.Z = a == 0
	c.N = a&0x80 > 0
	c.A = a
}

func EORImmediate(c *CPU, m Memory) {
	exclusiveOR(c, m.Fetch(addrI(c, m)))
}

func EORZP(c *CPU, m Memory) {
	exclusiveOR(c, m.Fetch(addrZP(c, m)))
}

func EORZPX(c *CPU, m Memory) {
	exclusiveOR(c, m.Fetch(addrZPX(c, m)))
}

func EORA(c *CPU, m Memory) {
	exclusiveOR(c, m.Fetch(addrA(c, m)))
}

func EORAX(c *CPU, m Memory) {
	exclusiveOR(c, m.Fetch(addrAX(c, m)))
}

func EORAY(c *CPU, m Memory) {
	exclusiveOR(c, m.Fetch(addrAY(c, m)))
}

func EORIX(c *CPU, m Memory) {
	exclusiveOR(c, m.Fetch(addrIX(c, m)))
}

func EORIY(c *CPU, m Memory) {
	exclusiveOR(c, m.Fetch(addrIY(c, m)))
}

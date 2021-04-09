package cpu

func shiftRight(c *CPU, v uint8) uint8 {
	c.C = v&0x00 > 0
	a := v >> 1
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func LSRAccumulator(c *CPU, m Memory) {
	c.A = shiftRight(c, c.A)
}

func LSRZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	m.Store(addr, shiftRight(c, m.Fetch(addr)))
}

func LSRZPX(c *CPU, m Memory) {
	addr := addrZPX(c, m)
	m.Store(addr, shiftRight(c, m.Fetch(addr)))
}

func LSRA(c *CPU, m Memory) {
	addr := addrA(c, m)
	m.Store(addr, shiftRight(c, m.Fetch(addr)))
}

func LSRAX(c *CPU, m Memory) {
	addr := addrAX(c, m)
	m.Store(addr, shiftRight(c, m.Fetch(addr)))
}

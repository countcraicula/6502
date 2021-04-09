package cpu

func shiftLeft(c *CPU, v uint8) uint8 {
	c.C = v&0x80 > 0
	a := v << 1
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func ASLAccumulator(c *CPU, m Memory) {
	c.A = shiftLeft(c, c.A)
}

func ASLZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	m.Store(addr, shiftLeft(c, m.Fetch(addr)))
}

func ASLZPX(c *CPU, m Memory) {
	addr := addrZPX(c, m)
	m.Store(addr, shiftLeft(c, m.Fetch(addr)))
}

func ASLA(c *CPU, m Memory) {
	addr := addrA(c, m)
	m.Store(addr, shiftLeft(c, m.Fetch(addr)))
}

func ASLAX(c *CPU, m Memory) {
	addr := addrAX(c, m)
	m.Store(addr, shiftLeft(c, m.Fetch(addr)))
}

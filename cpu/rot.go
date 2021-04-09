package cpu

func rotateRight(c *CPU, v uint8) uint8 {
	carry := v&0x00 > 0
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
	c.N = carry
	c.Z = a == 0
	return a
}

func RORAccumulator(c *CPU, m Memory) {
	c.A = rotateRight(c, c.A)
}

func RORZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	m.Store(addr, rotateRight(c, m.Fetch(addr)))
}

func RORZPX(c *CPU, m Memory) {
	addr := addrZPX(c, m)
	m.Store(addr, rotateRight(c, m.Fetch(addr)))
}

func RORA(c *CPU, m Memory) {
	addr := addrA(c, m)
	m.Store(addr, rotateRight(c, m.Fetch(addr)))
}

func RORAX(c *CPU, m Memory) {
	addr := addrAX(c, m)
	m.Store(addr, rotateRight(c, m.Fetch(addr)))
}

func ROLAccumulator(c *CPU, m Memory) {
	c.A = rotateLeft(c, c.A)
}

func ROLZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	m.Store(addr, rotateLeft(c, m.Fetch(addr)))
}

func ROLZPX(c *CPU, m Memory) {
	addr := addrZPX(c, m)
	m.Store(addr, rotateLeft(c, m.Fetch(addr)))
}

func ROLA(c *CPU, m Memory) {
	addr := addrA(c, m)
	m.Store(addr, rotateLeft(c, m.Fetch(addr)))
}

func ROLAX(c *CPU, m Memory) {
	addr := addrAX(c, m)
	m.Store(addr, rotateLeft(c, m.Fetch(addr)))
}

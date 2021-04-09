package cpu

func increment(c *CPU, v uint8) uint8 {
	v++
	c.Z = v == 0
	c.N = v&0x80 > 0
	return v
}

func INCZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	m.Store(addr, increment(c, m.Fetch(addr)))
}

func INCZPX(c *CPU, m Memory) {
	addr := addrZPX(c, m)
	m.Store(addr, increment(c, m.Fetch(addr)))
}

func INCA(c *CPU, m Memory) {
	addr := addrA(c, m)
	m.Store(addr, increment(c, m.Fetch(addr)))
}

func INCAX(c *CPU, m Memory) {
	addr := addrAX(c, m)
	m.Store(addr, increment(c, m.Fetch(addr)))
}

func INX(c *CPU, m Memory) {
	c.X = increment(c, c.X)
}

func INY(c *CPU, m Memory) {
	c.Y = increment(c, c.Y)
}

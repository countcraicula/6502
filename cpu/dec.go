package cpu

func decrement(c *CPU, v uint8) uint8 {
	v--
	c.Z = v == 0
	c.N = v&0x80 > 0
	return v
}

func DECZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	m.Store(addr, decrement(c, m.Fetch(addr)))
}

func DECZPX(c *CPU, m Memory) {
	addr := addrZPX(c, m)
	m.Store(addr, decrement(c, m.Fetch(addr)))
}

func DECA(c *CPU, m Memory) {
	addr := addrA(c, m)
	m.Store(addr, decrement(c, m.Fetch(addr)))
}

func DECAX(c *CPU, m Memory) {
	addr := addrAX(c, m)
	m.Store(addr, decrement(c, m.Fetch(addr)))
}

func DEX(c *CPU, m Memory) {
	c.X = decrement(c, c.X)
}

func DEY(c *CPU, m Memory) {
	c.Y = decrement(c, c.Y)
}

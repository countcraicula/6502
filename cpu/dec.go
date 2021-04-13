package cpu

func decrement(c *CPU, v uint8) uint8 {
	v--
	c.Z = v == 0
	c.N = v&0x80 > 0
	return v
}

func DEC(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, decrement(c, m.Fetch(addr)))
}

func DEX(c *CPU, m Memory) {
	c.X = decrement(c, c.X)
}

func DEY(c *CPU, m Memory) {
	c.Y = decrement(c, c.Y)
}

package cpu

func increment(c *CPU, v uint8) uint8 {
	v++
	c.Z = v == 0
	c.N = v&0x80 > 0
	return v
}

func INC(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, increment(c, m.Fetch(addr)))
}
func INX(c *CPU, m Memory) {
	c.X = increment(c, c.X)
}

func INY(c *CPU, m Memory) {
	c.Y = increment(c, c.Y)
}

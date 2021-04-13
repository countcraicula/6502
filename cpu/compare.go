package cpu

func compare(c *CPU, i, v uint8) {
	r := int8(i - v)
	c.C = r >= 0
	c.Z = r == 0
	c.N = r < 0
}

func CMP(c *CPU, m Memory, mode MemoryMode) {
	compare(c, c.A, m.Fetch(mode(c, m)))
}

func CPX(c *CPU, m Memory, mode MemoryMode) {
	compare(c, c.X, m.Fetch(mode(c, m)))
}

func CPY(c *CPU, m Memory, mode MemoryMode) {
	compare(c, c.Y, m.Fetch(mode(c, m)))
}

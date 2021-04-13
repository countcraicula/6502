package cpu

func STY(c *CPU, m Memory, mode MemoryMode) {
	m.Store(mode(c, m), c.Y)
}

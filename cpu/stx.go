package cpu

func STX(c *CPU, m Memory, mode MemoryMode) {
	m.Store(mode(c, m), c.X)
}

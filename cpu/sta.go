package cpu

func STA(c *CPU, m Memory, mode MemoryMode) {
	m.Store(mode(c, m), c.A)
}

package cpu

func setFlagsLDY(c *CPU) {
	c.N = c.Y&0x80 > 0
	c.Z = c.Y == 0
}
func LDY(c *CPU, m Memory, mode MemoryMode) {
	c.Y = m.Fetch(mode(c, m))
	c.N = c.Y&0x80 > 0
	c.Z = c.Y == 0
}

package cpu

func push(c *CPU, m Memory, v uint8) {

	m.Store(SPOffset+uint16(c.SP), v)
	c.SP++
}

func pop(c *CPU, m Memory) uint8 {
	v := m.Fetch(SPOffset + uint16(c.SP))
	c.SP--
	return v
}

func BRK(c *CPU, m Memory) {
	c.B = true
	push(c, m, uint8(c.PC>>8))
	push(c, m, uint8(c.PC&0xFF))
	push(c, m, c.GetFlags())
	c.PC = m.Fetch16(0xFFFE)
}

package cpu

func push(c *CPU, m Memory, v uint8) {

	m.Store(SPOffset+uint16(c.SP), v)
	c.SP--
}

func pull(c *CPU, m Memory) uint8 {
	c.SP++
	v := m.Fetch(SPOffset + uint16(c.SP))
	return v
}

func BRK(c *CPU, m Memory) {
	pc := c.PC + 1
	push(c, m, uint8(pc>>8))
	push(c, m, uint8(pc&0xFF))
	v := c.GetFlags()
	v |= 0x10
	push(c, m, v)
	c.I = true
	c.PC = m.Fetch16(0xFFFE)
}

func RTI(c *CPU, m Memory) {
	c.SetFlags(pull(c, m))
	c.PC = (uint16(pull(c, m)) + uint16(pull(c, m))<<8)
}

func RTS(c *CPU, m Memory) {
	c.PC = (uint16(pull(c, m)) + uint16(pull(c, m))<<8) + 1
}

func PHA(c *CPU, m Memory) {
	push(c, m, c.A)
}

func PHP(c *CPU, m Memory) {
	push(c, m, c.GetFlags())
}

func PLA(c *CPU, m Memory) {
	c.A = pull(c, m)
	c.Z = c.A == 0
	c.N = c.A&0x80 > 0
}

func PLP(c *CPU, m Memory) {
	c.SetFlags(pull(c, m))
}

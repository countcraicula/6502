package cpu

func push(c *CPU, m Memory, v uint8) {

	m.Store(SPOffset+uint16(c.SP), v)
	c.SP++
}

func pull(c *CPU, m Memory) uint8 {
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

func RTI(c *CPU, m Memory) {
	c.SetFlags(pull(c, m))
	c.PC = (uint16(pull(c, m))<<8 + uint16(pull(c, m)))
}

func RTS(c *CPU, m Memory) {
	c.PC = (uint16(pull(c, m))<<8 + uint16(pull(c, m)))
}

func PHA(c *CPU, m Memory) {
	push(c, m, c.A)
}

func PHP(c *CPU, m Memory) {
	push(c, m, c.GetFlags())
}

func PLA(c *CPU, m Memory) {
	c.A = pull(c, m)
}

func PLP(c *CPU, m Memory) {
	c.SetFlags(pull(c, m))
}

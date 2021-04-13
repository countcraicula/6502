package cpu

func push(c *CPU, m Memory, v uint8) {

	m.Store(uint16(c.SPH)<<8+uint16(c.SP), v)
	c.SP--
}

func pull(c *CPU, m Memory) uint8 {
	c.SP++
	v := m.Fetch(uint16(c.SPH)<<8 + uint16(c.SP))
	return v
}

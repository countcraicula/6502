package cpu

func CLC(c *CPU, m Memory) {
	c.C = false
}

func CLD(c *CPU, m Memory) {
	c.D = false
}

func CLI(c *CPU, m Memory) {
	c.I = false
}

func CLV(c *CPU, m Memory) {
	c.V = false
}

func SEC(c *CPU, m Memory) {
	c.C = true
}

func SED(c *CPU, m Memory) {
	c.D = true
}

func SEI(c *CPU, m Memory) {
	c.I = true
}

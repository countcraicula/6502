package cpu

func branch(c *CPU, m Memory) {
	c.PC = uint16(int32(c.PC) + int32(int8(c.PC)))
}

func BCC(c *CPU, m Memory) {
	if !c.C {
		branch(c, m)
	}
}

func BCS(c *CPU, m Memory) {
	if c.C {
		branch(c, m)
	}
}

func BEQ(c *CPU, m Memory) {
	if c.Z {
		branch(c, m)
	}
}

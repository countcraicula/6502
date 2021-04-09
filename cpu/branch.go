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

func BMI(c *CPU, m Memory) {
	if c.N {
		branch(c, m)
	}
}

func BNE(c *CPU, m Memory) {
	if !c.Z {
		branch(c, m)
	}
}

func BPL(c *CPU, m Memory) {
	if !c.N {
		branch(c, m)
	}
}

func BVC(c *CPU, m Memory) {
	if !c.V {
		branch(c, m)
	}
}

func BVS(c *CPU, m Memory) {
	if c.V {
		branch(c, m)
	}
}

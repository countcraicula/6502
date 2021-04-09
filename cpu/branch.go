package cpu

func branch(c *CPU, m Memory, flag bool) {
	v := int8(m.Fetch(addrI(c, m)))
	if flag {
		c.PC = uint16(int32(c.PC) + int32(v))
	}
}

func BCC(c *CPU, m Memory) {
	branch(c, m, !c.C)
}

func BCS(c *CPU, m Memory) {
	branch(c, m, c.C)
}

func BEQ(c *CPU, m Memory) {
	branch(c, m, c.Z)
}

func BMI(c *CPU, m Memory) {
	branch(c, m, c.N)
}

func BNE(c *CPU, m Memory) {
	branch(c, m, !c.Z)
}

func BPL(c *CPU, m Memory) {
	branch(c, m, !c.N)
}

func BVC(c *CPU, m Memory) {
	branch(c, m, !c.V)
}

func BVS(c *CPU, m Memory) {
	branch(c, m, c.V)
}

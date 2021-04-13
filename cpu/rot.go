package cpu

func rotateRight(c *CPU, v uint8) uint8 {
	carry := v&0x01 > 0
	a := v >> 1
	if c.C {
		a |= 0x80
	}
	c.C = carry
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func rotateLeft(c *CPU, v uint8) uint8 {
	carry := v&0x80 > 0
	a := v << 1
	if c.C {
		a |= 0x1
	}
	c.C = carry
	c.N = a&0x80 > 0
	c.Z = a == 0
	return a
}

func ROR(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, rotateRight(c, m.Fetch(addr)))
}

func RORAccumulator(c *CPU, m Memory) {
	c.A = rotateRight(c, c.A)
}

func ROL(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, rotateLeft(c, m.Fetch(addr)))
}

func ROLAccumulator(c *CPU, m Memory) {
	c.A = rotateLeft(c, c.A)
}

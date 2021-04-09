package cpu

func setTransferFlags(c *CPU, v uint8) {
	c.Z = v == 0
	c.N = v&0x80 > 0
}

func TAX(c *CPU, m Memory) {
	c.X = c.A
	setTransferFlags(c, c.X)
}

func TAY(c *CPU, m Memory) {
	c.Y = c.A
	setTransferFlags(c, c.Y)
}

func TSX(c *CPU, m Memory) {
	c.X = c.SP
	setTransferFlags(c, c.X)
}

func TXA(c *CPU, m Memory) {
	c.A = c.X
	setTransferFlags(c, c.A)
}

func TXS(c *CPU, m Memory) {
	c.SP = c.X
}

func TYA(c *CPU, m Memory) {
	c.A = c.Y
	setTransferFlags(c, c.A)
}

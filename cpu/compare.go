package cpu

func compare(c *CPU, i, v uint8) {
	r := int8(i - v)
	if r >= 0 {
		c.C = true

	}
	if r == 0 {
		c.Z = true
	}

	c.N = r < 0
}

func CMPImmediate(c *CPU, m Memory) {
	compare(c, c.A, m.Fetch(c.PC))
}

func CMPZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	compare(c, c.A, m.Fetch(addr))
}

func CMPZPX(c *CPU, m Memory) {
	addr := addrZPX(c, m)
	compare(c, c.A, m.Fetch(addr))
}

func CMPA(c *CPU, m Memory) {
	addr := addrA(c, m)
	compare(c, c.A, m.Fetch(addr))
}

func CMPAX(c *CPU, m Memory) {
	addr := addrAX(c, m)
	compare(c, c.A, m.Fetch(addr))
}

func CMPAY(c *CPU, m Memory) {
	addr := addrAY(c, m)
	compare(c, c.A, m.Fetch(addr))
}

func CMPIX(c *CPU, m Memory) {
	addr := addrIX(c, m)
	compare(c, c.A, m.Fetch(addr))
}

func CMPIY(c *CPU, m Memory) {
	addr := addrIY(c, m)
	compare(c, c.A, m.Fetch(addr))
}

func CPXImmediate(c *CPU, m Memory) {
	compare(c, c.X, m.Fetch(c.PC))
}

func CPXZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	compare(c, c.X, m.Fetch(addr))
}

func CPXA(c *CPU, m Memory) {
	addr := addrA(c, m)
	compare(c, c.X, m.Fetch(addr))
}

func CPYImmediate(c *CPU, m Memory) {
	compare(c, c.Y, m.Fetch(c.PC))
}

func CPYZP(c *CPU, m Memory) {
	addr := addrZP(c, m)
	compare(c, c.Y, m.Fetch(addr))
}

func CPYA(c *CPU, m Memory) {
	addr := addrA(c, m)
	compare(c, c.Y, m.Fetch(addr))
}

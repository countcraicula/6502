package cpu

func STAZP(c *CPU, m Memory) {
	m.Store(addrZP(c, m), c.A)
}

func STAZPX(c *CPU, m Memory) {
	m.Store(addrZPX(c, m), c.A)
}

func STAA(c *CPU, m Memory) {
	m.Store(addrA(c, m), c.A)
}

func STAAX(c *CPU, m Memory) {
	m.Store(addrAX(c, m), c.A)
}

func STAAY(c *CPU, m Memory) {
	m.Store(addrAY(c, m), c.A)
}

func STAIX(c *CPU, m Memory) {
	m.Store(addrIX(c, m), c.A)
}

func STAIY(c *CPU, m Memory) {
	m.Store(addrIY(c, m), c.A)
	setFlagsLDA(c)
}

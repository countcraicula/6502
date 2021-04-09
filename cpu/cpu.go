package cpu

type Clock uint

func (c Clock) Tick(count uint) bool {
	i := Clock(count)
	if c > i {
		c -= i
		return true
	}
	c = 0
	return false
}

func (c Clock) Add(i uint) {
	c += Clock(i)
}

type CPU struct {
	PC                  uint16
	SP                  uint8
	A, X, Y             uint8
	C, Z, I, D, B, V, N bool
}

const SPOffset = 0x0100

func (c *CPU) Reset() {
	c.PC = 0xFFFC
	c.SP = 0x00
	c.D = false
}

func (c *CPU) Execute(clock Clock, m Memory) {

	for clock > 0 {
		b := m.Fetch(c.PC)
		if !instructionTable[b].Execute(c, clock, m) {
			return
		}
		c.PC++
	}
}

func (c *CPU) GetFlags() uint8 {
	var v uint8
	if c.C {
		v |= 0x1
	}
	if c.Z {
		v |= 0x2
	}
	if c.I {
		v |= 0x4
	}
	if c.D {
		v |= 0x8
	}
	if c.B {
		v |= 0x10
	}
	if c.V {
		v |= 0x20
	}
	if c.N {
		v |= 0x40
	}
	return v
}

func (c *CPU) SetFlags(v uint8) {
	c.C = v&0x1 > 0
	c.Z = v&0x2 > 0
	c.I = v&0x4 > 0
	c.D = v&0x8 > 0
	c.B = v&0x10 > 0
	c.V = v&0x20 > 0
	c.N = v&0x40 > 0

}

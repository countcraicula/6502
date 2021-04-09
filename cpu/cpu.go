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

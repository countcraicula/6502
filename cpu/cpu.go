package cpu

import (
	"fmt"
)

type Clock struct {
	count      uint
	num        int
	singleStep bool
	stepped    bool
}

func NewClock(i int, singleStep bool) *Clock {
	return &Clock{
		num:        i,
		singleStep: singleStep,
	}
}

func (c *Clock) Tick(count int) bool {
	if c.singleStep && !c.stepped {
		c.stepped = true
		c.count += uint(count)
		return true
	}
	if c.stepped {
		return false
	}

	c.count += uint(count)
	if c.num < 0 {
		return true
	}
	if c.num > count {
		c.num -= count
		return true
	}
	c.num = 0
	return false
}

func (c *Clock) Step() {
	c.stepped = false
}

func (c *Clock) Add(i int) {
	c.num += i
}

type CPU struct {
	PC               uint16
	SP               uint8
	IR               uint8
	A, B, X, Y, ZR   uint8
	C, Z, I, D, V, N bool
	Halt, Wait       bool
	IRQ              chan bool
}

const SPOffset = 0x0100

func (c *CPU) Reset(m Memory) {
	//	c.PC = m.Fetch16(0xFFFC)
	c.PC = 0x0400
	c.SP = 0xFD
	c.D = false
	c.B = 0
	c.ZR = 0
}

func (c *CPU) String() string {
	return fmt.Sprintf("PC: 0x%x, SP: 0x%x, IR:0x%x, A: 0x%x, X: 0x%x, Y: 0x%x, C: %v, Z: %v, I: %v, D: %v,  V: %v, N: %v\n", c.PC, c.SP, c.IR, c.A, c.X, c.Y, c.C, c.Z, c.I, c.D, c.V, c.N)
}

func (c *CPU) Execute(clock *Clock, m Memory) {
	for {
		if c.Halt {
			return
		}
		if c.Wait {
			<-c.IRQ
		}
		pc := c.PC
		c.IR = m.Fetch(c.PC)
		ins := fastLookup[c.IR]
		if ins == nil {
			fmt.Printf("Unknown instruction 0x%x\n", c.IR)
			return
		}
		if !ins.Execute(c, clock, m) {
			return
		}
		log(c, clock)
		if c.PC == pc {
			fmt.Printf("Caught in a loop\n")
			fmt.Println(defaultLogger.String())
			return
		}
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
	v |= 0x10 // B only changed by interrupts.
	v |= 0x20 // Unregistered bit
	if c.V {
		v |= 0x40
	}
	if c.N {
		v |= 0x80
	}
	return v
}

func (c *CPU) SetFlags(v uint8) {
	c.C = v&0x1 > 0
	c.Z = v&0x2 > 0
	c.I = v&0x4 > 0
	c.D = v&0x8 > 0
	c.V = v&0x40 > 0
	c.N = v&0x80 > 0

}

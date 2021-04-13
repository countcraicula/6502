package cpu

import (
	"bytes"
	"fmt"
)

var defaultLogger *logger

var EnableLogging bool = true

func log(c *CPU, cl *Clock) {
	if defaultLogger != nil {
		defaultLogger.Log(c, cl)
	}
}

type logger struct {
	buffer []*state
	len    int
	head   int
}

type state struct {
	cpu   *CPU
	clock uint
}

func (s *state) String() string {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, instructionTable[s.cpu.IR])
	fmt.Fprintln(&buf, s.cpu)
	fmt.Fprintf(&buf, "Clock cycle: %v", s.clock)
	return buf.String()
}

func (l *logger) next() *state {
	var ret *state
	l.head++
	if l.head >= l.len {
		l.head = l.head % l.len
	}
	if l.buffer[l.head] == nil {
		ret = &state{
			cpu: &CPU{},
		}
		l.buffer[l.head] = ret
	} else {
		ret = l.buffer[l.head]
	}
	return ret
}

func (l *logger) Log(c *CPU, cl *Clock) {
	s := l.next()
	s.clock = cl.count
	nc := s.cpu
	nc.A = c.A
	nc.X = c.X
	nc.Y = c.Y
	nc.IR = c.IR
	nc.PC = c.PC
	nc.SP = c.SP
	nc.C = c.C
	nc.D = c.D
	nc.I = c.I
	nc.N = c.N
	nc.V = c.V
	nc.Z = c.Z
}

func (l *logger) String() string {
	if l == nil {
		return ""
	}
	var buf bytes.Buffer
	i := (l.head + 1) % l.len
	for i != l.head {
		i = (i + 1) % l.len
		fmt.Fprintln(&buf, l.buffer[i])
	}
	return buf.String()
}

package cpu

import "fmt"

type Memory []byte

func (m Memory) Fetch(addr uint16) byte {
	v := m[addr]
	fmt.Printf("Memory access at 0x%x(%v) 8bits -> 0x%x(%v)\n", addr, addr, v, v)
	return v
}

func (m Memory) Fetch16(addr uint16) uint16 {
	v := uint16(m[addr]) + (uint16(m[addr+1]) << 8)
	fmt.Printf("Memory access at 0x%x(%v) 16bits -> 0x%x(%v)\n", addr, addr, v, v)
	return v
}

func (m Memory) Store(addr uint16, value byte) {
	fmt.Printf("Value 0x%x(%v) stored at location 0x%x(%v)\n", value, value, addr, addr)
	m[addr] = value
}

package cpu

type Memory []byte

func (m Memory) Fetch(addr uint16) byte {
	v := m[addr]
	return v
}

func (m Memory) Fetch16(addr uint16) uint16 {
	v := uint16(m[addr]) + (uint16(m[addr+1]) << 8)
	return v
}

func (m Memory) Store(addr uint16, value byte) {
	m[addr] = value
}

package cpu

type Memory []byte

func (m Memory) Fetch(addr uint16) byte {
	return m[int(addr)]
}

func (m Memory) Fetch16(addr uint16) uint16 {
	return uint16(m[int(addr>>8)]) + (uint16(m[int(addr&0xFF)]) << 8)
}

func (m Memory) Store(addr uint16, value byte) {
	m[int(addr)] = value
}

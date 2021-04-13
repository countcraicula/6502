package cpu

func branch(c *CPU, m Memory, flag bool) {
	v := int8(m.Fetch(addrI(c, m)))
	if flag {
		c.PC = uint16(int32(c.PC) + int32(v))
	}
}

func setTransferFlags(c *CPU, v uint8) {
	c.Z = v == 0
	c.N = v&0x80 > 0
}

func BBR(c *CPU, m Memory, _ MemoryMode) {
	bit := uint8(0x1 << (c.IR >> 4))
	v := m.Fetch(addrI(c, m))
	branch(c, m, !(v&bit > 0))
}

func BBS(c *CPU, m Memory, _ MemoryMode) {
	bit := uint8(0x1 << ((c.IR >> 4) - 8))
	v := m.Fetch(addrI(c, m))
	branch(c, m, (v&bit > 0))
}

func RMB(c *CPU, m Memory, mode MemoryMode) {
	mask := uint8(0x1<<(c.IR>>4)) ^ 0xFF
	addr := mode(c, m)
	m.Store(addr, m.Fetch(addr)&mask)
}

func SMB(c *CPU, m Memory, mode MemoryMode) {
	mask := uint8(0x1 << ((c.IR >> 4) - 8))
	addr := mode(c, m)
	m.Store(addr, m.Fetch(addr)|mask)
}

func ADC(c *CPU, m Memory, mode MemoryMode) {
	v := m.Fetch(mode(c, m))
	addWithCarry(c, v)
}

func AND(c *CPU, m Memory, mode MemoryMode) {
	bitwiseAnd(c, m.Fetch(mode(c, m)))
}

func ASL(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, shiftLeft(c, m.Fetch(addr)))
}

func ASLAccumulator(c *CPU, m Memory, _ MemoryMode) {
	c.A = shiftLeft(c, c.A)
}

func BIT(c *CPU, m Memory, mode MemoryMode) {
	bitTest(c, m.Fetch(mode(c, m)))
}

func BCC(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, !c.C)
}

func BCS(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, c.C)
}

func BEQ(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, c.Z)
}

func BMI(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, c.N)
}

func BNE(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, !c.Z)
}

func BPL(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, !c.N)
}

func BRA(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, true)
}

func BRK(c *CPU, m Memory, _ MemoryMode) {
	pc := c.PC + 1
	push(c, m, uint8(pc>>8))
	push(c, m, uint8(pc&0xFF))
	v := c.GetFlags()
	v |= 0x10
	push(c, m, v)
	c.I = true
	c.PC = m.Fetch16(0xFFFE)
}

func BVC(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, !c.V)
}

func BVS(c *CPU, m Memory, _ MemoryMode) {
	branch(c, m, c.V)
}

func CMP(c *CPU, m Memory, mode MemoryMode) {
	compare(c, c.A, m.Fetch(mode(c, m)))
}

func CPX(c *CPU, m Memory, mode MemoryMode) {
	compare(c, c.X, m.Fetch(mode(c, m)))
}

func CPY(c *CPU, m Memory, mode MemoryMode) {
	compare(c, c.Y, m.Fetch(mode(c, m)))
}

func CLC(c *CPU, m Memory, _ MemoryMode) {
	c.C = false
}

func CLD(c *CPU, m Memory, _ MemoryMode) {
	c.D = false
}

func CLI(c *CPU, m Memory, _ MemoryMode) {
	c.I = false
}

func CLV(c *CPU, m Memory, _ MemoryMode) {
	c.V = false
}

func DEC(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, decrement(c, m.Fetch(addr)))
}

func DEX(c *CPU, m Memory, _ MemoryMode) {
	c.X = decrement(c, c.X)
}

func DEY(c *CPU, m Memory, _ MemoryMode) {
	c.Y = decrement(c, c.Y)
}

func EOR(c *CPU, m Memory, mode MemoryMode) {
	exclusiveOR(c, m.Fetch(mode(c, m)))
}

func INC(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, increment(c, m.Fetch(addr)))
}

func INX(c *CPU, m Memory, _ MemoryMode) {
	c.X = increment(c, c.X)
}

func INY(c *CPU, m Memory, _ MemoryMode) {
	c.Y = increment(c, c.Y)
}

func JMP(c *CPU, m Memory, mode MemoryMode) {
	c.PC = mode(c, m)
}

func JSR(c *CPU, m Memory, _ MemoryMode) {
	pc := c.PC + 1
	push(c, m, uint8(pc>>8))
	push(c, m, uint8(pc&0xFF))
	c.PC = addrA(c, m)
}

func LDA(c *CPU, m Memory, mode MemoryMode) {
	c.A = m.Fetch(mode(c, m))
	c.N = c.A&0x80 > 0
	c.Z = c.A == 0
}

func LDX(c *CPU, m Memory, mode MemoryMode) {
	c.X = m.Fetch(mode(c, m))
	c.N = c.X&0x80 > 0
	c.Z = c.X == 0
}

func LDY(c *CPU, m Memory, mode MemoryMode) {
	c.Y = m.Fetch(mode(c, m))
	c.N = c.Y&0x80 > 0
	c.Z = c.Y == 0
}

func LSR(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, shiftRight(c, m.Fetch(addr)))
}

func LSRAccumulator(c *CPU, m Memory, _ MemoryMode) {
	c.A = shiftRight(c, c.A)
}

func ORA(c *CPU, m Memory, mode MemoryMode) {
	inclusiveOR(c, m.Fetch(mode(c, m)))
}

func ROR(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, rotateRight(c, m.Fetch(addr)))
}

func RORAccumulator(c *CPU, m Memory, _ MemoryMode) {
	c.A = rotateRight(c, c.A)
}

func ROL(c *CPU, m Memory, mode MemoryMode) {
	addr := mode(c, m)
	m.Store(addr, rotateLeft(c, m.Fetch(addr)))
}

func ROLAccumulator(c *CPU, m Memory, _ MemoryMode) {
	c.A = rotateLeft(c, c.A)
}

func RTI(c *CPU, m Memory, _ MemoryMode) {
	c.SetFlags(pull(c, m))
	c.PC = (uint16(pull(c, m)) + uint16(pull(c, m))<<8)
}

func RTS(c *CPU, m Memory, _ MemoryMode) {
	c.PC = (uint16(pull(c, m)) + uint16(pull(c, m))<<8) + 1
}

func PHA(c *CPU, m Memory, _ MemoryMode) {
	push(c, m, c.A)
}

func PHX(c *CPU, m Memory, _ MemoryMode) {
	push(c, m, c.X)
}

func PHY(c *CPU, m Memory, _ MemoryMode) {
	push(c, m, c.Y)
}

func PHP(c *CPU, m Memory, _ MemoryMode) {
	push(c, m, c.GetFlags())
}

func pullRegister(c *CPU, m Memory) uint8 {
	v := pull(c, m)
	c.Z = v == 0
	c.N = v&0x80 > 0
	return v
}

func PLA(c *CPU, m Memory, _ MemoryMode) {
	c.A = pullRegister(c, m)
}

func PLX(c *CPU, m Memory, _ MemoryMode) {
	c.X = pullRegister(c, m)
}

func PLY(c *CPU, m Memory, _ MemoryMode) {
	c.Y = pullRegister(c, m)
}

func PLP(c *CPU, m Memory, _ MemoryMode) {
	c.SetFlags(pull(c, m))
}

func SEC(c *CPU, m Memory, _ MemoryMode) {
	c.C = true
}

func SED(c *CPU, m Memory, _ MemoryMode) {
	c.D = true
}

func SEI(c *CPU, m Memory, _ MemoryMode) {
	c.I = true
}

func STA(c *CPU, m Memory, mode MemoryMode) {
	m.Store(mode(c, m), c.A)
}

func STX(c *CPU, m Memory, mode MemoryMode) {
	m.Store(mode(c, m), c.X)
}

func STY(c *CPU, m Memory, mode MemoryMode) {
	m.Store(mode(c, m), c.Y)
}

func STZ(c *CPU, m Memory, mode MemoryMode) {
	m.Store(mode(c, m), 0)
}

func STP(c *CPU, _ Memory, _ MemoryMode) {
	c.Halt = true
}

func WAI(c *CPU, _ Memory, _ MemoryMode) {
	c.Wait = true
}

func SUB(c *CPU, m Memory, mode MemoryMode) {
	v := m.Fetch(mode(c, m))
	subWithCarry(c, v)
}

func TAX(c *CPU, m Memory, _ MemoryMode) {
	c.X = c.A
	setTransferFlags(c, c.X)
}

func TAY(c *CPU, m Memory, _ MemoryMode) {
	c.Y = c.A
	setTransferFlags(c, c.Y)
}

func TSX(c *CPU, m Memory, _ MemoryMode) {
	c.X = c.SP
	setTransferFlags(c, c.X)
}

func TXA(c *CPU, m Memory, _ MemoryMode) {
	c.A = c.X
	setTransferFlags(c, c.A)
}

func TXS(c *CPU, m Memory, _ MemoryMode) {
	c.SP = c.X
}

func TYA(c *CPU, m Memory, _ MemoryMode) {
	c.A = c.Y
	setTransferFlags(c, c.A)
}

func NOP(*CPU, Memory, MemoryMode) {}

package decoder

// Byte returns the value at the internal pointer and increments it accordingly
func (p *Packet) Byte() byte {
	if p.idx >= p.length {
		p.Err = ErrReadPastEndData
		return 0
	}
	v := p.buf[p.idx]
	p.idx += 1
	return v
}

// Uint16 returns the value at the internal pointer and increments it accordingly
func (p *Packet) Uint16() uint16 {
	if p.idx+1 >= p.length {
		p.Err = ErrReadPastEndData
		return 0
	}
	v := p.endian.Uint16(p.buf[p.idx:])
	p.idx += 2
	return v
}

// Uint32 returns the value at the internal pointer and increments it accordingly
func (p *Packet) Uint32() uint32 {
	if p.idx+3 >= p.length {
		p.Err = ErrReadPastEndData
		return 0
	}
	v := p.endian.Uint32(p.buf[p.idx:])
	p.idx += 4
	return v
}

// Uint64 returns the value at the internal pointer and increments it accordingly
func (p *Packet) Uint64() uint64 {
	if p.idx+7 >= p.length {
		p.Err = ErrReadPastEndData
		return 0
	}
	v := p.endian.Uint64(p.buf[p.idx:])
	p.idx += 8
	return v
}

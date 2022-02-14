package decoder

import (
	"encoding/binary"
)

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

// Uint24 returns the 24 bit (3 bytes) value as a Uint32 at the internal pointer and increments it accordingly
func (p *Packet) Uint24() uint32 {
	if p.idx+2 >= p.length {
		p.Err = ErrReadPastEndData
		return 0
	}
	// copy 3 bytes into the middle of a 5 byte slice making it easy to read it as little or big endian
	b := make([]byte, 5)
	copy(b[1:], p.buf[p.idx:p.idx+3])
	p.idx += 4
	if p.endian == binary.BigEndian {
		return p.endian.Uint32(b[0:])
	}
	return p.endian.Uint32(b[1:])
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

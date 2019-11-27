/*
	Decoder takes a []byte and provides safe & simple functions to read different
	types, and keeps track of the last read position.
*/
package decoder

import (
	"encoding/binary"
	"errors"
)

type Packet struct {
	buf    []byte           // The raw data with dummy 0's appended
	lenght int              // The actual data length
	idx    int              // The current idx we've read up to
	err    error            // The last error
	endian binary.ByteOrder // The endian to use for decoding
}

var ErrReadPastEndData = errors.New("read past of end of data")

// New returns a loaded packet ready for reading
func New(b []byte) *Packet {
	return &Packet{
		// Append to buf so we dont have to keep checking for index bounds
		buf:    append(b, make([]byte, 128)...),
		lenght: len(b),
		idx:    0,
		err:    nil,
		endian: binary.BigEndian,
	}
}

// SetLittleEndian set future read to be in little endian
func (p *Packet) SetLittleEndian() {
	p.endian = binary.LittleEndian
}

// SeBigEndian set future read to be in big endian
func (p *Packet) SetBigEndian() {
	p.endian = binary.BigEndian
}

// Reset moves the internal read point back to the start
func (p *Packet) Reset() {
	p.idx = 0
}

// Rewind moves the internal pointer backwards (or forward if passed a negative value)
func (p *Packet) Rewind(i int) error {
	idx := p.idx - i
	if idx < 0 || idx >= p.lenght {
		return ErrReadPastEndData
	}
	p.idx = idx
	return nil
}

// Byte returns the value at the internal pointer and increments it accordingly
func (p *Packet) Byte() (byte, error) {
	if p.idx >= p.lenght {
		p.err = ErrReadPastEndData
		return 0, p.err
	}
	v := p.buf[p.idx]
	p.idx += 1
	return v, nil
}

// Uint16 returns the value at the internal pointer and increments it accordingly
func (p *Packet) Uint16() (uint16, error) {
	if p.idx+1 >= p.lenght {
		p.err = ErrReadPastEndData
		return 0, p.err
	}
	v := p.endian.Uint16(p.buf[p.idx:])
	p.idx += 2
	return v, nil
}

// Uint32 returns the value at the internal pointer and increments it accordingly
func (p *Packet) Uint32() (uint32, error) {
	if p.idx+3 >= p.lenght {
		p.err = ErrReadPastEndData
		return 0, p.err
	}
	v := p.endian.Uint32(p.buf[p.idx:])
	p.idx += 4
	return v, nil
}

// Bits8 returns an array of bool where [0] is the right hand bit at the internal pointer and increments it accordingly
func (p *Packet) Bits8() ([]bool, error) {
	b, err := p.Byte()
	if err != nil {
		return nil, err
	}
	v := make([]bool, 8)
	for i := range v {
		v[i] = b&0x01 != 0
		b = b >> 1
	}
	return v, nil
}

// StringPrefixByteLen returns the string at internal pointer using the first byte as it's lenght and increments it accordingly
func (p *Packet) StringPrefixByteLen() (string, error) {
	if p.idx >= p.lenght {
		p.err = ErrReadPastEndData
		return "", p.err
	}
	l := int(p.buf[p.idx])
	p.idx++
	if p.idx+l >= p.lenght {
		p.err = ErrReadPastEndData
		return "", p.err
	}
	v := string(p.buf[p.idx : p.idx+l])
	p.idx += l
	return v, nil
}

// StringPrefixUint16Len returns the string at internal pointer using the first 2 bytes as it's lenght and increments it accordingly
func (p *Packet) StringPrefixUint16Len() (string, error) {
	x, err := p.Uint16()
	l := int(x)

	if err != nil {
		return "", err
	}
	if p.idx+l >= p.lenght {
		p.err = ErrReadPastEndData
		return "", p.err
	}
	v := string(p.buf[p.idx : p.idx+l])
	p.idx += l
	return v, nil
}

// CString returns the string at internal pointer (terminated with a 0x00) and increments it accordingly
func (p *Packet) CString() (string, error) {
	idxStart := p.idx
	for idx := p.idx; idx < p.lenght; idx++ {
		if p.buf[idx] == 0x00 {
			p.idx = idx + 1
			return string(p.buf[idxStart:idx]), nil
		}
	}
	p.err = ErrReadPastEndData
	return "", p.err
}

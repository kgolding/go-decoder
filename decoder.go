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
	Err    error            // The last error
	endian binary.ByteOrder // The endian to use for decoding
}

var ErrReadPastEndData = errors.New("read past of end of data")

// New returns a loaded packet ready for reading
func New(b []byte) *Packet {
	return &Packet{
		buf:    b,
		lenght: len(b),
		idx:    0,
		Err:    nil,
		endian: binary.BigEndian,
	}
}

// Index returns the current pointer position
func (p *Packet) Index() int {
	return p.idx
}

// PeekBytes returns the []byte data
func (p *Packet) PeekBytes() []byte {
	return p.buf
}

// PeekRemainingBytes returns the bytes from the current pointer postion
func (p *Packet) PeekRemainingBytes() []byte {
	return p.buf[p.idx:]
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
	p.Err = nil
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

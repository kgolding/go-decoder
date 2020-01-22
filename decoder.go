/*
	Decoder takes a []byte and provides safe & simple functions to read different
	types, and keeps track of the last read position.
*/
package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Packet struct {
	buf    []byte           // The raw data with dummy 0's appended
	length int              // The actual data length
	idx    int              // The current idx we've read up to
	Err    error            // The last error
	endian binary.ByteOrder // The endian to use for decoding
}

var ErrReadPastEndData = errors.New("read past of end of data")
var ErrReadInvalidLength = errors.New("invalid length")

// New returns a loaded packet ready for reading
func New(b []byte) *Packet {
	return &Packet{
		buf:    b,
		length: len(b),
		idx:    0,
		Err:    nil,
		endian: binary.BigEndian,
	}
}

// Seek Byteto the next instance of the given byte
func (p *Packet) SeekByte(b byte) bool {
	return p.Seek([]byte{b})
}

// Seek to the next instance of the given bytes
func (p *Packet) Seek(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	for i := p.idx; i < len(p.buf)+len(b)-1; i++ {
		if p.buf[i] == b[0] && bytes.Compare(p.buf[i:i+len(b)], b) == 0 {
			p.idx = i
			return true
		}
	}
	return false
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

// RemainingLength returns the number of bytes from the current pointer postion
func (p *Packet) RemainingLength() int {
	return len(p.buf) - p.idx
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
	if idx < 0 || idx >= p.length {
		return ErrReadPastEndData
	}
	p.idx = idx
	return nil
}

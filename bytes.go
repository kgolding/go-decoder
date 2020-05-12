package decoder

// Bytes returns the given number of bytes at the internal pointer and increments it accordingly
func (p *Packet) Bytes(length int) []byte {
	if length < 1 {
		return []byte{} // Ask for nothing... you get nothing
	}
	newIdx := p.idx + length
	if p.idx >= p.length || newIdx > p.length {
		p.Err = ErrReadPastEndData
		return nil
	}
	if newIdx > len(p.buf) { // Just in case p.length is bad
		newIdx = len(p.buf) - 1
	}
	b := p.buf[p.idx:newIdx]
	p.idx = newIdx
	return b
}

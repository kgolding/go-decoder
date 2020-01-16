package decoder

// Bytes returns the given number of bytes at the internal pointer and increments it accordingly
func (p *Packet) Bytes(length int) []byte {
	newIdx := p.idx + length
	if p.idx >= p.length || newIdx >= p.length {
		p.Err = ErrReadPastEndData
		return nil
	}
	b := p.buf[p.idx:newIdx]
	p.idx = newIdx
	return b
}

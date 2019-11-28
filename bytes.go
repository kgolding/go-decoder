package decoder

// Vytes returns the given number of bytes at the internal pointer and increments it accordingly
func (p *Packet) Bytes(length int) []byte {
	if p.idx+length-1 >= p.lenght {
		p.Err = ErrReadPastEndData
		return nil
	}
	b := p.buf[p.idx : p.idx+length]
	p.idx += length
	return b
}

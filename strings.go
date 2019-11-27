package decoder

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
func (p *Packet) StringPrefixUint16Len() string {
	l := int(p.Uint16())

	if p.err != nil {
		return ""
	}
	if p.idx+l >= p.lenght {
		p.err = ErrReadPastEndData
		return ""
	}
	v := string(p.buf[p.idx : p.idx+l])
	p.idx += l
	return v
}

// CString returns the string at internal pointer (terminated with a 0x00) and increments it accordingly
func (p *Packet) CString() string {
	idxStart := p.idx
	for idx := p.idx; idx < p.lenght; idx++ {
		if p.buf[idx] == 0x00 {
			p.idx = idx + 1
			return string(p.buf[idxStart:idx])
		}
	}
	p.err = ErrReadPastEndData
	return ""
}

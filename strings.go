package decoder

import (
	"bytes"
)

// StringByDelimiter returns the string at internal pointer using the given delimiter at the end marker
func (p *Packet) StringByDelimiter(delimiter byte) string {
	idx := p.idx

	index := bytes.IndexByte(p.buf[idx:], delimiter)
	if index == -1 {
		return ""
	}

	if p.idx+index >= p.length {
		p.Err = ErrReadPastEndData
		return ""
	}

	p.idx += index + 1
	return string(p.buf[idx : idx+index])
}

// StringPrefixByteLen returns the string at internal pointer using the first byte as it's lenght and increments it accordingly
func (p *Packet) StringPrefixByteLen() string {
	if p.idx >= p.length {
		p.Err = ErrReadPastEndData
		return ""
	}
	l := int(p.buf[p.idx])
	p.idx++
	if p.idx+l > p.length {
		p.Err = ErrReadPastEndData
		return ""
	}
	v := string(p.buf[p.idx : p.idx+l])
	p.idx += l
	return v
}

// StringZeroPadded returns the null padded string at internal pointer
func (p *Packet) StringZeroPadded(fixedLength int) string {
	if p.idx+fixedLength > p.length {
		p.Err = ErrReadPastEndData
		return ""
	}
	idx := p.idx
	p.idx += fixedLength

	b := p.buf[idx:p.idx]

	nullIndex := bytes.IndexByte(b, 0x00)
	if nullIndex == -1 {
		return string(b)
	}
	return string(p.buf[idx : idx+nullIndex])
}

// StringPrefixUint16Len returns the string at internal pointer using the first 2 bytes as it's lenght and increments it accordingly
func (p *Packet) StringPrefixUint16Len() string {
	l := int(p.Uint16())

	if p.Err != nil {
		return ""
	}
	if p.idx+l > p.length {
		p.Err = ErrReadPastEndData
		return ""
	}
	v := string(p.buf[p.idx : p.idx+l])
	p.idx += l
	return v
}

// CString returns the string at internal pointer (terminated with a 0x00) and increments it accordingly
func (p *Packet) CString() string {
	idxStart := p.idx
	for idx := p.idx; idx < p.length; idx++ {
		if p.buf[idx] == 0x00 {
			p.idx = idx + 1
			return string(p.buf[idxStart:idx])
		}
	}
	p.Err = ErrReadPastEndData
	return ""
}

// StringByWhitelist returns the string at internal pointer using the given whitelist bytes
func (p *Packet) StringByWhitelist(whitelist []byte) string {
	idxStart := p.idx
	var idx int
	for idx = p.idx; idx < p.length; idx++ {
		if bytes.IndexByte(whitelist, p.buf[idx]) == -1 {
			p.idx = idx
			break
		}
	}
	return string(p.buf[idxStart:idx])
}

// StringHex returns the string at internal pointer that has HEX [0-9a-xA-Z] chars
func (p *Packet) StringHex() string {
	return p.StringByWhitelist([]byte("0123456789abcdefABCDEF"))
}

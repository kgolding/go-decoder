package decoder

// AsciiInt returns the ASCII Integer at the internal pointer
func (p *Packet) AsciiInt() (v int) {
	negate := false
	nodata := true

	defer func() {
		if negate {
			v = -v
		}
		if nodata {
			p.Err = ErrReadNoData
		}
	}()

	var idx int
	for idx = p.idx; idx < p.length; idx++ {
		c := p.buf[idx]
		if idx == p.idx && c == '-' { // Leading "-""
			negate = true
		} else {
			if c >= '0' && c <= '9' {
				nodata = false
				v = v * 10
				v += int(c - '0')
			} else {
				break
			}
		}
	}
	p.idx = idx
	return
}

// AsciiUInt returns the ASCII Integer at the internal pointer
func (p *Packet) AsciiUInt() (v uint) {
	nodata := true

	defer func() {
		if nodata {
			p.Err = ErrReadNoData
		}
	}()

	var idx int
	for idx = p.idx; idx < p.length; idx++ {
		c := p.buf[idx]
		if c >= '0' && c <= '9' {
			nodata = false
			v = v * 10
			v += uint(c - '0')
		} else {
			break
		}
	}
	p.idx = idx
	return
}

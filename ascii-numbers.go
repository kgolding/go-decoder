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

	for cnt, c := range p.buf {
		if cnt == 0 && c == '-' {
			negate = true
			p.idx++
		} else {
			if c >= '0' && c <= '9' {
				p.idx++
				nodata = false
				v = v * 10
				v += int(c - '0')
			} else {
				return
			}
		}
	}
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

	for _, c := range p.buf {
		if c >= '0' && c <= '9' {
			p.idx++
			nodata = false
			v = v * 10
			v += uint(c - '0')
		} else {
			return
		}
	}
	return
}

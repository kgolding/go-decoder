package decoder

// Bits8 returns an array of bool where [0] is the right hand bit at the internal pointer and increments it accordingly
func (p *Packet) Bits8() []bool {
	b := p.Byte()
	if p.Err != nil {
		return nil
	}
	v := make([]bool, 8)
	for i := range v {
		v[i] = b&0x01 != 0
		b = b >> 1
	}
	return v
}

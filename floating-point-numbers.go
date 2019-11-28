package decoder

import "math"

// Float32 returns the value at the internal pointer and increments it accordingly
func (p *Packet) Float32() float32 {
	i := p.Uint32()
	if p.Err != nil {
		return 0
	}
	return math.Float32frombits(i)
}

// Float64 returns the value at the internal pointer and increments it accordingly
func (p *Packet) Float64() float64 {
	i := p.Uint64()
	if p.Err != nil {
		return 0
	}
	return math.Float64frombits(i)
}

package decoder

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

var data1 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}

func Test_Bytes(t *testing.T) {
	p := New(data1)
	p.Byte() // skip 1
	b := p.Bytes(3)
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if bytes.Compare(b, data1[1:4]) != 0 {
		t.Errorf("expected % X got % X", b, data1[1:4])
	}
	c := p.Byte()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if c != 0x04 {
		t.Errorf("expected % X got % X", 0x04, c)
	}

	p = New(data1)
	b = p.Bytes(0)
	if bytes.Compare(b, []byte{}) != 0 {
		t.Errorf("got % X expected {}", b)
	}

	b = p.Bytes(1)
	e := data1[0:1]
	if bytes.Compare(b, e) != 0 {
		t.Errorf("got % X expected % X", b, e)
	}

	b = p.Bytes(2)
	e = data1[1:3]
	if bytes.Compare(b, e) != 0 {
		t.Errorf("got % X expected % X", b, e)
	}

	b = p.Bytes(3)
	e = data1[3:6]
	if bytes.Compare(b, e) != 0 {
		t.Errorf("got % X expected % X", b, e)
	}
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}

	b = p.Bytes(9999)
	if p.Err == nil {
		t.Error("expected error reading past ed of data")
	}

}

func Test_Float64(t *testing.T) {
	float := float64(12345678.87654321)

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(float))

	p := New(b)
	f := p.Float64()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if f != float {
		t.Errorf("expected %f got %f", float, f)
	}
	binary.BigEndian.PutUint64(b, math.Float64bits(f))
}

func Test_Float32(t *testing.T) {
	float := float32(1234.4321)

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(float))

	p := New(b)
	f := p.Float32()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if f != float {
		t.Errorf("expected %f got %f", float, f)
	}
	binary.BigEndian.PutUint32(b, math.Float32bits(f))
}

func Test_CString(t *testing.T) {
	p := New([]byte{0x31, 0x32, 0x33, 0x00, 0xff}) // "123" 0x00, 0xff
	s := p.CString()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	b := p.Byte()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if b != 0xff {
		t.Errorf("expected 0xff got %X", b)
	}

	p = New([]byte{0x31, 0x32, 0x33, 0x00}) // "123" 0x00
	s = p.CString()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	// Test bad string
	p = New([]byte{0x31, 0x32, 0x33, 0xff}) // "123"
	s = p.CString()
	if p.Err != ErrReadPastEndData {
		t.Errorf("expected err: %s", p.Err)
	}
}

func Test_StringPrefixUint16Len(t *testing.T) {
	p := New([]byte{0x00, 0x03, 0x31, 0x32, 0x33, 0xff}) // 0x03 "123" 0xff
	s := p.StringPrefixUint16Len()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	b := p.Byte()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if b != 0xff {
		t.Errorf("expected 0xff got %X", b)
	}

	p = New([]byte{0x00, 0x03, 0x31, 0x32, 0x33}) // 0x03 "123"
	s = p.StringPrefixUint16Len()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}
}

func Test_StringPrefixByteLen(t *testing.T) {
	p := New([]byte{0x03, 0x31, 0x32, 0x33, 0xff}) // 0x03 "123" 0xff
	s := p.StringPrefixByteLen()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	b := p.Byte()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if b != 0xff {
		t.Errorf("expected 0xff got %X", b)
	}

	p = New([]byte{0x03, 0x31, 0x32, 0x33}) // 0x03 "123"
	s = p.StringPrefixByteLen()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

}

func Test_StringZeroPadded(t *testing.T) {
	p := New([]byte{0x03, 0x31, 0x32, 0x33, 0x00, 0x00, 0xff}) // 0x03 "123.." 0xff
	p.Byte()                                                   // STX

	s := p.StringZeroPadded(5)
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	b := p.Byte()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if b != 0xff {
		t.Errorf("expected 0xff got %X", b)
	}

	p = New([]byte{0x31, 0x32, 0x33, 0x00})
	s = p.StringZeroPadded(4)
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

}

func Test_StringByDelimiter(t *testing.T) {
	p := New([]byte{0xff, 0x31, 0x32, 0x33, 0x0A, 0xff}) // 0xff "123" LF 0xff
	p.Byte()                                             // clear first dummy byte

	s := p.StringByDelimiter(0x0A)
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	b := p.Byte()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if b != 0xff {
		t.Errorf("expected 0xff got %X", b)
	}

	// Read past end
	s = p.StringByDelimiter(0x0A)
	if s != "" {
		t.Errorf("expected string '' got '%s'", s)
	}
}

func Test_Rewind(t *testing.T) {
	p := New(data1)

	test := func() {
		p.Byte() // move past first byte
		if b2 := p.Byte(); b2 != 0x01 {
			t.Errorf("expected second byte to be 0x01 got %X. Idx is now %d", b2, p.idx)
		}
	}

	test()
	err := p.Rewind(2)
	if err != nil {
		t.Errorf("got unexpected err: %s", err)
	}
	test()

	p.Reset()
	p.Byte()
	p.Byte()
	err = p.Rewind(2)
	if err != nil {
		t.Errorf("got unexpected err: %s", err)
	}
	test()
}

func Test_Bit8(t *testing.T) {
	p := New(data1)

	p.Byte() // move past first byte

	v := p.Bits8()
	if p.Err != nil {
		t.Errorf("got unexpected err: %s", p.Err)
	}
	if len(v) != 8 {
		t.Errorf("expected 8 bits got %d", len(v))
	}

	for i, q := range v {
		switch i {
		case 0:
			if !q {
				t.Errorf("expect bit %d to be true got %t", i, q)
			}
		default:
			if q {
				t.Errorf("expect bit %d to be false got %t", i, q)
			}
		}
	}
}

func Test_Byte_Reset(t *testing.T) {
	p := New(data1)

	test := func() {
		for _, b := range data1 {
			v := p.Byte()
			if v != b {
				t.Errorf("expected %X got %X", b, v)
			}
			if p.Err != nil {
				t.Errorf("got unexpected err: %s", p.Err)
			}
		}
	}

	test()
	p.Reset()
	test()
}

func Test_Uint16(t *testing.T) {
	p := New(data1)

	for i := 0; i < len(data1); i += 2 {
		d := binary.BigEndian.Uint16(data1[i:])
		v := p.Uint16()
		if v != d {
			t.Errorf("expected %X got %X", d, v)
		}
		if p.Err != nil {
			t.Errorf("got unexpected err: %s", p.Err)
		}
	}
}

func Test_Uint32(t *testing.T) {
	p := New(data1)

	for i := 0; i < len(data1)-4; i += 4 {
		d := binary.BigEndian.Uint32(data1[i:])
		v := p.Uint32()
		if v != d {
			t.Errorf("expected %X got %X", d, v)
		}
		if p.Err != nil {
			t.Errorf("got unexpected err: %s", p.Err)
		}
	}
}

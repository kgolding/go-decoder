package decoder

import (
	"encoding/binary"
	"testing"
)

var data1 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}

func Test_CString(t *testing.T) {
	p := New([]byte{0x31, 0x32, 0x33, 0x00, 0xff}) // "123" 0x00, 0xff
	s := p.CString()
	if p.err != nil {
		t.Errorf("got unexpected err: %s", p.err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	b := p.Byte()
	if p.err != nil {
		t.Errorf("got unexpected err: %s", p.err)
	}
	if b != 0xff {
		t.Errorf("expected 0xff got %X", b)
	}

	// Test bad string
	p = New([]byte{0x31, 0x32, 0x33, 0xff}) // "123" 0xff
	s = p.CString()
	if p.err != ErrReadPastEndData {
		t.Errorf("expected err: %s", p.err)
	}

}

func Test_StringPrefixUint16Len(t *testing.T) {
	p := New([]byte{0x00, 0x03, 0x31, 0x32, 0x33, 0xff}) // 0x03 "123" 0xff
	s := p.StringPrefixUint16Len()
	if p.err != nil {
		t.Errorf("got unexpected err: %s", p.err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	b := p.Byte()
	if p.err != nil {
		t.Errorf("got unexpected err: %s", p.err)
	}
	if b != 0xff {
		t.Errorf("expected 0xff got %X", b)
	}
}

func Test_StringPrefixByteLen(t *testing.T) {
	p := New([]byte{0x03, 0x31, 0x32, 0x33, 0xff}) // 0x03 "123" 0xff
	s, err := p.StringPrefixByteLen()
	if err != nil {
		t.Errorf("got unexpected err: %s", err)
	}
	if s != "123" {
		t.Errorf("expected string '123' got '%s'", s)
	}

	b := p.Byte()
	if p.err != nil {
		t.Errorf("got unexpected err: %s", err)
	}
	if b != 0xff {
		t.Errorf("expected 0xff got %X", b)
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
	if p.err != nil {
		t.Errorf("got unexpected err: %s", p.err)
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
			if p.err != nil {
				t.Errorf("got unexpected err: %s", p.err)
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
		if p.err != nil {
			t.Errorf("got unexpected err: %s", p.err)
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
		if p.err != nil {
			t.Errorf("got unexpected err: %s", p.err)
		}
	}
}

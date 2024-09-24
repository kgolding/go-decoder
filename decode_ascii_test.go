package decoder

import (
	"fmt"
	"testing"
)

func Test_AsciiIntUintWithPrependedData(t *testing.T) {
	s := "abcd#1234xxx"

	dec := New([]byte(s))
	dec.StringByDelimiter('#')
	x := dec.AsciiInt()
	if x != 1234 {
		t.Errorf("expected 1234, but got %d", x)
	}

	dec = New([]byte(s))
	dec.StringByDelimiter('#')
	z := dec.AsciiUInt()
	if z != 1234 {
		t.Errorf("expected 1234, but got %d", z)
	}
}

func Test_AsciiInt(t *testing.T) {
	tests := []struct {
		expect      int
		shouldError bool
		input       []byte
	}{
		{0, true, []byte("")},
		{0, true, []byte("NONUMBERSHERE")},
		{0, true, []byte("X")},
		{5, false, []byte("5")},
		{5, false, []byte("5X")},
		{0, false, []byte("0")},
		{0, false, []byte("0X")},
		{123, false, []byte("123")},
		{123, false, []byte("123X")},
		{-123, false, []byte("-123")},
		{-123, false, []byte("-123X")},
		{1, false, []byte{0x31, 0x02, 0x31}},
	}

	for _, test := range tests {
		p := New(test.input)
		f := p.AsciiInt()
		if test.shouldError && p.Err == nil {
			t.Errorf("expected err, but got nil: '%s'", string(test.input))
		}
		if !test.shouldError && p.Err != nil {
			t.Errorf("got unexpected err: %s: '%s'", p.Err, string(test.input))
		}
		if f != test.expect {
			t.Errorf("expected %d got %d: '%s'", test.expect, f, string(test.input))
		}
		if !test.shouldError { // Check that the idx has been moved forward
			l := len(fmt.Sprintf("%d", f))
			q := len(test.input) - len(p.PeekRemainingBytes())
			if q != l {
				t.Errorf("expected bytes used to be %d got %d: '%s'", l, q, string(test.input))
			}
		}
	}
}

func Test_AsciiUInt(t *testing.T) {
	tests := []struct {
		expect      uint
		shouldError bool
		input       []byte
	}{
		{0, true, []byte("")},
		{0, true, []byte("NONUMBERSHERE")},
		{0, true, []byte("X")},
		{5, false, []byte("5")},
		{5, false, []byte("5X")},
		{0, false, []byte("0")},
		{0, false, []byte("0X")},
		{123, false, []byte("123")},
		{123, false, []byte("123X")},
		{0, true, []byte("-123")},
		{0, true, []byte("-123X")},
		{1, false, []byte{0x31, 0x02, 0x31}},
	}

	for _, test := range tests {
		p := New(test.input)
		f := p.AsciiUInt()
		if test.shouldError && p.Err == nil {
			t.Errorf("expected err, but got nil: '%s'", string(test.input))
		}
		if !test.shouldError && p.Err != nil {
			t.Errorf("got unexpected err: %s: '%s'", p.Err, string(test.input))
		}
		if f != test.expect {
			t.Errorf("expected %d got %d: '%s'", test.expect, f, string(test.input))
		}
		if !test.shouldError { // Check that the idx has been moved forward
			l := len(fmt.Sprintf("%d", f))
			q := len(test.input) - len(p.PeekRemainingBytes())
			if q != l {
				t.Errorf("expected bytes used to be %d got %d: '%s'", l, q, string(test.input))
			}
		}
	}
}

package decoder

import "testing"

func TestStringHex(t *testing.T) {
	dot := byte('.')
	hash := byte('#')

	tests := []struct {
		skip     int
		expect   string
		input    []byte
		nextByte *byte
	}{
		{0, "0", []byte("0"), nil},
		{0, "", []byte(""), nil},
		{0, "", []byte("HELLO"), nil},
		{2, "f", []byte("xxf"), nil},
		{2, "e", []byte("xxeXX"), nil},
		{2, "", []byte("xxXX"), nil},
		{0, "FfEeDDCcBbAa9876543210", []byte("FfEeDDCcBbAa9876543210"), nil},
		{5, "FfEeDDCcBbAa9876543210", []byte("AAAAAFfEeDDCcBbAa9876543210."), &dot},
		{1, "000000", []byte("L000000#1234[#1234|1140 00 007]_22:49:34,01-22-2012"), &hash},
		{1, "1234", []byte("#1234[#1234|1140 00 007]_22:49:34,01-22-2012"), nil},
	}

	for _, test := range tests {
		p := New(test.input)
		if test.skip > 0 {
			p.Bytes(test.skip)
		}
		f := p.StringHex()
		if f != test.expect {
			t.Errorf("expected %s got '%s': '%s'", test.expect, f, string(test.input))
		}
		if test.nextByte != nil {
			nb := p.Byte()
			if *test.nextByte != nb {
				t.Errorf("expected next byte %X got '%X': '%s'", *test.nextByte, nb, string(test.input))
			}
		}
	}
}

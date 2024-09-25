package decoder

import "testing"

func TestStringHex(t *testing.T) {
	tests := []struct {
		skip   int
		expect string
		input  []byte
	}{
		{0, "0", []byte("0")},
		{0, "", []byte("")},
		{0, "", []byte("HELLO")},
		{2, "f", []byte("xxf")},
		{2, "e", []byte("xxeXX")},
		{2, "", []byte("xxXX")},
		{0, "FfEeDDCcBbAa9876543210", []byte("FfEeDDCcBbAa9876543210")},
		{5, "FfEeDDCcBbAa9876543210", []byte("AAAAAFfEeDDCcBbAa9876543210.")},
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
	}
}

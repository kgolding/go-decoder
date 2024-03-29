package decoder

import (
	"encoding/binary"
	"testing"
)

func TestUint24(t *testing.T) {
	tests := []struct {
		data   []byte
		endian binary.ByteOrder
		val    uint32
	}{
		{[]byte{0x01, 0x23, 0x45}, binary.BigEndian, 0x00012345},
		{[]byte{0x45, 0x23, 0x01}, binary.LittleEndian, 0x00012345},
	}

	marker := byte(0xfe)

	for _, test := range tests {
		dec := New(append(test.data, marker))
		if test.endian == binary.BigEndian {
			dec.SetBigEndian()
		} else {
			dec.SetLittleEndian()
		}
		v := dec.Uint24()
		if v != test.val {
			t.Errorf("with % X expected %x got %x", test.data, test.val, v)
		}
		nextByte := dec.Byte()
		if nextByte != marker {
			t.Errorf("with % X expected next byte to be %x got %x", test.data, marker, nextByte)
		}
		if dec.Err != nil {
			t.Error(dec.Err)
		}
	}
}

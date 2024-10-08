package main

import (
	"log"

	"github.com/kgolding/go-decoder"
)

type Data struct {
	Uint16 uint16
	String string
}

func main() {
	// Here's a sample made up packet of data
	buf := []byte{
		decoder.STX, // STX Byte
		0xde, 0xad,  // Uint16 0xdead!
		0x41, 0x42, 0x43, decoder.NULL, // String "ABC" NULL
		decoder.ETX, // ETX Byte
	}

	Example1(buf)
	Example2(buf)

}

// Example2 shows typical usage
func Example2(buf []byte) {
	log.SetPrefix("Example2: ")

	// Create a load the data into a new *Packet
	dec := decoder.New(buf)

	// Create a Data instance to put the packet into
	mydata := Data{}

	// Decode the packet, accessing the elements in the order they appear
	if dec.Byte() != decoder.STX {
		log.Fatalln("Missing STX")
	}

	mydata.Uint16 = dec.Uint16()
	mydata.String = dec.CString()

	if dec.Byte() != decoder.ETX {
		log.Fatalln("Missing ETX")
	}

	// Check if there were any errors along the way
	if err := dec.Err; err != nil {
		log.Fatalln(err)
	}

	log.Printf("Read packet in Data{} OK: %+v\n", mydata)
}

// Example1 is a verbose form of decoding a packet
func Example1(buf []byte) {
	log.SetPrefix("Example1: ")

	// Create a load the data into a new *Packet
	dec := decoder.New(buf)

	// Read the first byte and valid it as an STX
	if b := dec.Byte(); b != decoder.STX {
		if err := dec.Err; err != nil {
			log.Fatalln("decoding STX:", err)
		}
		log.Fatalf("Expected STX as start byte, got %X\n", b)
	}
	log.Println("Read start STX byte ok")

	// Read the uint16
	if u := dec.Uint16(); u != 0xdead {
		if err := dec.Err; err != nil {
			log.Fatalln("decoding uint16:", err)
		}
		log.Fatalf("Expected Uint16 to be 0xdead, got %X\n", u)
	}
	log.Print("Read uint16 ok")

	// Read the null terminated string
	if s := dec.CString(); s != "ABC" {
		if err := dec.Err; err != nil {
			log.Fatalln("decoding string:", err)
		}
		log.Fatalf("Expected string to be 'ABC', got 'ABC'\n", s)
	}
	log.Print("Read string ok")

	// Read the last byte and validate it as an ETX
	if b := dec.Byte(); b != decoder.ETX {
		if err := dec.Err; err != nil {
			log.Fatalln("decoding ETX:", err)
		}
		log.Fatalf("Expected ETX as end byte, got %X\n", b)
	}
	log.Println("Read ETX byte ok")

}

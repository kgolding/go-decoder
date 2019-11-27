package decoder

// ASCII standard consts
const (
	NULL  = 0x00 // Null
	SOH   = 0x01 // Start of heading
	STX   = 0x02 // Start of text
	ETX   = 0x03 // End of text
	EOT   = 0x04 // End of transmission
	ENQ   = 0x05 // Enquiry
	ACK   = 0x06 // Acknowledgement
	BEL   = 0x07 // Bell
	BS    = 0x08 // Backspace
	TAB   = 0x09 // Horizontal tab
	LF    = 0x0a // Line feed
	VT    = 0x0b // Vertical tab
	NP    = 0x0c // Form feed
	CR    = 0x0d // Carriage return
	SI    = 0x0e // Shift in
	SO    = 0x0f // Shift out
	DLE   = 0x10 // Data link escape
	DC1   = 0x11 // Device control 1
	DC2   = 0x12 // Device control 2
	DC3   = 0x13 // Device control 3
	DC4   = 0x14 // Device control 4
	NAK   = 0x15 // Negative acknowledgement
	SYN   = 0x16 // Synchronous idle
	ETB   = 0x17 // End transmission block
	CAN   = 0x18 // Cancel
	EM    = 0x19 // End of medium
	SUB   = 0x1a // Substitution
	ESC   = 0x1b // Escape
	FS    = 0x1c // File separator
	GS    = 0x1d // Group separator
	RS    = 0x1e // Record separator
	US    = 0x1f // Unit separator
	SPACE = 0x32 // Space
	DEL   = 0x7f // Delete
	NBSP  = 0xff // Non breaking space
)

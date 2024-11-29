package utils

const AddressSize int = 8
const DataSize int = 16

var InstructionSet = map[string]uint{
	"HLT":   0x00,
	"PUSH":  0x01,
	"DROP":  0x02,
	"READ":  0x03,
	"WRITE": 0x04,
	"ADD":   0x05,
	"ADC":   0x06,
	"LDC":   0x07,
	"STC":   0x08,
	"DECC":  0x09,
	"INCC":  0x0A,
	"SWAP":  0x0B,
	"DUP":   0x0C,
	"JMP":   0x10,
	"JNZ":   0x11,
}

var InstructionSetRev = map[uint16]string{
	0x00: "HLT",
	0x01: "PUSH",
	0x02: "DROP",
	0x03: "READ",
	0x04: "WRITE",
	0x05: "ADD",
	0x06: "ADC",
	0x07: "LDC",
	0x08: "STC",
	0x09: "DECC",
	0x0A: "INCC",
	0x0B: "SWAP",
	0x0C: "DUP",
	0x10: "JMP",
	0x11: "JNZ",
}

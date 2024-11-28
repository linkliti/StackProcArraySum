package utils

var AddressSize int = 8
var DataSize int = 16

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

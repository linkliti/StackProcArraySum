package processor

import (
	"emulator/utils"
	"fmt"
)

type Processor struct {
	PC     uint8
	Stack  *Stack // Stack + SP
	CX     uint16
	FLAGS  [1]bool
	Memory [512]uint16
	Shift  uint16
}

func (p *Processor) Init() *Processor {
	p.Stack = NewStack(16)
	p.Shift = 256
	return p
}

func (p *Processor) String(isInst bool) string {
	// [true/false] -> [1/0]
	flags := make([]int, len(p.FLAGS))
	for i, flag := range p.FLAGS {
		flags[i] = utils.Bool2int(flag)
	}
	decodedPCMem := fmt.Sprintf("%d", p.Memory[p.PC])
	if isInst {
		inst, ok := utils.InstructionSetRev[p.Memory[p.PC]]
		decodedPCMem = inst
		if !ok {
			decodedPCMem = "???"
		}
	}

	pcPart := fmt.Sprintf("PC: 0x%x[%s]", p.PC, decodedPCMem)
	cxPart := fmt.Sprintf("CX: 0x%x", p.CX)
	stackPart := fmt.Sprintf("Stack: %s", p.Stack)
	spPart := fmt.Sprintf("SP: 0x%x", p.Stack.SP)
	flagsPart := fmt.Sprintf("FLAGS: %v", flags)

	return fmt.Sprintf("Processor{%s,  \t%s,\t%s,\t%s,\t%s}", pcPart, cxPart, flagsPart, spPart, stackPart)
}

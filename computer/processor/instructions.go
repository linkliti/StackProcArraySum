package processor

import (
	"emulator/utils"
	"log/slog"
	"os"
)

func (p *Processor) Start() {
	var data uint16
	var t1 uint16
	var t2 uint16
	var tsum uint32
	const maxDataVal = uint32(1<<utils.DataSize) - 1

	for {
		slog.Info(p.String(true))
		data = p.Memory[p.PC]
		inst := utils.InstructionSetRev[data]
		switch inst {
		case "HLT":
			return
		case "PUSH":
			p.PC++
			slog.Info(p.String(false))
			data = p.Memory[p.PC]
			p.Stack.Push(data)
		case "DROP":
			p.Stack.Pop()
		case "READ":
			t1 = p.Stack.Pop().(uint16)
			p.Stack.Push(p.Memory[t1])
		case "WRITE":
			t1 = p.Stack.Pop().(uint16)
			t2 = p.Stack.Pop().(uint16)
			p.Memory[t1] = t2
		case "ADD":
			t1 = p.Stack.Pop().(uint16)
			t2 = p.Stack.Pop().(uint16)
			tsum = uint32(t1) + uint32(t2)
			p.Stack.Push(uint16(tsum))
			if tsum > maxDataVal {
				p.FLAGS[0] = true
			} else {
				p.FLAGS[0] = false
			}
		case "ADC":
			t1 = p.Stack.Pop().(uint16)
			t2 = p.Stack.Pop().(uint16)
			tsum = uint32(t1) + uint32(t2) + uint32(utils.Bool2int(p.FLAGS[0]))
			p.Stack.Push(uint16(tsum))
			if tsum > maxDataVal {
				slog.Error("ADC: Overflow", "tsum", tsum)
				os.Exit(1)
			}
		case "LDC":
			t1 = p.Stack.Pop().(uint16)
			p.CX = t1
		case "STC":
			p.Stack.Push(p.CX)
		case "DECC":
			p.CX--
		case "INCC":
			p.CX++
		case "SWAP":
			t1 = p.Stack.Pop().(uint16)
			t2 = p.Stack.Pop().(uint16)
			p.Stack.Push(t1)
			p.Stack.Push(t2)
		case "DUP":
			t1 = p.Stack.Pop().(uint16)
			p.Stack.Push(t1)
			p.Stack.Push(t1)
		case "JMP":
			t1 = p.Stack.Pop().(uint16)
			p.PC = uint8(t1)
			continue
		case "JNZ":
			t1 = p.Stack.Pop().(uint16)
			if p.CX != 0 {
				p.PC = uint8(t1)
				continue
			}
		default:
			slog.Error("Unknown instruction: " + inst)
			os.Exit(1)
		}
		p.PC++
	}
}

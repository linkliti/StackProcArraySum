package readers

import (
	"bufio"
	"emulator/computer/processor"
	"emulator/utils"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func MemoryFill(binFile *os.File, comp *processor.Processor) {
	isInst := false
	reader := bufio.NewScanner(binFile)
	for reader.Scan() {
		text := reader.Text()
		if text == "" {
			isInst = true
			slog.Debug("Skipping empty line")
			continue
		}
		parts := strings.Split(text, " ")
		if len(parts) != 2 {
			slog.Error("Invalid binary format", "text", text)
			os.Exit(1)
		}
		addr, err1 := strconv.ParseUint(parts[0], 2, utils.AddressSize)
		data, err2 := strconv.ParseUint(parts[1], 2, utils.DataSize)
		if err1 != nil || err2 != nil {
			slog.Error("Invalid binary value", "text", text)
			os.Exit(1)
		}
		slog.Debug("Filling memory", "isInst", isInst, "addr", addr, "data", data)
		if isInst {
			comp.Memory[addr] = uint16(data)
		} else {
			comp.Memory[addr+uint64(comp.Shift)] = uint16(data)
		}
	}
}

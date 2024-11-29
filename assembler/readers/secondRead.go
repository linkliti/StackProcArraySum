package readers

import (
	"bufio"
	"emulator/utils"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func SecondRead(inputFile *os.File, outputFile *os.File, jmpMarkers map[string]uint) {
	slog.Debug("Starting second read")
	instructionNum := new(uint)
	if _, ok := jmpMarkers["START"]; ok {
		// JMP START and marker shift
		for k := range jmpMarkers {
			jmpMarkers[k] += 3
		}
		slog.Debug("Shifted markers", "markers", utils.PrintMap(jmpMarkers))
		slog.Debug("Writing PUSH START JMP", "startAddress", jmpMarkers["START"])
		WriteBinCount(outputFile, instructionNum, utils.InstructionSet["PUSH"])
		WriteBinCount(outputFile, instructionNum, jmpMarkers["START"])
		WriteBinCount(outputFile, instructionNum, utils.InstructionSet["JMP"])
	}
	reader := bufio.NewScanner(inputFile)
	ok := reader.Scan()
	for ok {
		text := CleanText(reader.Text())
		if text == "" {
			slog.Debug("Skipping empty line")
			ok = reader.Scan()
			continue
		} else if strings.HasPrefix(text, ".") {
			slog.Debug("Found section", "text", text)
			if text == ".DATA" {
				slog.Debug("Skipping .DATA")
				ok = reader.Scan()
				text = CleanText(reader.Text())
				for ok && !strings.HasPrefix(text, ".") {
					ok = reader.Scan()
					text = CleanText(reader.Text())
				}
				continue
			}
			slog.Debug("Parsing section", "section", text)
			ok = instructParser(reader, outputFile, instructionNum, jmpMarkers)
			continue
		}
		slog.Error("Invalid section", "text", text)
		os.Exit(1)
	}
	if *instructionNum == 0 {
		WriteBin(outputFile, 0, 0)
	}
	outputFile.WriteString("\n")
}

func instructParser(reader *bufio.Scanner, outputFile *os.File, instructionNum *uint, jmpMarkers map[string]uint) bool {
	sectionName := CleanText(reader.Text())
	sectionName = string([]rune(sectionName)[1:]) // Remove .
	slog.Debug("Start parsing section", "section", sectionName, "instruction", *instructionNum)
	ok := reader.Scan()
	text := CleanText(reader.Text())
	for ok && !strings.HasPrefix(text, ".") {
		if len(text) == 0 {
			slog.Debug("Skipping empty line")
			ok = reader.Scan()
			text = CleanText(reader.Text())
			continue
		}
		parts := strings.Split(text, " ")
		instruction := parts[0]
		// if _, ok := utils.InstructionSet[instruction]; !ok {
		// 	slog.Error("Invalid instruction", "instruction", instruction)
		// 	os.Exit(1)
		// }
		slog.Debug("Found instruction", "instruction", text)
		switch instruction {
		case "PUSH":
			data, err := strconv.ParseUint(parts[1], 10, utils.DataSize)
			if err != nil {
				slog.Error("Invalid value for instruction", "text", text)
				os.Exit(1)
			}
			WriteBinCount(outputFile, instructionNum, utils.InstructionSet["PUSH"])
			WriteBinCount(outputFile, instructionNum, data)
		case "JMP", "JNZ":
			data := jmpMarkers[parts[1]]
			WriteBinCount(outputFile, instructionNum, utils.InstructionSet["PUSH"])
			WriteBinCount(outputFile, instructionNum, data)
			WriteBinCount(outputFile, instructionNum, utils.InstructionSet[instruction])
		default:
			WriteBinCount(outputFile, instructionNum, utils.InstructionSet[instruction])
		}
		ok = reader.Scan()
		text = CleanText(reader.Text())
	}
	slog.Debug("Finished parsing section", "section", sectionName, "instruction", *instructionNum)
	return ok
}

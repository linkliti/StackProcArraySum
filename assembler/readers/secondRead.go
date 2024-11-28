package readers

import (
	"bufio"
	"log/slog"
	"os"
	"processor/assembler/utils"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func SecondRead(inputFile *os.File, outputFile *os.File, jmpMarkers *orderedmap.OrderedMap[string, uint]) {
	slog.Debug("Starting second read")
	instructionNum := new(uint)
	if address, ok := jmpMarkers.Get("START"); ok {
		utils.WriteBinCount(outputFile, instructionNum, utils.InstructionSet["PUSH"])
		utils.WriteBinCount(outputFile, instructionNum, address)
		utils.WriteBinCount(outputFile, instructionNum, utils.InstructionSet["JMP"])
	}
	reader := bufio.NewScanner(inputFile)
	ok := reader.Scan()
	for ok {
		text := utils.CleanText(reader.Text())
		if text == "" {
			slog.Debug("Skipping empty line")
			ok = reader.Scan()
			continue
		} else if strings.HasPrefix(text, ".") {
			slog.Debug("Found section", "text", text)
			if text == ".DATA" {
				slog.Debug("Skipping .DATA")
				ok = reader.Scan()
				text = utils.CleanText(reader.Text())
				for ok && !strings.HasPrefix(text, ".") {
					ok = reader.Scan()
					text = utils.CleanText(reader.Text())
				}
				continue
			}
			slog.Debug("Parsing section", "section", text)
			ok = instructParser(reader, outputFile, instructionNum)
			continue
		}
		slog.Error("Invalid section", "text", text)
		os.Exit(1)
	}
	if *instructionNum == 0 {
		utils.WriteBin(outputFile, 0, 0)
	}
	outputFile.WriteString("\n")
}

func instructParser(reader *bufio.Scanner, outputFile *os.File, instructionNum *uint) bool {
	sectionName := utils.CleanText(reader.Text())
	sectionName = string([]rune(sectionName)[1:]) // Remove .
	slog.Debug("Start parsing section", "section", sectionName, "instruction", *instructionNum)
	ok := reader.Scan()
	text := utils.CleanText(reader.Text())
	for ok && !strings.HasPrefix(text, ".") {
		if len(text) == 0 {
			slog.Debug("Skipping empty line")
			ok = reader.Scan()
			text = utils.CleanText(reader.Text())
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
			// if len(parts) != 2 {
			// 	slog.Error("Invalid instruction format", "text", text)
			// 	os.Exit(1)
			// }
			utils.WriteBinCount(outputFile, instructionNum, utils.InstructionSet["PUSH"])
			utils.WriteBinCount(outputFile, instructionNum, utils.InstructionSet[parts[1]])
		case "JMP", "JNZ":
			// if len(parts) != 2 {
			// 	slog.Error("Invalid instruction format", "text", text)
			// 	os.Exit(1)
			// }
			utils.WriteBinCount(outputFile, instructionNum, utils.InstructionSet["PUSH"])
			utils.WriteBinCount(outputFile, instructionNum, utils.InstructionSet[parts[1]])
			utils.WriteBinCount(outputFile, instructionNum, utils.InstructionSet[instruction])
		default:
			// if len(parts) != 1 {
			// 	slog.Error("Invalid instruction format", "text", text)
			// 	os.Exit(1)
			// }
			utils.WriteBinCount(outputFile, instructionNum, utils.InstructionSet[instruction])
		}
		ok = reader.Scan()
		text = utils.CleanText(reader.Text())
	}
	slog.Debug("Finished parsing section", "section", sectionName, "instruction", *instructionNum)
	return ok
}

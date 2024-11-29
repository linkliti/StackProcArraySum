package readers

import (
	"bufio"
	"emulator/utils"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func FirstRead(inputFile *os.File, outputFile *os.File, jmpMarkers map[string]uint) {
	reader := bufio.NewScanner(inputFile)
	slog.Debug("Starting first read")
	var instructionNum uint = 0
	var dataNum uint = 0
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
				slog.Debug("Parsing .DATA")
				ok = dataParser(reader, outputFile, &dataNum)
				continue
			}
			slog.Debug("Reading section", "section", text)
			ok = instructReader(reader, &instructionNum, jmpMarkers)
			continue
		}
		slog.Error("Invalid section", "text", text)
		os.Exit(1)
	}
	if dataNum == 0 {
		WriteBin(outputFile, 0, 0)
	}
	outputFile.WriteString("\n")
}

func instructReader(reader *bufio.Scanner, instructionNum *uint, jmpMarkers map[string]uint) bool {
	sectionName := CleanText(reader.Text())
	sectionName = string([]rune(sectionName)[1:]) // Remove .
	jmpMarkers[sectionName] = *instructionNum
	slog.Debug("Saved section", "section", sectionName, "instruction", *instructionNum)
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
		if _, ok := utils.InstructionSet[instruction]; !ok {
			slog.Error("Invalid instruction", "instruction", instruction)
			os.Exit(1)
		}
		switch instruction {
		case "PUSH":
			if len(parts) != 2 {
				slog.Error("Invalid instruction format", "text", text)
				os.Exit(1)
			}
			*instructionNum += 2
		case "JMP", "JNZ":
			if len(parts) != 2 {
				slog.Error("Invalid instruction format", "text", text)
				os.Exit(1)
			}
			*instructionNum += 3
		default:
			if len(parts) != 1 {
				slog.Error("Invalid instruction format", "text", text)
				os.Exit(1)
			}
			*instructionNum += 1
		}
		ok = reader.Scan()
		text = CleanText(reader.Text())
	}
	slog.Debug("Finished reading section", "section", sectionName, "instruction", *instructionNum)
	return ok
}

func dataParser(reader *bufio.Scanner, outputFile *os.File, dataNum *uint) bool {
	ok := reader.Scan()
	text := CleanText(reader.Text())
	for ok && !strings.HasPrefix(text, ".") {
		if len(text) == 0 {
			slog.Debug("Skipping empty line")
			ok = reader.Scan()
			text = CleanText(reader.Text())
			continue
		}
		nums := strings.Split(text, " ")
		for _, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				slog.Error("Failed to parse number", "numStr", numStr)
				os.Exit(1)
			}
			WriteBin(outputFile, *dataNum, num)
			*dataNum++
		}
		ok = reader.Scan()
		text = CleanText(reader.Text())
	}
	slog.Debug("Finished parsing .DATA")
	return ok
}

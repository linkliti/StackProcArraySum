package readers

import (
	"bufio"
	"log/slog"
	"os"
	"processor/assembler/utils"
	"strconv"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func FirstRead(inputFile *os.File, outputFile *os.File, jmpMarkers *orderedmap.OrderedMap[string, uint]) {
	reader := bufio.NewScanner(inputFile)
	slog.Debug("Starting first read")
	var instructionNum uint = 0
	var dataNum uint = 0
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
		utils.WriteBin(outputFile, 0, 0)
	}
	outputFile.WriteString("\n")
}

func instructReader(reader *bufio.Scanner, instructionNum *uint, jmpMarkers *orderedmap.OrderedMap[string, uint]) bool {
	sectionName := utils.CleanText(reader.Text())
	sectionName = string([]rune(sectionName)[1:]) // Remove .
	jmpMarkers.Set(sectionName, *instructionNum)
	slog.Debug("Saved section", "section", sectionName, "instruction", *instructionNum)
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
		text = utils.CleanText(reader.Text())
	}
	slog.Debug("Finished reading section", "section", sectionName, "instruction", *instructionNum)
	return ok
}

func dataParser(reader *bufio.Scanner, outputFile *os.File, dataNum *uint) bool {
	ok := reader.Scan()
	text := utils.CleanText(reader.Text())
	for ok && !strings.HasPrefix(text, ".") {
		if len(text) == 0 {
			slog.Debug("Skipping empty line")
			ok = reader.Scan()
			text = utils.CleanText(reader.Text())
			continue
		}
		nums := strings.Split(text, " ")
		for _, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				slog.Error("Failed to parse number", "numStr", numStr)
				os.Exit(1)
			}
			utils.WriteBin(outputFile, *dataNum, num)
			*dataNum++
		}
		ok = reader.Scan()
		text = utils.CleanText(reader.Text())
	}
	slog.Debug("Finished parsing .DATA")
	return ok
}

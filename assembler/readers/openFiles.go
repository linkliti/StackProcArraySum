package readers

import (
	"log/slog"
	"os"
)

func OpenASMFile(inputFilename string) *os.File {

	inputFile, err := os.Open(inputFilename)
	if err != nil {
		slog.Error("Failed to open input file", "error", err)
		os.Exit(1)
	}
	slog.Debug("File opened", "inputFile", inputFilename)
	return inputFile
}

func OpenBINFile(outputFilename string) *os.File {
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		slog.Error("Failed to create output file", "error", err)
		os.Exit(1)
	}
	slog.Debug("File created", "outputFile", outputFilename)
	return outputFile
}

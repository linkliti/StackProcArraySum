package main

import (
	"emulator/assembler/readers"
	"emulator/utils"
	"flag"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	inputFilename := flag.String("f", "../arraysum.asm", "Input assembly file")
	outputFilename := flag.String("o", "../arraysum.bin", "Output binary file")
	flag.Parse()
	inputFile := readers.OpenASMFile(*inputFilename)
	defer inputFile.Close()
	outputFile := readers.OpenBINFile(*outputFilename)
	defer outputFile.Close()
	jmpMarkers := map[string]uint{}
	readers.FirstRead(inputFile, outputFile, jmpMarkers)
	slog.Debug("Finished first read", "markers", utils.PrintMap(jmpMarkers))
	inputFile.Close()
	inputFile = readers.OpenASMFile(*inputFilename)
	readers.SecondRead(inputFile, outputFile, jmpMarkers)
	slog.Debug("Finished second read")
}

package main

import (
	"flag"
	"log/slog"
	"processor/assembler/readers"
	"processor/assembler/utils"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	inputFilename := flag.String("f", "./arraysum.asm", "Input assembly file")
	outputFilename := flag.String("o", "./arraysum.bin", "Output binary file")
	flag.Parse()
	inputFile := readers.OpenASMFile(*inputFilename)
	defer inputFile.Close()
	outputFile := readers.OpenBINFile(*outputFilename)
	defer outputFile.Close()
	var jmpMarkers = orderedmap.New[string, uint]()
	readers.FirstRead(inputFile, outputFile, jmpMarkers)
	slog.Debug("Finished first read", "markers", utils.PrintOrderedMap(jmpMarkers))
	inputFile.Close()
	inputFile = readers.OpenASMFile(*inputFilename)
	readers.SecondRead(inputFile, outputFile, jmpMarkers)
	slog.Debug("Finished second read")
}

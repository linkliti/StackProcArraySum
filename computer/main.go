package main

import (
	"emulator/computer/processor"
	"emulator/computer/readers"
	"flag"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	binFilename := flag.String("f", "../arraysum.bin", "Output binary file")
	flag.Parse()
	binFile := readers.OpenBINFile(*binFilename)
	defer binFile.Close()
	comp := new(processor.Processor).Init()
	readers.MemoryFill(binFile, comp)
	comp.Start()
}

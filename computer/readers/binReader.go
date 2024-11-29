package readers

import (
	"log/slog"
	"os"
)

func OpenBINFile(binFilename string) *os.File {
	binFile, err := os.Open(binFilename)
	if err != nil {
		slog.Error("Failed to open bin file", "error", err)
		os.Exit(1)
	}
	slog.Debug("File opened", "binFile", binFilename)
	return binFile
}

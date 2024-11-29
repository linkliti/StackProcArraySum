package readers

import (
	"emulator/utils"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/exp/constraints"
)

func CleanText(text string) string {
	parts := strings.Split(text, ";")
	text = strings.TrimSpace(parts[0])
	return text
}

func PadBinaryString(binStr string, size int) string {
	runeCount := utf8.RuneCountInString(binStr) // Letter count
	if runeCount <= size {
		return strings.Repeat("0", size-runeCount) + binStr
	}
	slog.Warn("Exceeded binary size", "binStr", binStr, "size", size)
	return string([]rune(binStr)[runeCount-size:])
}

func WriteBin[K constraints.Integer, V constraints.Integer](outputFile *os.File, i K, data V) {
	iBinStr := PadBinaryString(strconv.FormatInt(int64(i), 2), utils.AddressSize)
	dataBinStr := PadBinaryString(strconv.FormatInt(int64(data), 2), utils.DataSize)
	slog.Debug("Writing to output file", "i", i, "data", data)
	_, err := outputFile.WriteString(iBinStr + " " + dataBinStr + "\n")
	if err != nil {
		slog.Error("Failed to write to output file", "error", err)
		os.Exit(1)
	}
}

func WriteBinCount[V constraints.Integer](outputFile *os.File, count *uint, data V) {
	WriteBin(outputFile, *count, data)
	*count++
}

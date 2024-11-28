package utils

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	orderedmap "github.com/wk8/go-ordered-map/v2"
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
	iBinStr := PadBinaryString(strconv.FormatInt(int64(i), 2), AddressSize)
	dataBinStr := PadBinaryString(strconv.FormatInt(int64(data), 2), DataSize)
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

func PrintOrderedMap[V constraints.Integer](om *orderedmap.OrderedMap[string, V]) string {
	var result string
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		if result != "" {
			result += " "
		}
		result += fmt.Sprintf("%v:%v", pair.Key, pair.Value)
	}
	return fmt.Sprintf("orderedmap[%s]", result)
}

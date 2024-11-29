package main

import (
	"emulator/utils"
	"log/slog"
	"strconv"
)

func main() {
	str := "0000000000001100"
	data, err := strconv.ParseUint(str, 2, utils.DataSize)
	if err != nil {
		panic(err)
	}
	slog.Info("Data", "data", uint16(data))
}

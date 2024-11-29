package utils

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func PrintMap[V constraints.Integer](om map[string]V) string {
	var result string
	for k, v := range om {
		if result != "" {
			result += " "
		}
		result += fmt.Sprintf("%v:%v", k, v)
	}
	return fmt.Sprintf("map[%s]", result)
}

func Bool2int(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

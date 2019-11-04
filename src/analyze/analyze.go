package analyze

import (
	"fmt"
	"strings"
)

// Analyze analyze input string.
func Analyze(inputString string) {
	separatedArray := splitIntoArray(inputString)

	for index, item := range separatedArray {
		fmt.Printf("index: %d, item: %s\n", index, item)
	}
}

func splitIntoArray(inputString string) []string {
	return strings.Split(inputString, "\n")
}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func UnpackString(input string) (string, error) {
	var result strings.Builder
	var prevRune rune
	for i, r := range input {
		if unicode.IsDigit(r) {
			if i == 0 || unicode.IsDigit(prevRune) {
				return "", fmt.Errorf("incorrect string")
			}
			count, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(prevRune), count-1))
		} else {
			result.WriteRune(r)
		}
		prevRune = r
	}
	return result.String(), nil
}

func main() {
	examples := []string{"a4bc2d5e", "abcd", "45", ""}
	for _, example := range examples {
		unpacked, err := UnpackString(example)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("%s\n", unpacked)
		}
	}
}

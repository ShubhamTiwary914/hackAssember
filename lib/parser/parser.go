package assembler

import (
	"fmt"
	"strconv"
	Str "strings"
)

// region parsers
func To15BitBinary(numStr string) string {
	num, _ := strconv.Atoi(numStr)
	return fmt.Sprintf("%015b", num)
}

func RemoveComment(line *string) {
	commentIndex := Str.Index(*line, "//")
	if commentIndex != -1 {
		(*line) = (*line)[:commentIndex]
	}
}

func RemoveWhitespaceAndEmptyLines(line *string) bool {
	*line = Str.ReplaceAll(*line, " ", "")
	*line = Str.ReplaceAll(*line, "\t", "")
	*line = Str.TrimSpace(*line)
	return len(*line) > 0
}

// *Parse C-instruction
func ParseCInstruction(instruction *string) (string, string, string) {
	comp := "null"
	dest := "null"
	jump := "null"

	// Split by ';' for jump
	parts := Str.Split(*instruction, ";")
	if len(parts) > 1 {
		jump = parts[1]
	}

	// Split by '=' for comp and dest
	equalParts := Str.Split(parts[0], "=")
	if len(equalParts) > 1 {
		dest = equalParts[0]
		comp = equalParts[1]
	} else {
		comp = equalParts[0]
	}

	// Print results
	return comp, dest, jump
}

// "(label)" --> "label"
func ExtractLabel(input string) string {
	return input[1 : len(input)-1]
}

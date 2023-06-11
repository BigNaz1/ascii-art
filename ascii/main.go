package main

import (
	"fmt"
	"os"
	"strings"
	"work"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	// Writing arguments in a single string
	str := os.Args[1]
	for _, v := range os.Args[2:] {
		str += " " + v
	}

	// Replacing escaped newlines with newlines
	str = replaceEscapedNewlines(str)

	// Splitting the string into parts
	parts := strings.Split(str, "\n")
	var filteredParts []string
	for _, part := range parts {
		// Filtering out empty parts
		if part != "" {
			filteredParts = append(filteredParts, part)
		}
	}

	// Printing the parts
	for i, part := range filteredParts {
		// Adding the part to the result
		var res strings.Builder
		for j := 0; j < 8; j++ {
			for _, letter := range part {
				res.WriteString(work.Findline(2 + int(letter-' ')*9 + j))
			}
			fmt.Println(res.String())
			res.Reset()
		}
		// Adding a single newline after each part except the last one
		//fix soon
		if i < len(filteredParts)-1 {
			fmt.Print()
		}
	}
}

func replaceEscapedNewlines(input string) string {
	var result strings.Builder
	escaped := false
	for _, r := range input {
		if escaped {
			if r == 'n' {
				result.WriteRune('\n')
			} else {
				result.WriteRune('\\')
				result.WriteRune(r)
			}
			escaped = false
		} else if r == '\\' {
			escaped = true
		} else {
			result.WriteRune(r)
		}
	}
	if escaped {
		result.WriteRune('\\')
	}
	return result.String()
}
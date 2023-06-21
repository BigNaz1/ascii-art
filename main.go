package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

var asciiArt []string

func main() {
	if len(os.Args) == 1 {
		return
	}

	// Load ASCII art from file
	LoadAsciiArt()

	// Writing arguments in a single string
	str := os.Args[1]
	for _, v := range os.Args[2:] {
		str += " " + v
	}

	// Checking whether str contain "\\n" or not
	severallines := strings.Contains(str, "\\n")

	// Writing text line by line into res
	var res strings.Builder
	if severallines {
		args := strings.Split(str, "\\n")
		for _, word := range args {
			if word == "" {
				fmt.Println()
				continue
			}
			for i := 0; i < 8; i++ {
				for _, letter := range word {
					res.WriteString(getLine(1 + int(letter-' ')*9 + i))
				}
				if res.Len() > 0 {
					fmt.Println(res.String())
					res.Reset()
				}
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for _, letter := range str {
				res.WriteString(getLine(1 + int(letter-' ')*9 + i))
			}
			fmt.Println(res.String())
			res.Reset()
		}
	}
}


func LoadAsciiArt() {
	f, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}
}

func getLine(num int) string {
	if num < len(asciiArt) {
		return asciiArt[num]
	}
	return ""
}

package HussainAl7lo

import (
	"os"
	"bufio"
	"fmt"
)

var asciiArt []string

func Hasooni() {
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

func GetAsciiArt() []string {
	return asciiArt
}
package file

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func FromLinesToArray(lines []string) [][]string {
	var arr [][]string

	for _, line := range lines {
		words := strings.Fields(line)
		arr = append(arr, words)
	}
	return arr
}

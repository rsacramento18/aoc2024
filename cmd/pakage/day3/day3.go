package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rsacramento18/aoc2024/cmd/pakage/file"
)

func Run() {
	lines := file.ReadFile("../pakage/day3/input.txt")
	part1(lines)
}

func part1(arr []string) {
	sum := 0
	r := regexp.MustCompile(`mul\(\d*,\d*\)`)
	for _, line := range arr {
		matches := r.FindAllStringSubmatch(line, -1)
		fmt.Println("matches", matches)

		for _, operation := range matches {
			fmt.Println("operation", operation[0])
			fmt.Println("num1", string(operation[0][4]))
			fmt.Println("num2", string(operation[0][6]))

			num1, _ := strconv.Atoi(string(operation[0][4]))
			num2, _ := strconv.Atoi(string(operation[0][6]))

			sum += num1 * num2
		}
	}
	fmt.Println("Day3 ---- part1 ---- sum ----", sum)
}

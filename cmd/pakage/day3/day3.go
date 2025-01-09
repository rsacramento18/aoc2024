package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rsacramento18/aoc2024/cmd/pakage/file"
)

func Run() {
	lines := file.ReadFile("../pakage/day3/input4.txt")
	part1(lines)
	part2(lines)
}

func part1(arr []string) {
	sum := 0
	r := regexp.MustCompile(`mul\(\d*,\d*\)`)
	for _, line := range arr {
		matches := r.FindAllStringSubmatch(line, -1)

		for _, operation := range matches {
			substr := string(operation[0][4 : len(operation[0])-1])
			s := strings.Split(substr, ",")

			num1, _ := strconv.Atoi(s[0])
			num2, _ := strconv.Atoi(s[1])

			sum += num1 * num2
		}
	}
	fmt.Println("Day3 ---- part1 ---- sum ----", sum)
}

func part2(arr []string) {
	sum := 0
	r := regexp.MustCompile(`(don't\(\)|do\(\))|(mul\(\d*,\d*\))`)
	enabled := true
	for _, line := range arr {
		matches := r.FindAllString(line, -1)

		for _, match := range matches {
			switch match {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled == true {
					substr := string(match[4 : len(match)-1])
					s := strings.Split(substr, ",")

					num1, _ := strconv.Atoi(s[0])
					num2, _ := strconv.Atoi(s[1])

					sum += num1 * num2
				}
			}
		}
	}
	fmt.Println("Day3 ---- part2 ---- sum ----", sum)
}

func checkWord(s string) string {
	for i, char := range s {
		if string(char) == "(" {
			word := string(s[0:i])
			return word
		}
	}
	return ""
}

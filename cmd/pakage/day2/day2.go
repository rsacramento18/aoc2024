package day2

import (
	"fmt"
	"strconv"

	"github.com/rsacramento18/aoc2024/cmd/pakage/file"
)

func Run() {
	lines := file.ReadFile("../pakage/day2/input2.txt")
	arr := file.FromLinesToArray(lines)
	part1(arr)
	part2(arr)
}

func part1(arr [][]string) {
	count := 0

	for _, report := range arr {
		a, _ := strconv.Atoi(report[0])
		b, _ := strconv.Atoi(report[1])
		increasing := b-a > 0

		fmt.Println(report)
		isSafe := true
		for i := range report {
			if i > 0 {
				curr, _ := strconv.Atoi(report[i])
				before, _ := strconv.Atoi(report[i-1])

				diffIncreasing := curr - before
				diffDecreasing := before - curr

				if increasing {
					isSafe = 1 <= diffIncreasing && diffIncreasing <= 3
				} else {
					isSafe = 1 <= diffDecreasing && diffDecreasing <= 3
				}
				if isSafe == false {
					break
				}
			}
		}
		fmt.Println("isSafe", isSafe)
		if isSafe == true {
			count += 1
		}
	}
	fmt.Println("Day2 ---- part1 ---- count ----", count)
}

func part2(arr [][]string) {
	count := 0

	for _, report := range arr {
		a, _ := strconv.Atoi(report[0])
		b, _ := strconv.Atoi(report[1])
		increasing := b-a > 0

		isSafe := true
		error := false
		fmt.Println(report)
		for i := range report {
			if i > 0 {
				curr, _ := strconv.Atoi(report[i])
				before, _ := strconv.Atoi(report[i-1])

				diffIncreasing := curr - before
				diffDecreasing := before - curr

				if increasing {
					isSafe = 1 <= diffIncreasing && diffIncreasing <= 3
				} else {
					isSafe = 1 <= diffDecreasing && diffDecreasing <= 3
				}
				if isSafe == false && error == false {
					isSafe = true
					error = true
					fmt.Println("--------", i)
					report = remove(report, i)
					fmt.Println(report)
					i = 0
				} else if error == true {
					break
				}
			}
		}
		if isSafe == true {
			count += 1
		}
	}
	fmt.Println("Day2 ---- part1 ---- count ----", count)
}

func remove(slice []string, s int) []string {
	if s < 0 || s >= len(slice) {
		return append(slice[:s], slice[s+1:]...)
	} else {
		return slice
	}
}

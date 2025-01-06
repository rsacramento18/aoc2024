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

		fmt.Println(report)
		isSafe := true
		err := false
		for i := range report {
			if i > 0 {
				curr, _ := strconv.Atoi(report[i])
				var before int
				if err == true {
					before, _ = strconv.Atoi(report[i-2])
				} else {
					before, _ = strconv.Atoi(report[i-1])
				}

				diffIncreasing := curr - before
				diffDecreasing := before - curr

				if increasing {
					isSafe = 1 <= diffIncreasing && diffIncreasing <= 3
				} else {
					isSafe = 1 <= diffDecreasing && diffDecreasing <= 3
				}
				if isSafe == false && err == false {
					isSafe = true
					err = true
				} else if isSafe == false && err == true {
					break
				}
			}
		}
		fmt.Println("isSafe", isSafe)
		if isSafe == true {
			count += 1
		}
	}
	fmt.Println("Day2 ---- part2 ---- count ----", count)
}

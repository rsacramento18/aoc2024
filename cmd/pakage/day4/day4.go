package day4

import (
	"fmt"

	"github.com/rsacramento18/aoc2024/cmd/pakage/file"
)

type tuple struct {
	a int
	b int
}

func Run() {
	lines := file.ReadFile("../pakage/day4/input2.txt")
	arr := file.FromLinesToArrayWithoutWhiteSpaces(lines)
	part1(arr)
	part2(arr)
}

func part1(arr [][]string) {
	sum := 0

	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[x]); y++ {
			if arr[x][y] == "X" {
				sum += checkN(arr, x, y, "")
				sum += checkS(arr, x, y, "")
				sum += checkE(arr, x, y, "")
				sum += checkW(arr, x, y, "")
				sum += checkNE(arr, x, y, "")
				sum += checkSE(arr, x, y, "")
				sum += checkNW(arr, x, y, "")
				sum += checkSW(arr, x, y, "")
			}
		}
	}
	fmt.Println("Day4 ---- part1 ---- sum ----", sum)
}

func part2(arr [][]string) {
	sum := 0

	directions := []tuple{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	for x := 1; x < len(arr)-1; x++ {
		for y := 1; y < len(arr[x])-1; y++ {
			if arr[x][y] == "A" {
				var letter string
				for _, d := range directions {
					letter += arr[d.a+x][d.b+y]
				}
				if letter == "MMSS" || letter == "MSSM" || letter == "SSMM" || letter == "SMMS" {
					sum += 1
				}
			}
		}
	}
	fmt.Println("Day4 ---- part2 ---- sum ----", sum)
}

func inside(x int, y int, arr [][]string) bool {
	return x >= 0 && x < len(arr) && y >= 0 && y < len(arr[x])
}

func checkN(arr [][]string, x int, y int, res string) int {
	if res == "XMAS" {
		return 1
	} else if inside(x, y, arr) {
		res += arr[x][y]
		return checkN(arr, x-1, y, res)
	} else {
		return 0
	}
}

func checkS(arr [][]string, x int, y int, res string) int {
	if res == "XMAS" {
		return 1
	} else if inside(x, y, arr) {
		res += arr[x][y]
		return checkS(arr, x+1, y, res)
	} else {
		return 0
	}
}

func checkE(arr [][]string, x int, y int, res string) int {
	if res == "XMAS" {
		return 1
	} else if inside(x, y, arr) {
		res += arr[x][y]
		return checkE(arr, x, y+1, res)
	} else {
		return 0
	}
}

func checkW(arr [][]string, x int, y int, res string) int {
	if res == "XMAS" {
		return 1
	} else if inside(x, y, arr) {
		res += arr[x][y]
		return checkW(arr, x, y-1, res)
	} else {
		return 0
	}
}

func checkNE(arr [][]string, x int, y int, res string) int {
	if res == "XMAS" {
		return 1
	} else if inside(x, y, arr) {
		res += arr[x][y]
		return checkNE(arr, x-1, y+1, res)
	} else {
		return 0
	}
}

func checkNW(arr [][]string, x int, y int, res string) int {
	if res == "XMAS" {
		return 1
	} else if inside(x, y, arr) {
		res += arr[x][y]
		return checkNW(arr, x-1, y-1, res)
	} else {
		return 0
	}
}

func checkSE(arr [][]string, x int, y int, res string) int {
	if res == "XMAS" {
		return 1
	} else if inside(x, y, arr) {
		res += arr[x][y]
		return checkSE(arr, x+1, y+1, res)
	} else {
		return 0
	}
}

func checkSW(arr [][]string, x int, y int, res string) int {
	if res == "XMAS" {
		return 1
	} else if inside(x, y, arr) {
		res += arr[x][y]
		return checkSW(arr, x+1, y-1, res)
	} else {
		return 0
	}
}

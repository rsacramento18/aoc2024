package day6

import (
	"fmt"

	"github.com/rsacramento18/aoc2024/cmd/pakage/file"
)

func Run() {
	fmt.Println("Day 6")
	lines := file.ReadFile("../pakage/day6/input2.txt")
	arr := file.FromLinesToArrayWithoutWhiteSpaces(lines)

	part1(arr)
}

func part1(arr [][]string) {
	d := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	steps := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == "^" {
				steps = walk(arr, i, j, 0, 0, d)
			}
		}
	}
	for _, a := range arr {
		fmt.Println(a)
	}
	fmt.Println("Day6 ---- part1 ---- steps ----", steps)
}

func walk(arr [][]string, i, j, ct, steps int, d [][2]int) int {
	cd := d[ct]
	x := i + cd[0]
	y := j + cd[1]
	if x >= len(arr) || x < 0 || y >= len(arr[0]) || y < 0 {
		if arr[i][j] != "X" {
			steps++
		}
		arr[i][j] = "X"
		return steps
	} else if arr[x][y] == "#" {
		fmt.Println("ct", ct)
		ct = (ct + 1) % 4
		return walk(arr, i, j, ct, steps, d)
	} else {
		if arr[i][j] != "X" {
			steps++
		}
		arr[i][j] = "X"
		return walk(arr, x, y, ct, steps, d)
	}
}

package day4

import (
	"fmt"

	"github.com/rsacramento18/aoc2024/cmd/pakage/file"
)

func Run() {
	lines := file.ReadFile("../pakage/day4/input.txt")
	arr := file.FromLinesToArray(lines)
	part1(arr)
}

type Coordinate int

const (
	N Coordinate = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func part1(arr [][]string) {
}

func checkN(arr [][]string, x int, y int) (string, error) {
	if y > 0 {
		return arr[x-1][y], nil
	}
	return "", fmt.Errorf("Coordinate out of bounds")
}

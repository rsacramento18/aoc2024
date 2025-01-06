package day1

import (
	"fmt"
	"math"
	"strconv"

	"github.com/rsacramento18/aoc2024/cmd/pakage/file"
)

type distance struct {
	x int
	s int
	r int
}

func Run() {
	lines := file.ReadFile("../pakage/day1/input2.txt")
	arr := file.FromLinesToArray(lines)
	part1(arr)
	part2(arr)
}

func part1(arr [][]string) {
	orderedArr := orderPart1(arr)
	sum := totalDistance(orderedArr)
	fmt.Println("Part1 --- SUM --->", sum)
}

func part2(arr [][]string) {
	distances := []distance{}

	for _, coords := range arr {
		x, _ := strconv.Atoi(coords[0])

		seen := false
		for i := range distances {
			if distances[i].x == x {
				distances[i].s += 1
				seen = true
				break
			}
		}
		if seen == true {
			continue
		}

		newDist := distance{x: x, s: 1, r: 0}
		distances = append(distances, newDist)
	}

	for _, coords := range arr {
		r, _ := strconv.Atoi(coords[1])

		for i := range distances {
			if distances[i].x == r {
				distances[i].r += 1
			}
		}
	}

	sum := 0

	for _, dist := range distances {
		sum += dist.x * dist.s * dist.r
	}

	fmt.Println("Part2 --- SUM --->", sum)
}

func orderPart1(arr [][]string) [][]string {
	for _, coordsA := range arr {
		for _, coordsB := range arr {
			if coordsA[0] < coordsB[0] {
				temp := coordsB[0]
				coordsB[0] = coordsA[0]
				coordsA[0] = temp
			}

			if coordsA[1] < coordsB[1] {
				temp := coordsB[1]
				coordsB[1] = coordsA[1]
				coordsA[1] = temp
			}

		}
	}
	return arr
}

func totalDistance(arr [][]string) int {
	var sum int = 0
	for _, coords := range arr {
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		count := x - y

		sum += int(math.Abs(float64(count)))
	}
	return sum
}

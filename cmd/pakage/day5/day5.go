package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rules struct {
	x int
	y int
}

func Run() {
	rules, updates := getRulesUpdates()
	// for _, r := range rules {
	// 	fmt.Println(r)
	// }
	for _, u := range updates {
		fmt.Println(u)
	}

	part1(rules, updates)
}

func part1(rules []Rules, updates [][]int) {
	sum := 0
	for _, update := range updates {
		valid := true
		for i := range update {
			val := update[i]
			for _, rule := range rules {
				if rule.x == val {
					for j := i; j > 0; j-- {
						if update[j] == rule.y {
							valid = false
							break
						}
					}
				} else if rule.y == val {
					for j := i; j < len(update); j++ {
						if update[j] == rule.x {
							valid = false
							break
						}
					}
				}
				if !valid {
					break
				}
			}
		}
		if valid {
			middle := len(update) / 2
			sum += update[middle]
		}
	}
	fmt.Println(sum)
}

func getRulesUpdates() ([]Rules, [][]int) {
	var rules []Rules
	var updates [][]int
	file, err := os.Open("../pakage/day5/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	split := func(r rune) bool {
		return r == ',' || r == '|'
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			lineArr := []int{}
			for _, n := range strings.FieldsFunc(scanner.Text(), split) {

				j, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				lineArr = append(lineArr, j)
			}

			if len(lineArr) > 2 {
				updates = append(updates, lineArr)
			} else {
				rule := Rules{x: lineArr[0], y: lineArr[1]}
				rules = append(rules, rule)
			}
		}
	}
	return rules, updates
}

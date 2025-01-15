package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Rules struct {
	x int
	y int
}

func Run() {
	rules, updates := getRulesUpdates()
	// for _, u := range updates {
	// 	fmt.Println(u)
	// }

	notSafeUpdates := part1(rules, updates)
	part2(rules, notSafeUpdates)
}

func part1(rules []Rules, updates [][]int) [][]int {
	sum := 0
	notSafeUpdates := [][]int{}
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
		} else {
			notSafeUpdates = append(notSafeUpdates, update)
		}
	}
	fmt.Println("part1", sum)
	return notSafeUpdates
}

func part2(rules []Rules, nfu [][]int) {
	sum := 0
	for _, update := range nfu {
	INNERLOOP:
		problemFound := false
		for index := 1; index < len(update); index++ {
			toCheckFirst := []int{update[index-1], update[index]}
			if !ArrayContainsDeep(rules, toCheckFirst) {
				reflect.Swapper(update)(index-1, index)
				problemFound = true
			}
		}
		if problemFound {
			goto INNERLOOP
		}
		middle := len(update) / 2
		sum += update[middle]

	}
	fmt.Println("part2", sum)
}

func ArrayContainsDeep(rules []Rules, target []int) bool {
	for _, rule := range rules {
		if reflect.DeepEqual(rule.x, target[0]) && reflect.DeepEqual(rule.y, target[1]) {
			return true
		}
	}
	return false
}

func getRulesUpdates() ([]Rules, [][]int) {
	var rules []Rules
	var updates [][]int
	file, err := os.Open("../pakage/day5/input2.txt")
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

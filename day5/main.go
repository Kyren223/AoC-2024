package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	left  int
	right int
}

func main() {
	file, _ := os.ReadFile("example.txt")
	input := string(file)
	fmt.Println("Part 1 Example:", Part1(input))

	file, _ = os.ReadFile("input.txt")
	input = string(file)
	fmt.Println("Part 1:", Part1(input))

	file, _ = os.ReadFile("example.txt")
	input = string(file)
	fmt.Println("Part 2:", Part2(input))

	file, _ = os.ReadFile("input.txt")
	input = string(file)
	fmt.Println("Part 2:", Part2(input))
}

func Part2(input string) int {
	split := strings.Split(input, "\n\n")

	var rules []Rule

	ruleLines := strings.Split(split[0], "\n")
	for _, line := range ruleLines {
		if line == "" {
			continue
		}
		rule := strings.Split(line, "|")
		left, _ := strconv.ParseInt(rule[0], 10, 64)
		right, _ := strconv.ParseInt(rule[1], 10, 64)
		rules = append(rules, Rule{
			left:  int(left),
			right: int(right),
		})
	}

	var updates [][]int

	updateLines := strings.Split(split[1], "\n")
	for i, line := range updateLines {
		if line == "" {
			continue
		}
		updates = append(updates, []int{})
		update := strings.Split(line, ",")
		for _, up := range update {
			num, _ := strconv.ParseInt(up, 10, 64)
			updates[i] = append(updates[i], int(num))
		}
	}

	// fmt.Println("Rules:", rules, "\nUpdates:", updates)

	sum := 0
	for _, update := range updates {
		safe := true
		for _, rule := range rules {
			i := -1
			j := -1
			for k, up := range update {
				if rule.left == up {
					i = k
				} else if rule.right == up {
					j = k
				}
			}
			if i != -1 && j != -1 && i > j {
				// fmt.Println("Broke Rule: ", rule, "I", i, "J", j, "IAT", update[i], "JAT", update[j])
				safe = false
				break
			}
		}
		// fmt.Println("IsSafe:", safe, "Update:", update)
		if !safe {
			slices.SortFunc(update, func(a, b int) int {
				for _, rule := range rules {
					if rule.left == a && rule.right == b {
						return -1
					}
					if rule.left == b && rule.right == a {
						return 1
					}
				}
				return 0
			})
			// fmt.Println("After Sort:", update, "\nRules:", rules)
			middleIndex := (len(update) / 2)
			middle := update[middleIndex]
			// fmt.Println("Len:", len(update), "MI:", middleIndex, "Middle:", middle)
			sum += middle
		}
	}

	return sum
}

func Part1(input string) int {
	split := strings.Split(input, "\n\n")

	var rules []Rule

	ruleLines := strings.Split(split[0], "\n")
	for _, line := range ruleLines {
		if line == "" {
			continue
		}
		rule := strings.Split(line, "|")
		left, _ := strconv.ParseInt(rule[0], 10, 64)
		right, _ := strconv.ParseInt(rule[1], 10, 64)
		rules = append(rules, Rule{
			left:  int(left),
			right: int(right),
		})
	}

	var updates [][]int

	updateLines := strings.Split(split[1], "\n")
	for i, line := range updateLines {
		if line == "" {
			continue
		}
		updates = append(updates, []int{})
		update := strings.Split(line, ",")
		for _, up := range update {
			num, _ := strconv.ParseInt(up, 10, 64)
			updates[i] = append(updates[i], int(num))
		}
	}

	// fmt.Println("Rules:", rules, "\nUpdates:", updates)

	sum := 0
	for _, update := range updates {
		safe := true
		for _, rule := range rules {
			i := -1
			j := -1
			for k, up := range update {
				if rule.left == up {
					i = k
				} else if rule.right == up {
					j = k
				}
			}
			if i != -1 && j != -1 && i > j {
				// fmt.Println("Broke Rule: ", rule, "I", i, "J", j, "IAT", update[i], "JAT", update[j])
				safe = false
				break
			}
		}
		// fmt.Println("IsSafe:", safe, "Update:", update)
		if safe {
			middleIndex := (len(update) / 2)
			middle := update[middleIndex]
			// fmt.Println("Len:", len(update), "MI:", middleIndex, "Middle:", middle)
			sum += middle
		}
	}

	return sum
}
